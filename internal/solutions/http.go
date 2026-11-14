package solutions

import (
	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup) {
	api.GET("/projects/:projectId/solutions", scaffold.NotImplemented)
	api.POST("/projects/:projectId/solutions", scaffold.NotImplemented)
	api.GET("/users/:userId/solutions", scaffold.NotImplemented)
	api.GET("/solutions", scaffold.NotImplemented)

	item := api.Group("/solutions/:solutionId")
	item.GET("", scaffold.NotImplemented)
	item.PATCH("", scaffold.NotImplemented)
	item.GET("/revisions", scaffold.NotImplemented)
	item.PUT("/useful", scaffold.NotImplemented)
	item.DELETE("/useful", scaffold.NotImplemented)
	item.GET("/reviews", scaffold.NotImplemented)
	item.PUT("/reviews/me", scaffold.NotImplemented)
}
