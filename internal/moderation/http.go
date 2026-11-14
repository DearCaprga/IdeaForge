package moderation

import (
	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup) {
	api.POST("/reports", scaffold.NotImplemented)

	review := api.Group("/admin/moderation")
	review.GET("/items", scaffold.NotImplemented)
	review.GET("/items/:itemId", scaffold.NotImplemented)
	review.POST("/items/:itemId/checks", scaffold.NotImplemented)
	review.PATCH("/items/:itemId/decision", scaffold.NotImplemented)
	review.GET("/reports", scaffold.NotImplemented)
	review.PATCH("/reports/:reportId/decision", scaffold.NotImplemented)
	review.GET("/audit", scaffold.NotImplemented)

	catalog := api.Group("/admin/catalog")
	catalog.GET("/categories", scaffold.NotImplemented)
	catalog.POST("/categories", scaffold.NotImplemented)
	catalog.PATCH("/categories/:categoryId", scaffold.NotImplemented)
	catalog.DELETE("/categories/:categoryId", scaffold.NotImplemented)
	catalog.GET("/tags", scaffold.NotImplemented)
	catalog.POST("/tags", scaffold.NotImplemented)
	catalog.PATCH("/tags/:tagId", scaffold.NotImplemented)
	catalog.DELETE("/tags/:tagId", scaffold.NotImplemented)

	api.GET("/admin/analytics/usage", scaffold.NotImplemented)
}
