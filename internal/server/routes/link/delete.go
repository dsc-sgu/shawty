package linkroutes

import (
	"net/http"

	"github.com/dsc-sgu/shawty/internal/database"
	linkdto "github.com/dsc-sgu/shawty/internal/server/dto/link"
	"github.com/dsc-sgu/shawty/internal/server/html/render"
	linktempls "github.com/dsc-sgu/shawty/internal/server/html/templs/link"
	"github.com/dsc-sgu/shawty/internal/server/routes/common"

	"github.com/gin-gonic/gin"
)

func DeleteLink(c *gin.Context) {
	var query linkdto.DeleteQuery
	if err := c.ShouldBind(&query); err != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	lv, exists, err := database.C.FindLinkByName(c, query.Name)
	if err != nil {
		common.InternalError(c)
		return
	}
	if !exists {
		common.LinkNotFound(c)
		return
	}

	if err := database.C.DeleteLink(c, lv.Id); err != nil {
		common.InternalError(c)
		return
	}

	r := render.New(c, linktempls.LinkRow(lv, -1, true))
	c.Render(http.StatusOK, r)
}