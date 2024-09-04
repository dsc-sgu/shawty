package webroutes

import (
	"net/http"

	"github.com/dsc-sgu/shawty/internal/config"
	"github.com/dsc-sgu/shawty/internal/database"
	"github.com/dsc-sgu/shawty/internal/models"
	"github.com/dsc-sgu/shawty/internal/random"
	webdto "github.com/dsc-sgu/shawty/internal/server/dto/webui"
	"github.com/dsc-sgu/shawty/internal/server/html/render"
	linktempls "github.com/dsc-sgu/shawty/internal/server/html/templs/link"
	"github.com/dsc-sgu/shawty/internal/server/routes/webui/common"
	"github.com/dsc-sgu/shawty/internal/util"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func NewLink(c *gin.Context) {
	form := webdto.LinkCreate{}
	r := render.New(c, linktempls.CreateForm(form))
	c.Render(http.StatusOK, r)
}

func PostLink(c *gin.Context) {
	var form webdto.LinkCreate
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
		form.ValidateName()
		if form.Errors.Any() {
			c.Status(http.StatusBadRequest)
			return
		}

		if taken, err := database.C.IsNameTaken(c, form.Data.Name); err != nil {
			common.InternalError(c)
			return
		} else if taken {
			form.Errors.Name = util.LinkAlreadyExistsText
		}
	}

	if form.Errors.Any() {
		r := render.New(c, linktempls.CreateForm(form))
		c.Render(http.StatusOK, r)
		return
	}

	id, err := uuid.NewV7()
	if err != nil {
		common.InternalError(c)
		return
	}

	sl := models.Link{
		Id:          id,
		Name:        form.Data.Name,
		Target:      form.Data.Target,
		CreatedFrom: util.CreatedFromWebUi,
	}
	if err := database.C.SaveLink(c, sl); err != nil {
		common.InternalError(c)
		return
	}

	r := render.New(
		c,
		linktempls.Result(config.C.Ssl, config.C.Domain, form.Data.Name),
	)
	c.Render(http.StatusOK, r)
}

func DeleteLink(c *gin.Context) {
	param := c.Param("id")
	id, err := uuid.Parse(param)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	lv, exists, err := database.C.FindLinkById(c, id)
	if err != nil {
		common.InternalError(c)
		return
	}
	if !exists {
		common.LinkNotFound(c)
		return
	}

	if err := database.C.DeleteLink(c, lv.Id); err != nil {
		common.InternalError(c)
		return
	}

	r := render.New(c, linktempls.LinkRow(lv, -1, true))
	c.Render(http.StatusOK, r)
}

func GetLinks(c *gin.Context) {
	var query webdto.LinkFetchQuery
	if err := c.ShouldBind(&query); err != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	params := webdto.LinkFetchParams{Query: query}
	if params.Query.Page == 0 {
		params.Query.Page = 1 // if unset
	}

	lv, err := database.C.GetLinksWithVisits(
		c,
		params.Query.Page-1,
		config.C.Pagination.LinksPerPage,
	)
	if err != nil {
		common.InternalError(c)
		return
	}

	params.Data = lv
	if params.Query.Page == 1 {
		r := render.New(c, linktempls.Browse(params))
		c.Render(http.StatusOK, r)
		return
	}

	r := render.New(c, linktempls.LinkRows(params))
	c.Render(http.StatusOK, r)
}
