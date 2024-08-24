package routes

import (
	"net/http"

	"github.com/dsc-sgu/shawty/internal/config"
	"github.com/dsc-sgu/shawty/internal/database"
	"github.com/dsc-sgu/shawty/internal/random"
	"github.com/dsc-sgu/shawty/internal/server/dto"
	"github.com/dsc-sgu/shawty/internal/server/html/render"
	"github.com/dsc-sgu/shawty/internal/server/html/templates"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetCreate(c *gin.Context) {
	form := dto.CreateForm{WithSecret: len(config.C.SharedSecret) != 0}
	r := render.New(c, templates.CreateForm(form))
	c.Render(http.StatusOK, r)
}

func PostCreate(c *gin.Context) {
	form := dto.CreateForm{
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

	if len(form.Data.Name) == 0 {
	loop:
		for {
			select {
			case <-c.Done():
				return
			default:
				form.Data.Name = random.RandSeq(10)
				if taken, err := database.C.IsNameTaken(c, form.Data.Name); err != nil {
					internalError(c)
					return
				} else if !taken {
					break loop
				}
			}
		}
	} else {
		if taken, err := database.C.IsNameTaken(c, form.Data.Name); err != nil {
			internalError(c)
			return
		} else if taken {
			form.Errors.Name = "a link with this name already exists"
		}
	}

	if form.Errors.Any() {
		r := render.New(c, templates.CreateForm(form))
		c.Render(http.StatusOK, r)
		return
	}

	id, err := uuid.NewV7()
	if err != nil {
		internalError(c)
		return
	}

	sl := database.ShortenedLink{
		Id:          id,
		Name:        form.Data.Name,
		Target:      form.Data.Target,
		CreatedFrom: "web_ui",
	}
	if err := database.C.SaveLink(c, sl); err != nil {
		internalError(c)
		return
	}

	r := render.New(
		c,
		templates.CreateResult(config.C.Ssl, config.C.Domain, form.Data.Name),
	)
	c.Render(http.StatusOK, r)
}
