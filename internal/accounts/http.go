package accounts

import (
	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	auth.POST("/register", scaffold.NotImplemented)
	auth.POST("/login", scaffold.NotImplemented)
	auth.POST("/refresh", scaffold.NotImplemented)
	auth.POST("/logout", scaffold.NotImplemented)

	api.GET("/users/:userId", scaffold.NotImplemented)

	me := api.Group("/me")
	me.GET("", scaffold.NotImplemented)
	me.PATCH("", scaffold.NotImplemented)
	me.GET("/settings", scaffold.NotImplemented)
	me.PATCH("/settings", scaffold.NotImplemented)
}
