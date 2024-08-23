package routes

import (
	"net/http"

	"github.com/dsc-sgu/shawty/internal/server/html/render"
	"github.com/dsc-sgu/shawty/internal/server/html/templates"
	"github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {
	r := render.New(c, templates.IndexPage())
	c.Render(http.StatusOK, r)
}
