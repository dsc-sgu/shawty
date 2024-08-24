package dto

type DeleteFormData struct {
	// Shared secret, if specified in the config.
	Secret string `form:"secret,omitempty"`
	// Optional name of the link. Will be displayed
	// in the resulting URL.
	Name string `form:"name,omitempty,min=1,max=256"`
}

type DeleteFormErrors struct {
	Secret string
	Name   string
}

// Are there any errors?
func (e DeleteFormErrors) Any() bool {
	return len(e.Secret) != 0 || len(e.Name) != 0
}

type DeleteForm struct {
	WithSecret bool
	Data       CreateFormData
	Errors     CreateFormErrors
}
