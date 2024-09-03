package routes

import (
	"net/http"

	"github.com/dsc-sgu/shawty/internal/config"
	"github.com/dsc-sgu/shawty/internal/log"
	authdto "github.com/dsc-sgu/shawty/internal/server/dto/auth"
	"github.com/dsc-sgu/shawty/internal/server/html/render"
	"github.com/dsc-sgu/shawty/internal/server/html/templs"
	authtempls "github.com/dsc-sgu/shawty/internal/server/html/templs/auth"
	"github.com/dsc-sgu/shawty/internal/server/routes/common"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func PostAuth(c *gin.Context) {
	var data authdto.AuthFormData
	if err := c.ShouldBind(&data); err != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	form := authdto.AuthForm{Data: data}

	if form.Data.Secret != config.C.SharedSecret {
		log.S.Debugw("Failed authentication attempt", "host", c.Request.Host)
		form.Errors.Secret = "incorrect secret"
		r := render.New(c, authtempls.AuthForm(form))
		c.Render(http.StatusOK, r)
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	signed, err := token.SignedString([]byte(config.C.JwtSecret))
	if err != nil {
		log.S.Debugw("Failed to parse the session token", "error", err)
		common.InternalError(c)
		return
	}

	c.SetCookie("session", signed, 0, "", "", config.C.Ssl, true)
	r := render.New(c, templs.Home())
	c.Render(http.StatusOK, r)
}
