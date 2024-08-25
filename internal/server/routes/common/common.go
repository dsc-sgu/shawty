package common

import (
	"net/http"

	"github.com/dsc-sgu/shawty/internal/server/html/render"
	"github.com/dsc-sgu/shawty/internal/server/html/templs"
	"github.com/dsc-sgu/shawty/internal/util"
	"github.com/gin-gonic/gin"
)

func InternalError(c *gin.Context) {
	c.Header("HX-Retarget", "main")
	c.Header("HX-Reswap", "outerHTML")
	r := render.New(c, templs.Error(util.InternalErrorText))
	c.Render(http.StatusOK, r)
}

func LinkNotFound(c *gin.Context) {
	c.Header("HX-Retarget", "main")
	c.Header("HX-Reswap", "outerHTML")
	r := render.New(c, templs.Error(util.LinkNotFoundText))
	c.Render(http.StatusOK, r)
}
