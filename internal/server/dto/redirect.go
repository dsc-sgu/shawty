package dto

type RedirectMetadata struct {
	// A tag that you can add to gather more statistics
	// about each redirect that is being performed.
	Tag string `form:"t,omitempty"`
}
