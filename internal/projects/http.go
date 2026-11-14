package projects

import (
	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup) {
	api.GET("/home", scaffold.NotImplemented)
	api.GET("/search/suggestions", scaffold.NotImplemented)
	api.GET("/catalog/options", scaffold.NotImplemented)

	collection := api.Group("/projects")
	collection.GET("", scaffold.NotImplemented)
	collection.POST("", scaffold.NotImplemented)
	api.GET("/users/:userId/projects", scaffold.NotImplemented)

	item := collection.Group("/:projectId")
	item.GET("", scaffold.NotImplemented)
	item.PATCH("", scaffold.NotImplemented)
	item.DELETE("", scaffold.NotImplemented)
	item.GET("/preview", scaffold.NotImplemented)
	item.POST("/publish", scaffold.NotImplemented)
	item.PUT("/favorite", scaffold.NotImplemented)
	item.DELETE("/favorite", scaffold.NotImplemented)
	item.GET("/work", scaffold.NotImplemented)
	item.POST("/work", scaffold.NotImplemented)
	item.PATCH("/work", scaffold.NotImplemented)
	item.PATCH("/work/stages/:stageId", scaffold.NotImplemented)

	api.GET("/me/projects", scaffold.NotImplemented)
}
