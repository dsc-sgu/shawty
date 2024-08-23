package dto

type ErrorData struct {
	// Message to display on the error page, if any.
	Message string `form:"msg,omitempty"`
}
