package moderation

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes attaches the mock endpoints to the API group.
func RegisterRoutes(api *gin.RouterGroup) {
	api.POST("/reports", submitReport)

	review := api.Group("/admin/moderation")
	review.GET("/items", listQueue)
	review.GET("/items/:itemId", getQueueItem)
	review.POST("/items/:itemId/checks", checkItem)
	review.PATCH("/items/:itemId/decision", decideItem)
	review.GET("/reports", listReports)
	review.PATCH("/reports/:reportId/decision", decideReport)
	review.GET("/audit", listAuditEntries)

	catalog := api.Group("/admin/catalog")
	catalog.GET("/categories", listCategories)
	catalog.POST("/categories", createCategory)
	catalog.PATCH("/categories/:categoryId", updateCategory)
	catalog.DELETE("/categories/:categoryId", deleteCategory)
	catalog.GET("/tags", listTags)
	catalog.POST("/tags", createTag)
	catalog.PATCH("/tags/:tagId", updateTag)
	catalog.DELETE("/tags/:tagId", deleteTag)

	api.GET("/admin/analytics/usage", getUsageMetrics)
}

func listQueue(c *gin.Context) {
	c.JSON(http.StatusOK, fixtureQueue)
}

func getQueueItem(c *gin.Context) {
	c.JSON(http.StatusOK, fixtureQueueItem)
}

func checkItem(c *gin.Context) {
	c.JSON(http.StatusOK, fixtureChecks)
}

func decideItem(c *gin.Context) {
	item := fixtureQueueItem
	item.Status = ReviewApproved
	c.JSON(http.StatusOK, item)
}

func submitReport(c *gin.Context) {
	c.JSON(http.StatusCreated, fixtureReport)
}

func listReports(c *gin.Context) {
	c.JSON(http.StatusOK, fixtureReports)
}

func decideReport(c *gin.Context) {
	report := fixtureReport
	report.Status = ReportResolved
	c.JSON(http.StatusOK, report)
}

func listAuditEntries(c *gin.Context) {
	c.JSON(http.StatusOK, fixtureAudit)
}

func listCategories(c *gin.Context) {
	c.JSON(http.StatusOK, fixtureCategories)
}

func createCategory(c *gin.Context) {
	c.JSON(http.StatusCreated, fixtureCategory)
}

func updateCategory(c *gin.Context) {
	c.JSON(http.StatusOK, fixtureCategory)
}

func deleteCategory(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func listTags(c *gin.Context) {
	c.JSON(http.StatusOK, fixtureTags)
}

func createTag(c *gin.Context) {
	c.JSON(http.StatusCreated, fixtureTag)
}

func updateTag(c *gin.Context) {
	c.JSON(http.StatusOK, fixtureTag)
}

func deleteTag(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func getUsageMetrics(c *gin.Context) {
	c.JSON(http.StatusOK, fixtureUsage)
}
