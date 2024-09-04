package webdto

import (
	"regexp"

	"github.com/dsc-sgu/shawty/internal/models"
	"github.com/dsc-sgu/shawty/internal/util"
)

type LinkFetchQuery struct {
	// Page number to query for.
	Page int `form:"page,min=1,omitempty"`
}

type LinkFetchParams struct {
	// Query parameters.
	Query LinkFetchQuery
	// Data to display in the view.
	Data []models.LinkWithVisits
}

type LinkCreateData struct {
	// Optional name of the link. Will be displayed
	// in the resulting URL.
	Name string `form:"name,omitempty,max=256"`
	// Target link to redirect to.
	Target string `form:"target,min=3,max=512"`
}

type LinkCreateErrors struct {
	Name   string
	Target string
}

// Are there any errors?
func (e LinkCreateErrors) Any() bool {
	return len(e.Name) != 0 || len(e.Target) != 0
}

type LinkCreate struct {
	Data   LinkCreateData
	Errors LinkCreateErrors
}

var nameRegex = regexp.MustCompile(util.LinkNameRegex)

func (cl *LinkCreate) ValidateName() {
	if !nameRegex.MatchString(cl.Data.Name) {
		cl.Errors.Name = `name must only contain letters a-z, digits 0-9, symbol "-", ` +
			`and be of length [1;256]`
	}
}
