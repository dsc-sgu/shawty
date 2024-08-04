package middleware

import (
	"net/http"
	"regexp"

	"github.com/dsc-sgu/atcc/internal/config"
	"github.com/dsc-sgu/atcc/internal/util"
	"github.com/gin-gonic/gin"
)

// Middleware to perform access control.
func AuthMiddleware(allowUri []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqPath := []byte(c.Request.URL.Path)
		allow := util.Any(util.Map(allowUri, func(au string) bool {
			res, _ := regexp.Match(au, reqPath)
			return res
		}))
		if len(config.C.SharedSecret) == 0 || allow {
			c.Next()
			return
		}

		key := c.GetHeader("X-Shared-Secret")
		if config.C.SharedSecret != key {
			c.Status(http.StatusForbidden)
			c.Abort()
			return
		}

		c.Next()
	}
}
