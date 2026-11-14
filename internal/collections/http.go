package collections

import (
	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup) {
	collections := api.Group("/collections")
	collections.GET("", scaffold.NotImplemented)
	collections.POST("", scaffold.NotImplemented)
	collections.GET("/:collectionId", scaffold.NotImplemented)
	collections.PATCH("/:collectionId", scaffold.NotImplemented)
	collections.DELETE("/:collectionId", scaffold.NotImplemented)
}
