package routes

import (
	"net/http"

	"github.com/dsc-sgu/shawty/internal/config"
	"github.com/dsc-sgu/shawty/internal/database"
	"github.com/dsc-sgu/shawty/internal/server/dto"
	"github.com/dsc-sgu/shawty/internal/server/html/render"
	"github.com/dsc-sgu/shawty/internal/server/html/templates"

	"github.com/gin-gonic/gin"
)

func GetDelete(c *gin.Context) {
	form := dto.DeleteForm{WithSecret: len(config.C.SharedSecret) != 0}
	r := render.New(c, templates.DeleteForm(form))
	c.Render(http.StatusOK, r)
}

func PostDelete(c *gin.Context) {
	form := dto.DeleteForm{
		WithSecret: len(config.C.SharedSecret) != 0,
	}
	if err := c.ShouldBind(&form); err != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	if len(config.C.SharedSecret) != 0 {
		if len(form.Data.Secret) == 0 {
			form.Errors.Secret = "provide a secret"
		} else if form.Data.Secret != config.C.SharedSecret {
			form.Errors.Secret = "the secret is incorrect"
		}
	}

	sl, exists, err := database.C.FindLinkByName(c, form.Data.Name)
	if err != nil {
		internalError(c)
		return
	} else if !exists {
		form.Errors.Name = "a link with this name doesn't exist"
	}

	if form.Errors.Any() {
		r := render.New(c, templates.DeleteForm(form))
		c.Render(http.StatusOK, r)
		return
	}

	if err := database.C.DeleteLink(c, sl.Id); err != nil {
		internalError(c)
		return
	}

	r := render.New(c, templates.DeleteResult())
	c.Render(http.StatusOK, r)
}
