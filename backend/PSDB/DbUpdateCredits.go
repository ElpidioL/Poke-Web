package PSDB

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	Defaults "github.com/ElpidioL/Poke-Web/Defaults"
	_ "github.com/lib/pq"
)

func UpdateCredits(Price float32, UserId int, pokemon Defaults.Pokemon) error {
	//storing the info to access the DB
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		DbInfo.Host, DbInfo.Port, DbInfo.User, DbInfo.Password, DbInfo.Dbname)

	//starting db
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//don't forget to close it
	defer db.Close()
	//ping to check if its working
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	sqlStatement := `SELECT credits,user_info FROM users_info WHERE id = $1;`
	row := db.QueryRow(sqlStatement, UserId)
	var dbCredits float32
	var dbInfo string
	err = row.Scan(&dbCredits, &dbInfo)

	switch err {

	case sql.ErrNoRows:
		return errors.New("credits does not exist?? userId invalid??")
	case nil:
		if dbCredits > Price {

			sqlStatement := `UPDATE users_info
					SET credits = $1
					WHERE id = $2;`
			_, err = db.Exec(sqlStatement, (dbCredits - Price), UserId)
			if err != nil {
				return errors.New("fail to update credits")
			}

			fmt.Println(string(dbInfo))
			dbStruct := Defaults.PokeList{}
			err = json.Unmarshal([]byte(string(dbInfo)), &dbStruct)
			if err != nil {
				fmt.Println(err)
				return errors.New("fail to update credits") //////////////////////////////////////////////
			}
			dbStruct.PokeToHatch.PokeId = append(dbStruct.PokeToHatch.PokeId, pokemon.PokeId)
			dbStruct.PokeToHatch.PokeDate = append(dbStruct.PokeToHatch.PokeDate, "10/10/10")

			newJson, err := json.Marshal(dbStruct)
			if err != nil {
				fmt.Println(err)
				return errors.New("fail to update credits1") //////////////////////////////////////////////
			}

			sqlStatement = `UPDATE users_info
					SET user_info = $1
					WHERE id = $2;`
			_, err = db.Exec(sqlStatement, newJson, UserId)
			if err != nil {
				fmt.Println(err)
				return errors.New("fail to update credits2") ///////////////////////////
			}

			return nil
		}
		return errors.New("low credits")
	default:
		panic(err)

	}
}
