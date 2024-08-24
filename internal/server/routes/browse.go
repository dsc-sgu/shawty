package routes

import (
	"net/http"

	"github.com/dsc-sgu/shawty/internal/config"
	"github.com/dsc-sgu/shawty/internal/database"
	"github.com/dsc-sgu/shawty/internal/server/dto"
	"github.com/dsc-sgu/shawty/internal/server/html/render"
	"github.com/dsc-sgu/shawty/internal/server/html/templates"
	"github.com/gin-gonic/gin"
)

func GetBrowse(c *gin.Context) {
	var params dto.LinksViewParams
	if err := c.ShouldBind(&params); err != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	if params.PageNumber == 0 {
		params.PageNumber = 1
	}

	slv, err := database.C.GetLinksVisits(
		c,
		params.PageNumber,
		config.C.Pagination.LinksPerPage,
	)
	if err != nil {
		internalError(c)
		return
	}

	r := render.New(c, templates.LinksView(params, slv))
	c.Render(http.StatusOK, r)
}
