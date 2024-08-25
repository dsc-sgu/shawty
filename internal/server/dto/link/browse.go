package linkdto

import "github.com/dsc-sgu/shawty/internal/database"

type ViewQuery struct {
	// Page number to query for.
	Page int `form:"page,min=1,omitempty"`
}

type ViewParams struct {
	// Query parameters.
	Query ViewQuery
	// Data to display in the view.
	Data []database.LinkWithVisits
}
