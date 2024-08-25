package routes

import (
	"net/http"

	"github.com/dsc-sgu/shawty/internal/server/html/render"
	"github.com/dsc-sgu/shawty/internal/server/html/templs"
	"github.com/gin-gonic/gin"
)

func GetHome(c *gin.Context) {
	r := render.New(c, templs.Home())
	c.Render(http.StatusOK, r)
}
