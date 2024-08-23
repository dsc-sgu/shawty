package routes

import (
	"net/http"

	"github.com/dsc-sgu/shawty/internal/database"
	"github.com/gin-gonic/gin"
)

func GetS(c *gin.Context) {
	linkName := c.Param("name")
	sl, exists, err := database.C.FindLinkByName(c, linkName)
	if err != nil {
		internalError(c)
		return
	}

	if !exists {
		linkNotFound(c)
		return
	}

	// TODO(evgenymng): update view statistics

	c.Redirect(http.StatusMovedPermanently, sl.Target)
}
