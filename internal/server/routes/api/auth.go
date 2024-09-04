package apiroutes

import (
	"fmt"
	"net/http"

	"github.com/dsc-sgu/shawty/internal/config"
	"github.com/dsc-sgu/shawty/internal/log"
	apidto "github.com/dsc-sgu/shawty/internal/server/dto/api"
	"github.com/dsc-sgu/shawty/internal/server/routes/api/common"
	"github.com/dsc-sgu/shawty/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func PostAuth(c *gin.Context) {
	var data apidto.AuthData
	if err := c.ShouldBind(&data); err != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	form := apidto.Auth{Data: data}

	if form.Data.Secret != config.C.SharedSecret {
		log.S.Debugw("Failed authentication attempt", "host", c.Request.Host)
		form.Errors.Secret = "incorrect secret"
		c.JSON(http.StatusUnauthorized, apidto.Error{
			Error:  fmt.Errorf(util.AuthFailedText),
			Detail: form.Errors,
		})
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	signed, err := token.SignedString([]byte(config.C.JwtSecret))
	if err != nil {
		log.S.Debugw("Failed to sign the session token", "error", err)
		common.InternalError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"session": signed})
}
