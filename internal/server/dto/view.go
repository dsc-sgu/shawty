package dto

type LinksViewParams struct {
	// Shared secret, if specified in the config.
	Secret string `form:"secret,omitempty"`
	// Page number to query for.
	PageNumber int `form:"page,min=1,omitempty"`
}
