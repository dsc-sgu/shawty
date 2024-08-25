package routes

import (
	"net/http"

	"github.com/dsc-sgu/shawty/internal/database"
	"github.com/dsc-sgu/shawty/internal/server/routes/common"
	"github.com/gin-gonic/gin"
)

func Redirect(c *gin.Context) {
	linkName := c.Param("name")
	if len(linkName) == 0 {
		c.Status(http.StatusBadRequest)
		return
	}
	sl, exists, err := database.C.FindLinkByName(c, linkName)
	if err != nil {
		common.InternalError(c)
		return
	}
	if !exists {
		common.LinkNotFound(c)
		return
	}

	// TODO(evgenymng): update view statistics

	c.Redirect(http.StatusMovedPermanently, sl.Target)
}
