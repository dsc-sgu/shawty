package linkroutes

import (
	"net/http"

	"github.com/dsc-sgu/shawty/internal/config"
	"github.com/dsc-sgu/shawty/internal/database"
	"github.com/dsc-sgu/shawty/internal/server/dto"
	"github.com/dsc-sgu/shawty/internal/server/html/render"
	"github.com/dsc-sgu/shawty/internal/server/html/templates"
	"github.com/dsc-sgu/shawty/internal/server/routes/common"
	"github.com/gin-gonic/gin"
)

func GetLinks(c *gin.Context) {
	var query dto.LinksViewQuery
	if err := c.ShouldBind(&query); err != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	params := dto.LinksParams{Query: query}
	if params.Query.Page == 0 {
		// if unset
		params.Query.Page = 1
	}

	lv, err := database.C.GetLinksVisits(
		c,
		params.Query.Page-1,
		config.C.Pagination.LinksPerPage,
	)
	if err != nil {
		common.InternalError(c)
		return
	}

	params.Data = lv
	if params.Query.Page == 1 {
		r := render.New(c, templates.Browse(params))
		c.Render(http.StatusOK, r)
		return
	}

	r := render.New(c, templates.LinkRows(params))
	c.Render(http.StatusOK, r)
}
