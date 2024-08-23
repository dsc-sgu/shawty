package routes

import (
	"net/http"

	"github.com/dsc-sgu/shawty/internal/server/html/render"
	"github.com/dsc-sgu/shawty/internal/server/html/templates"
	"github.com/dsc-sgu/shawty/internal/util"
	"github.com/gin-gonic/gin"
)

func internalError(c *gin.Context) {
	c.Header("HX-Retarget", "main")
	c.Header("HX-Reswap", "innerHTML")
	r := render.New(c, templates.Error(util.InternalErrorText))
	c.Render(http.StatusOK, r)
}

func linkNotFound(c *gin.Context) {
	c.Header("HX-Retarget", "main")
	c.Header("HX-Reswap", "innerHTML")
	r := render.New(c, templates.Error(util.LinkNotFoundText))
	c.Render(http.StatusOK, r)
}
