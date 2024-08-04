package routes

import (
	"fmt"
	"net/http"

	"github.com/dsc-sgu/atcc/internal/config"
	"github.com/dsc-sgu/atcc/internal/database"
	"github.com/dsc-sgu/atcc/internal/log"
	"github.com/dsc-sgu/atcc/internal/random"
	"github.com/dsc-sgu/atcc/internal/server/dto"
	"github.com/dsc-sgu/atcc/internal/server/html/render"
	"github.com/dsc-sgu/atcc/internal/server/html/templates"
	"github.com/gin-gonic/gin"
)

func GetCreate(c *gin.Context) {
	withSecret := len(config.C.SharedSecret) != 0
	r := render.New(c, templates.CreateForm(withSecret, dto.CreateForm{}, nil))
	c.Render(http.StatusOK, r)
}

func PostCreate(c *gin.Context) {
	withSecret := len(config.C.SharedSecret) != 0

	var form dto.CreateForm
	if err := c.ShouldBind(&form); err != nil {
		r := render.New(c, templates.CreateForm(withSecret, form, err))
		c.Render(http.StatusUnprocessableEntity, r)
		return
	}

	log.S.Debugw("Request", "req", form)

	if len(config.C.SharedSecret) != 0 && form.Secret != config.C.SharedSecret {
		r := render.New(
			c,
			templates.CreateForm(
				withSecret,
				form,
				fmt.Errorf("Provide a secret"),
			),
		)
		c.Render(http.StatusUnprocessableEntity, r)
		return
	}

	if len(form.Name) == 0 {
	loop:
		for {
			select {
			case <-c.Done():
				return
			default:
				form.Name = random.RandSeq(10)
				if taken, err := database.C.IsNameTaken(c, form.Name); err != nil {
					r := render.New(c, templates.ErrorBox("Oops, something is very wrong..."))
					c.Render(http.StatusInternalServerError, r)
					return
				} else if !taken {
					break loop
				}
			}
		}
	} else {
		if taken, err := database.C.IsNameTaken(c, form.Name); err != nil {
			r := render.New(c, templates.ErrorBox("Oops, something is very wrong..."))
			c.Render(http.StatusInternalServerError, r)
			return
		} else if taken {
			r := render.New(
				c,
				templates.CreateForm(
					withSecret,
					form,
					fmt.Errorf("A link with this name already exists"),
				),
			)
			c.Render(http.StatusConflict, r)
			return
		}
	}

	sl := database.ShortenedLink{
		Name:        form.Name,
		Target:      form.Target,
		CreatedFrom: "web_ui",
	}
	if err := database.C.SaveLink(c, sl); err != nil {
		r := render.New(
			c,
			templates.ErrorBox("Oops, something is very wrong..."),
		)
		c.Render(http.StatusInternalServerError, r)
		return
	}

	r := render.New(c, templates.Result(config.C.Domain, form.Name))
	c.Render(http.StatusOK, r)
}
