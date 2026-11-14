package notifications

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup) {
	notifications := api.Group("/me/notifications")
	notifications.GET("", listNotifications)
	notifications.PUT("/:notificationId/read", markRead)
	notifications.GET("/settings", getSettings)
	notifications.PATCH("/settings", updateSettings)
}

func listNotifications(c *gin.Context) {
	c.JSON(http.StatusOK, mockNotifications())
}

func markRead(c *gin.Context) {
	c.JSON(http.StatusOK, mockReadNotification())
}

func getSettings(c *gin.Context) {
	c.JSON(http.StatusOK, mockSettings(FrequencyImmediate))
}

func updateSettings(c *gin.Context) {
	c.JSON(http.StatusOK, mockSettings(FrequencyDaily))
}
