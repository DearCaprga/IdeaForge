package sources

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes attaches the mock endpoints to the API group.
func RegisterRoutes(api *gin.RouterGroup) {
	admin := api.Group("/admin")
	admin.GET("/sources", listSources)
	admin.POST("/sources", registerSource)
	admin.GET("/sources/:sourceId", getSource)
	admin.PATCH("/sources/:sourceId", updateSource)
	admin.PATCH("/sources/:sourceId/verification", verifySource)
	admin.GET("/sources/:sourceId/imports", listImports)
	admin.POST("/sources/:sourceId/imports", startImport)
	admin.POST("/sources/:sourceId/drafts", createManualDraft)
	admin.GET("/imports/:importId", getImport)
	admin.PATCH("/imports/:importId/review", reviewImport)
	admin.POST("/imports/:importId/publish", publishImport)
}

func listSources(c *gin.Context) {
	c.JSON(http.StatusOK, fixtureSources)
}

func registerSource(c *gin.Context) {
	c.JSON(http.StatusCreated, fixturePendingSource)
}

func getSource(c *gin.Context) {
	c.JSON(http.StatusOK, fixtureSource)
}

func updateSource(c *gin.Context) {
	c.JSON(http.StatusOK, fixtureSource)
}

func verifySource(c *gin.Context) {
	c.JSON(http.StatusOK, fixtureSource)
}

func listImports(c *gin.Context) {
	c.JSON(http.StatusOK, fixtureImports)
}

func startImport(c *gin.Context) {
	c.JSON(http.StatusAccepted, fixtureQueuedImport)
}

func getImport(c *gin.Context) {
	c.JSON(http.StatusOK, fixtureImport)
}

func createManualDraft(c *gin.Context) {
	c.JSON(http.StatusCreated, fixtureDraft)
}

func reviewImport(c *gin.Context) {
	c.JSON(http.StatusOK, fixtureApprovedImport)
}

func publishImport(c *gin.Context) {
	c.JSON(http.StatusOK, fixturePublishedImport)
}
