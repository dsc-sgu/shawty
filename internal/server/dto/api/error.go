package apidto

type Error struct {
	Error  error `json:"error"`
	Detail any   `json:"detail,omitempty"`
}
