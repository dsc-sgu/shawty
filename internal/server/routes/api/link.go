package apiroutes

import (
	"fmt"
	"net/http"

	"github.com/dsc-sgu/shawty/internal/config"
	"github.com/dsc-sgu/shawty/internal/database"
	"github.com/dsc-sgu/shawty/internal/models"
	"github.com/dsc-sgu/shawty/internal/random"
	apidto "github.com/dsc-sgu/shawty/internal/server/dto/api"
	"github.com/dsc-sgu/shawty/internal/server/routes/api/common"
	"github.com/dsc-sgu/shawty/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PostLink(c *gin.Context) {
	var body apidto.LinkCreateData
	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusUnprocessableEntity, apidto.Error{Error: err})
		return
	}

	form := apidto.LinkCreate{Data: body}
	form.ValidateTarget()

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
			common.BadRequest(c, form.Errors)
			return
		}

		if taken, err := database.C.IsNameTaken(c, form.Data.Name); err != nil {
			common.InternalError(c)
			return
		} else if taken {
			form.Errors.Name = util.LinkAlreadyExistsText
		}
	}

	id, err := uuid.NewV7()
	if err != nil {
		common.InternalError(c)
		return
	}

	link := models.Link{
		Id:          id,
		Name:        form.Data.Name,
		Target:      form.Data.Target,
		CreatedFrom: util.CreatedFromRestApi,
	}
	if err := database.C.SaveLink(c, link); err != nil {
		common.InternalError(c)
		return
	}

	c.Header("Location", fmt.Sprintf("/api/links/%v", link.Id))
	c.Status(http.StatusCreated)
}

func GetLinks(c *gin.Context) {
	var query apidto.LinkFetchQuery
	if err := c.ShouldBind(&query); err != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	if query.Page <= 0 {
		query.Page = 1 // if unset or garbage
	}

	links, err := database.C.GetLinksWithVisits(
		c,
		query.Page-1,
		config.C.Pagination.LinksPerPage,
	)
	if err != nil {
		common.InternalError(c)
		return
	}

	c.JSON(http.StatusOK, links)
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

	c.Status(http.StatusOK)
}
