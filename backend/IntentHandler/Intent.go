package IntentHandler

/////seria interessante se eu fizesse as verificações corretas para ter certeza que as infos que vieram
///// possuem as propriedades certas do intent, visto que o intent fica no front end e pode ser alterado.

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	DB "github.com/ElpidioL/Poke-Web/PSDB"
	pass "github.com/ElpidioL/Poke-Web/PasswordHandler"
	Defaults "github.com/ElpidioL/Poke-Web/Structure"
	"github.com/rs/xid"
)

var UserInfo Defaults.UserInfo

func Intentions(choice []byte) string {
	Intents := Defaults.IntentDefine{}
	err := json.Unmarshal([]byte(string(choice)), &Intents)
	if err != nil {
		return fmt.Sprintf(`{"intent":"error", "msg":"%s, fail to Parse JSON"}`, err.Error())
	}

	if Intents.Intent == "colour" {

		LoginToken := Defaults.TokenAcess{}
		err = json.Unmarshal([]byte(string(choice)), &LoginToken)
		if err != nil {
			return fmt.Sprintf(`{"intent":"error", "msg":"%s, fail to Parse JSON"}`, err.Error())
		}
		info, err := DB.LoginUserToken(LoginToken.Colour, LoginToken.Email)
		if err != nil {
			return `{"intent":"error", "msg":"Fail to Colour"}`
		}

		vl, err := SaveInfo(info)
		if err != nil {
			return `{"intent":"error", "msg":"Fail to HandShake"}`
		}
		UserInfo = vl
		return info
	}

	if Intents.Intent == "register" || Intents.Intent == "login" {
		registerUser := Defaults.Register{}
		err = json.Unmarshal([]byte(string(choice)), &registerUser)
		if err != nil {
			return fmt.Sprintf(`{"intent":"error", "msg":"%s, fail to Parse JSON"}`, err.Error())
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

		NewPokemon := Defaults.Pokemon{}
		err = json.Unmarshal([]byte(string(choice)), &NewPokemon)
		if err != nil {
			return fmt.Sprintf(`{"intent":"error", "msg":"%s, fail to Parse JSON"}`, err.Error())
		}

		//fmt.Println(NewPokemon.Handshake.Colour)

		return `{"intent":"success", "msg":"sucess"}`
	}

	return `{"intent":"error", "msg":"Not in a if"}`
}

func SaveInfo(msg string) (Defaults.UserInfo, error) {
	info := Defaults.UserInfo{}
	err := json.Unmarshal([]byte(string(msg)), &info)
	info.Session = xid.New().Bytes()
	if err != nil {
		return info, fmt.Errorf(`{"intent":"error", "msg":"%s, fail to Parse JSON"}`, err.Error())
		//return info, fmt.Sprintf(`{"intent":"error", "msg":"%s, fail to Parse JSON"}`, err.Error())
	}
	return info, nil
}

func CompareSessions(session []byte) error {
	x := strings.Compare(string(session), string(UserInfo.Session))
	if x == 0 {
		return nil
	}
	return errors.New("wrong sessions")
}
