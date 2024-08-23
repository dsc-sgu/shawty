package dto

type CreateFormData struct {
	// Shared secret, if specified in the config.
	Secret string `form:"secret,omitempty"`
	// Optional name of the link. Will be displayed
	// in the resulting URL.
	Name string `form:"name,omitempty,min=1,max=256"`
	// Target link to redirect to.
	Target string `form:"target,min=3,max=512"`
}

type CreateFormErrors struct {
	Secret string
	Name   string
	Target string
}

// Are there any errors?
func (e CreateFormErrors) Any() bool {
	return len(e.Secret) != 0 || len(e.Name) != 0 || len(e.Target) != 0
}

type CreateForm struct {
	WithSecret bool
	Data       CreateFormData
	Errors     CreateFormErrors
}
