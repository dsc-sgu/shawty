package linkroutes

import (
	"net/http"
	"regexp"

	"github.com/dsc-sgu/shawty/internal/config"
	"github.com/dsc-sgu/shawty/internal/database"
	"github.com/dsc-sgu/shawty/internal/random"
	"github.com/dsc-sgu/shawty/internal/server/dto"
	"github.com/dsc-sgu/shawty/internal/server/html/render"
	"github.com/dsc-sgu/shawty/internal/server/html/templates"
	"github.com/dsc-sgu/shawty/internal/server/routes/common"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func NewLink(c *gin.Context) {
	form := dto.CreateLinkForm{}
	r := render.New(c, templates.CreateForm(form))
	c.Render(http.StatusOK, r)
}

var nameRegex = regexp.MustCompile(`^[a-z0-9\-]{1,256}$`)

func PostLink(c *gin.Context) {
	var form dto.CreateLinkForm
	if err := c.ShouldBind(&form); err != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
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
					common.InternalError(c)
					return
				} else if !taken {
					break loop
				}
			}
		}
	} else {
		if !nameRegex.Match([]byte(form.Data.Name)) {
			c.Status(http.StatusBadRequest)
			return
		}

		if taken, err := database.C.IsNameTaken(c, form.Data.Name); err != nil {
			common.InternalError(c)
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
		common.InternalError(c)
		return
	}

	sl := database.Link{
		Id:          id,
		Name:        form.Data.Name,
		Target:      form.Data.Target,
		CreatedFrom: "web_ui",
	}
	if err := database.C.SaveLink(c, sl); err != nil {
		common.InternalError(c)
		return
	}

	r := render.New(
		c,
		templates.CreateResult(config.C.Ssl, config.C.Domain, form.Data.Name),
	)
	c.Render(http.StatusOK, r)
}
