package dto

type CreateFormData struct {
	// Optional name of the link. Will be displayed
	// in the resulting URL.
	Name string `form:"name,omitempty,max=256"`
	// Target link to redirect to.
	Target string `form:"target,min=3,max=512"`
}

type CreateFormErrors struct {
	Name   string
	Target string
}

// Are there any errors?
func (e CreateFormErrors) Any() bool {
	return len(e.Name) != 0 || len(e.Target) != 0
}

type CreateLinkForm struct {
	Data   CreateFormData
	Errors CreateFormErrors
}
