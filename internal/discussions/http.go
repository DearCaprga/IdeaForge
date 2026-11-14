package discussions

import (
	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup) {
	api.GET("/projects/:projectId/comments", scaffold.NotImplemented)
	api.POST("/projects/:projectId/comments", scaffold.NotImplemented)
	api.GET("/solutions/:solutionId/comments", scaffold.NotImplemented)
	api.POST("/solutions/:solutionId/comments", scaffold.NotImplemented)
	api.GET("/comments", scaffold.NotImplemented)

	item := api.Group("/comments/:commentId")
	item.PATCH("", scaffold.NotImplemented)
	item.DELETE("", scaffold.NotImplemented)
	item.GET("/replies", scaffold.NotImplemented)
	item.POST("/replies", scaffold.NotImplemented)
	item.PUT("/reactions/:kind", scaffold.NotImplemented)
	item.DELETE("/reactions/:kind", scaffold.NotImplemented)
	item.PUT("/useful", scaffold.NotImplemented)
}
