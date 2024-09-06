package middleware

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/dsc-sgu/shawty/internal/config"
	"github.com/dsc-sgu/shawty/internal/log"
	"github.com/dsc-sgu/shawty/internal/server/auth"
	"github.com/dsc-sgu/shawty/internal/util"
	"github.com/gin-gonic/gin"
)

// Middleware to check for client authentication on
// every request.
func AuthMiddleware(protected []string) gin.HandlerFunc {
	compiled := util.Map(protected, func(re string) *regexp.Regexp {
		return regexp.MustCompile(re)
	})

	return func(c *gin.Context) {
		if len(config.C.SharedSecret) == 0 {
			c.Next()
			return
		}

		shouldProtect := util.Any(
			util.Map(compiled, func(re *regexp.Regexp) bool {
				return re.Match([]byte(c.Request.URL.Path))
			}),
		)

		if shouldProtect {
			session, err := c.Cookie("session")
			if err != nil {
				header := c.GetHeader("Authorization")
				session, err = processAuthHeader(header)
				if err != nil {
					authFailed(c, err)
					return
				}
			}

			if err = auth.CheckSession(session); err != nil {
				authFailed(c, err)
				return
			}
		}

		c.Next()
	}
}

func authFailed(c *gin.Context, err error) {
	log.S.Debugw("Failed authentication attempt", "error", err)
	c.Status(http.StatusForbidden)
	c.Abort()
}

func processAuthHeader(header string) (string, error) {
	if !strings.HasPrefix(header, "Bearer ") {
		return "", fmt.Errorf("authorization header has invalid format")
	}

	return strings.TrimPrefix(header, "Bearer "), nil
}
