package util

const (
	AppName    = "Shawty"
	AppVersion = "0.0.1"
)

const (
	InternalErrorText     = "oops, something is very wrong..."
	LinkNotFoundText      = "link was not found"
	LinkAlreadyExistsText = "a link with this name already exists"
	BadRequestText        = "data validation failed"
	AuthFailedText        = "failed authentication attempt"
)

const (
	LinkNameRegex = `^[a-z0-9\-]{1,256}$`
)

const (
	CreatedFromWebUi   = "web_ui"
	CreatedFromRestApi = "rest_api"
)
