package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPing(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
