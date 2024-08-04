package dto

type CreateForm struct {
	// Shared secret, if specified in the config.
	Secret string `form:"secret,omitempty"`
	// Optional name of the link. Will be displayed
	// in the resulting URL.
	Name string `form:"name,omitempty,min=1,max=256"`
	// Target link to redirect to.
	Target string `form:"target,min=3,max=512"`
}
