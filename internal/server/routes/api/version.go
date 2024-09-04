package apiroutes

import (
	"net/http"

	"github.com/dsc-sgu/shawty/internal/util"
	"github.com/gin-gonic/gin"
)

func GetVersion(c *gin.Context) {
	c.String(http.StatusOK, "%s-%s", util.AppName, util.AppVersion)
}
