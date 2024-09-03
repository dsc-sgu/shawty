package authdto

type AuthForm struct {
	Data   AuthFormData
	Errors AuthFormErrors
}

type AuthFormData struct {
	Secret string `form:"secret,min=1"`
}

type AuthFormErrors struct {
	Secret string
}
