package notifications

import (
	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup) {
	notifications := api.Group("/me/notifications")
	notifications.GET("", scaffold.NotImplemented)
	notifications.PUT("/:notificationId/read", scaffold.NotImplemented)
	notifications.GET("/settings", scaffold.NotImplemented)
	notifications.PATCH("/settings", scaffold.NotImplemented)
}
