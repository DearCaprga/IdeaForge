package sources

import (
	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup) {
	admin := api.Group("/admin")
	admin.GET("/sources", scaffold.NotImplemented)
	admin.POST("/sources", scaffold.NotImplemented)
	admin.GET("/sources/:sourceId", scaffold.NotImplemented)
	admin.PATCH("/sources/:sourceId", scaffold.NotImplemented)
	admin.PATCH("/sources/:sourceId/verification", scaffold.NotImplemented)
	admin.GET("/sources/:sourceId/imports", scaffold.NotImplemented)
	admin.POST("/sources/:sourceId/imports", scaffold.NotImplemented)
	admin.POST("/sources/:sourceId/drafts", scaffold.NotImplemented)
	admin.GET("/imports/:importId", scaffold.NotImplemented)
	admin.PATCH("/imports/:importId/review", scaffold.NotImplemented)
	admin.POST("/imports/:importId/publish", scaffold.NotImplemented)
}
