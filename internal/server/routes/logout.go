package routes

import (
	"net/http"

	"github.com/dsc-sgu/shawty/internal/config"
	authdto "github.com/dsc-sgu/shawty/internal/server/dto/auth"
	"github.com/dsc-sgu/shawty/internal/server/html/render"
	authtempls "github.com/dsc-sgu/shawty/internal/server/html/templs/auth"
	"github.com/gin-gonic/gin"
)

func GetLogout(c *gin.Context) {
	c.SetCookie("session", "", 0, "", "", config.C.Ssl, true)
	r := render.New(c, authtempls.AuthForm(authdto.AuthForm{}))
	c.Render(http.StatusOK, r)
}
