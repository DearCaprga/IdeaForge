package teams

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup) {
	api.GET("/teams", listTeams)
	api.GET("/me/teams", listTeams)
	api.GET("/projects/:projectId/teams", listTeams)
	api.POST("/projects/:projectId/teams", createTeam)

	team := api.Group("/teams/:teamId")
	team.GET("", getTeam)
	team.PATCH("", updateTeam)
	team.GET("/applications", listApplications)
	team.POST("/applications", apply)
	team.PATCH("/applications/:applicationId", reviewApplication)
}

func listTeams(c *gin.Context) {
	c.JSON(http.StatusOK, mockTeams())
}

func createTeam(c *gin.Context) {
	c.JSON(http.StatusCreated, mockTeam())
}

func getTeam(c *gin.Context) {
	c.JSON(http.StatusOK, mockTeam())
}

func updateTeam(c *gin.Context) {
	c.JSON(http.StatusOK, mockTeam())
}

func listApplications(c *gin.Context) {
	c.JSON(http.StatusOK, mockApplications())
}

func apply(c *gin.Context) {
	c.JSON(http.StatusCreated, mockApplication())
}

func reviewApplication(c *gin.Context) {
	c.JSON(http.StatusOK, mockReviewedApplication())
}
