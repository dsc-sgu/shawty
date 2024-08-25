package dto

import "github.com/dsc-sgu/shawty/internal/database"

type LinksViewQuery struct {
	// Page number to query for.
	Page int `form:"page,min=1,omitempty"`
}

type LinksParams struct {
	// Query parameters.
	Query LinksViewQuery
	// Data to display in the view.
	Data []database.LinkWithVisits
}
