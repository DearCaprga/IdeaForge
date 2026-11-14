package accounts

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	auth.POST("/register", register)
	auth.POST("/login", login)
	auth.POST("/refresh", refresh)
	auth.POST("/logout", logout)

	api.GET("/users/:userId", getPublicProfile)

	me := api.Group("/me")
	me.GET("", getPrivateProfile)
	me.PATCH("", updateProfile)
	me.GET("/settings", getSettings)
	me.PATCH("/settings", updateSettings)
}

func register(c *gin.Context) {
	c.JSON(http.StatusCreated, mockAuthResponse())
}

func login(c *gin.Context) {
	c.JSON(http.StatusOK, mockAuthResponse())
}

func refresh(c *gin.Context) {
	c.JSON(http.StatusOK, mockAuthResponse())
}

func logout(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func getPublicProfile(c *gin.Context) {
	c.JSON(http.StatusOK, mockProfile())
}

func getPrivateProfile(c *gin.Context) {
	c.JSON(http.StatusOK, mockPrivateProfile())
}

func updateProfile(c *gin.Context) {
	c.JSON(http.StatusOK, mockPrivateProfile())
}

func getSettings(c *gin.Context) {
	c.JSON(http.StatusOK, Settings{Theme: ThemeSystem})
}

func updateSettings(c *gin.Context) {
	c.JSON(http.StatusOK, Settings{Theme: ThemeDark})
}
