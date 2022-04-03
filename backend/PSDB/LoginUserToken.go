package PSDB

import (
	"database/sql"
	"errors"
	"fmt"

	sanitizer "github.com/ElpidioL/Poke-Web/PasswordHandler"
)

func LoginUserToken(token string, email string) (string, error) {
	//storing the info to access the DB
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		DbInfo.Host, DbInfo.Port, DbInfo.User, DbInfo.Password, DbInfo.Dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	sqlStatement := `SELECT id,token FROM login_token WHERE email = $1`
	row := db.QueryRow(sqlStatement, email)
	var dbId int
	var dbToken string
	err = row.Scan(&dbId, &dbToken)

	switch err {
	case sql.ErrNoRows:
		return "", errors.New("invalid token")
	case nil:
		err := sanitizer.CheckHash(token, dbToken)
		if err != nil {
			return "", err
		}
		sqlStatement := `SELECT credits, user_info, last_update FROM users_info WHERE id = $1`
		row := db.QueryRow(sqlStatement, dbId)
		var dbcredits float32
		var dbinfo string
		var last string
		err = row.Scan(&dbcredits, &dbinfo, &last)
		switch err {
		case sql.ErrNoRows:
			return "", errors.New("something went really wrong")
		case nil:
			return fmt.Sprintf(`{"intent":"setInfo","credits":%v, "last":"%s", "dbId":%v, "info":"%v"}`, dbcredits, last, dbId, []byte(dbinfo)), nil
		default:
			panic(err)
		}
	default:
		panic(err)
	}
}
