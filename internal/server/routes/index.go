package routes

import (
	"net/http"

	"github.com/dsc-sgu/shawty/internal/config"
	"github.com/dsc-sgu/shawty/internal/server/auth"
	"github.com/dsc-sgu/shawty/internal/server/html/render"
	"github.com/dsc-sgu/shawty/internal/server/html/templs"
	"github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {
	sessionCookie, err := c.Cookie("session")
	if len(config.C.SharedSecret) != 0 &&
		(err != nil || auth.CheckSession(sessionCookie) != nil) {
		r := render.New(c, templs.IndexPage(false))
		c.Render(http.StatusOK, r)
		return
	}

	r := render.New(c, templs.IndexPage(true))
	c.Render(http.StatusOK, r)
}
