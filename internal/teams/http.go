package teams

import (
	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup) {
	api.GET("/teams", scaffold.NotImplemented)
	api.GET("/me/teams", scaffold.NotImplemented)
	api.GET("/projects/:projectId/teams", scaffold.NotImplemented)
	api.POST("/projects/:projectId/teams", scaffold.NotImplemented)

	team := api.Group("/teams/:teamId")
	team.GET("", scaffold.NotImplemented)
	team.PATCH("", scaffold.NotImplemented)
	team.GET("/applications", scaffold.NotImplemented)
	team.POST("/applications", scaffold.NotImplemented)
	team.PATCH("/applications/:applicationId", scaffold.NotImplemented)
}
