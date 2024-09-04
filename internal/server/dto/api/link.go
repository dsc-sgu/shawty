package apidto

import (
	"regexp"

	"github.com/dsc-sgu/shawty/internal/util"
)

type LinkFetchQuery struct {
	// Page number to query for.
	Page int `form:"page,min=1,omitempty"`
}

type LinkCreate struct {
	Data   LinkCreateData
	Errors LinkCreateErrors
}

type LinkCreateData struct {
	// Optional name of the link. Will be displayed
	// in the resulting URL.
	Name string `json:"name,omitempty"`
	// Target link to redirect to.
	Target string `json:"target"`
}

type LinkCreateErrors struct {
	Name   string `json:"name,omitempty"`
	Target string `json:"target,omitempty"`
}

// Are there any errors?
func (e LinkCreateErrors) Any() bool {
	return len(e.Name) != 0 || len(e.Target) != 0
}

var nameRegex = regexp.MustCompile(util.LinkNameRegex)

func (cl *LinkCreate) ValidateTarget() {
	if cl.Data.Target == "" {
		cl.Errors.Target = "target cannot be empty"
	}
}

func (cl *LinkCreate) ValidateName() {
	if !nameRegex.MatchString(cl.Data.Name) {
		cl.Errors.Name = `name must only contain letters a-z, digits 0-9, symbol "-", ` +
			`and be of length [1;256]`
	}
}
