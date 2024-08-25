package routes

import (
	"net/http"

	"github.com/dsc-sgu/shawty/internal/database"
	"github.com/dsc-sgu/shawty/internal/server/dto"
	"github.com/dsc-sgu/shawty/internal/server/routes/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Redirect(c *gin.Context) {
	linkName := c.Param("name")
	if len(linkName) == 0 {
		c.Status(http.StatusBadRequest)
		return
	}

	var query dto.RedirectMetadata
	if err := c.ShouldBind(&query); err != nil {
		c.Status(http.StatusUnprocessableEntity)
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

	id, err := uuid.NewV7()
	if err != nil {
		common.InternalError(c)
		return
	}

	visit := database.Visit{
		Id:     id,
		LinkId: sl.Id,
		Tag:    query.Tag,
		Host:   c.RemoteIP(),
	}
	if err := database.C.SaveVisit(c, visit); err != nil {
		common.InternalError(c)
		return
	}

	c.Redirect(http.StatusMovedPermanently, sl.Target)
}
