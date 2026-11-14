package scaffold

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NotImplemented is the temporary HTTP handler for scaffolded routes.
func NotImplemented(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, Problem{
		Code:    "not_implemented",
		Message: "Операция пока не реализована",
	})
}
