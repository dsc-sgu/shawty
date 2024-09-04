package webdto

type Auth struct {
	Data   AuthData
	Errors AuthErrors
}

type AuthData struct {
	Secret string `form:"secret,min=1"`
}

type AuthErrors struct {
	Secret string
}
