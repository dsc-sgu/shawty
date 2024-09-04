package common

import (
	"fmt"
	"net/http"

	apidto "github.com/dsc-sgu/shawty/internal/server/dto/api"
	"github.com/dsc-sgu/shawty/internal/util"
	"github.com/gin-gonic/gin"
)

func InternalError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, apidto.Error{
		Error: fmt.Errorf(util.InternalErrorText),
	})
}

func LinkNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, apidto.Error{
		Error: fmt.Errorf(util.LinkNotFoundText),
	})
}

func BadRequest(c *gin.Context, detail any) {
	c.JSON(http.StatusBadRequest, apidto.Error{
		Error:  fmt.Errorf(util.BadRequestText),
		Detail: detail,
	})
}
