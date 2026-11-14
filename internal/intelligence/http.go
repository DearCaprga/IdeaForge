package intelligence

import (
	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup) {
	api.GET("/me/recommendations", scaffold.NotImplemented)
	tools := api.Group("/intelligence")
	tools.POST("/author-assistance", scaffold.NotImplemented)
	tools.POST("/similar-projects", scaffold.NotImplemented)
	tools.POST("/difficulty", scaffold.NotImplemented)
	tools.POST("/relevance", scaffold.NotImplemented)
}
