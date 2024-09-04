package apidto

type Auth struct {
	Data   AuthData
	Errors AuthErrors
}

type AuthData struct {
	Secret string `json:"secret"`
}

type AuthErrors struct {
	Secret string
}
