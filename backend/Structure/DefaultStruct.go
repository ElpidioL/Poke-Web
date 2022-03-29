package structure

type Register struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Intent   string `json:"intent"`
	Msg      string `json:"msg"`
}

type IntentDefine struct {
	Intent string `json:"intent"`
	Msg    string `json:"msg"`
}
type TokenAcess struct {
	Intent string `json:"intent"`
	Colour string `json:"colour"`
	Email  string `json:"email"`
}
type Pokemon struct {
	Pokemon   string `json:"pokemon"`
	PokeId    int    `json:"pokeId"`
	Intent    string `json:"intent"`
	Handshake struct {
		Intent string `json:"intent"`
		Colour string `json:"colour"`
		Email  string `json:"email"`
	}
}

type UserInfo struct {
	Intent  string `json:"intent"`
	Credits string `json:"credits"`
	Info    string `json:"info"`
	Last    string `json:"last"`
	DbId    string `json:"dbId"`
	Session []byte
}
