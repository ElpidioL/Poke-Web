package IntentHandler

/////seria interessante se eu fizesse as verificações corretas para ter certeza que as infos que vieram
///// possuem as propriedades certas do intent, visto que o intent fica no front end e pode ser alterado.

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	Defaults "github.com/ElpidioL/Poke-Web/Defaults"
	DB "github.com/ElpidioL/Poke-Web/PSDB"
	pass "github.com/ElpidioL/Poke-Web/PasswordHandler"
	"github.com/rs/xid"
)

var UserInfo Defaults.UserInfo

func Intentions(choice []byte) string {
	Intents := Defaults.IntentDefine{}
	err := json.Unmarshal([]byte(string(choice)), &Intents)
	if err != nil {
		return fmt.Sprintf(`{"intent":"error", "msg":"%s, fail to Parse JSON1"}`, err.Error())
	}

	if Intents.Intent == "colour" {

		LoginToken := Defaults.TokenAcess{}
		err = json.Unmarshal([]byte(string(choice)), &LoginToken)
		if err != nil {
			return fmt.Sprintf(`{"intent":"error", "msg":"%s, fail to Parse JSON2"}`, err.Error())
		}
		info, err := DB.LoginUserToken(LoginToken.Colour, LoginToken.Email)
		if err != nil {
			return `{"intent":"error", "msg":"Fail to Colour"}`
		}
		vl, err := SaveInfo(info)
		if err != nil {
			fmt.Println(err)
			return fmt.Sprintf(`{"intent":"error", "msg":"%s, fail to HandShake"}`, err.Error())
		}
		UserInfo = vl
		return fmt.Sprintf(`{"intent":"%s","credits":%v,"info":"%s","last":"%s","session":"%s","dbId":%v}`, vl.Intent, vl.Credits, vl.Info, vl.Last, vl.Session, vl.DbId)
	}

	if Intents.Intent == "register" || Intents.Intent == "login" {
		registerUser := Defaults.Register{}
		err = json.Unmarshal([]byte(string(choice)), &registerUser)
		if err != nil {
			return fmt.Sprintf(`{"intent":"error", "msg":"%s, fail to Parse JSON3"}`, err.Error())
		}
		registerUser, err = pass.Sanitizer(registerUser)
		if err != nil {
			return `{"intent":"error", "msg":"Fail to sanitize"}`
		}

		if Intents.Intent == "register" {
			err := DB.UserRegister(registerUser.Email, registerUser.Name, registerUser.Password)
			if err != nil {
				return fmt.Sprintf(`{"intent":"error", "msg":"%s"}`, err.Error())
			}
			return `{"intent":"success", "msg":"sucess"}`

		}
		if Intents.Intent == "login" {
			token, err := pass.CreateToken()
			if err != nil {
				return `{"intent":"error", "msg":"Fail to create Colour"}`
			}
			hashToken, err := pass.SmallHash(token)
			if err != nil {
				return `{"intent":"error", "msg":"Fail to create HColour"}`
			}
			err = DB.LoginUser(registerUser.Email, registerUser.Password, hashToken)
			if err != nil {
				return fmt.Sprintf(`{"intent":"error", "msg":"%s"}`, err.Error())
			}
			return fmt.Sprintf(`{"intent":"colour","colour":"%s", "email":"%s"}`, token, registerUser.Email)
		}
	}
	if Intents.Intent == "pokemon" {
		if UserInfo.Credits > Defaults.PokemonPrice {
			NewPokemon := Defaults.Pokemon{}
			err := json.Unmarshal([]byte(string(choice)), &NewPokemon)
			if err != nil {
				return fmt.Sprintf(`{"intent":"error", "msg":"%s, fail to Parse JSON4"}`, err.Error())
			}

			err = CompareSessions(NewPokemon.Session)
			if err != nil {
				return fmt.Sprintf(`{"intent":"error", "msg":"%s, wrong session"}`, err.Error())
			}

			err = DB.UpdateCredits(Defaults.PokemonPrice, UserInfo.DbId, NewPokemon)
			if err != nil {
				return fmt.Sprintf(`{"intent":"error", "msg":"%s, DB error"}`, err.Error())
			}
			UserInfo.Credits -= Defaults.PokemonPrice
			fmt.Println(NewPokemon)
			return `{"intent":"success", "msg":"sucess"}`
		}
		return `{"intent":"msg", "msg":"Not enough credits"}`
	}

	return `{"intent":"error", "msg":"Not in a if"}`
}

func SaveInfo(msg string) (Defaults.UserInfo, error) { // some really ugly things happening here.
	test := strings.SplitAfter(string(msg), ":")
	test[11] = strings.Replace()
	poke := strings.Join(test[7:], " ")
	test[7] = `"null"}`
	final := strings.Join(test[:8], " ")
	info := Defaults.UserInfo{}
	err := json.Unmarshal([]byte(string(final)), &info)
	if err != nil {
		return info, fmt.Errorf(`{"intent":"error", "msg":"%s, fail to Parse JSON5"}`, err.Error())
	}
	info.Session = xid.New().String()
	info.Info = poke[1:]
	//fmt.Println(poke[1:20])
	//	fmt.Println(info)
	return info, nil
}

func CompareSessions(session string) error {
	x := strings.Compare(string(session), string(UserInfo.Session))
	if x == 0 {
		return nil
	}
	return errors.New("wrong sessions")
}
