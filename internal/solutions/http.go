package solutions

import (
	"net/http"

	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup) {
	api.GET("/projects/:projectId/solutions", mockListSolutions)
	api.POST("/projects/:projectId/solutions", mockPublishSolution)
	api.GET("/users/:userId/solutions", mockListSolutions)
	api.GET("/solutions", mockListSolutions)

	item := api.Group("/solutions/:solutionId")
	item.GET("", mockGetSolution)
	item.PATCH("", mockUpdateSolution)
	item.GET("/revisions", mockListRevisions)
	item.PUT("/useful", mockAddUsefulMark)
	item.DELETE("/useful", mockRemoveUsefulMark)
	item.GET("/reviews", mockListReviews)
	item.PUT("/reviews/me", mockSubmitReview)
}

func mockListSolutions(c *gin.Context) {
	c.JSON(http.StatusOK, scaffold.Page[Solution]{Items: []Solution{mockSolution()}, Total: 1, Limit: 20})
}

func mockGetSolution(c *gin.Context) {
	c.JSON(http.StatusOK, mockSolution())
}

func mockPublishSolution(c *gin.Context) {
	c.JSON(http.StatusCreated, mockSolution())
}

func mockUpdateSolution(c *gin.Context) {
	c.JSON(http.StatusOK, mockSolution())
}

func mockListRevisions(c *gin.Context) {
	solution := mockSolution()
	revision := Revision{
		ID: "revision-1", SolutionID: solution.ID, EditorID: solution.Author.ID,
		Content: solution.Content, PublishedAt: solution.CreatedAt,
	}
	c.JSON(http.StatusOK, scaffold.Page[Revision]{Items: []Revision{revision}, Total: 1, Limit: 20})
}

func mockAddUsefulMark(c *gin.Context) {
	c.JSON(http.StatusOK, UsefulMark{SolutionID: "solution-1", UserID: "user-1", CreatedAt: samplePublishedAt})
}

func mockRemoveUsefulMark(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func mockListReviews(c *gin.Context) {
	c.JSON(http.StatusOK, scaffold.Page[Review]{Items: []Review{mockReview()}, Total: 1, Limit: 20})
}

func mockSubmitReview(c *gin.Context) {
	c.JSON(http.StatusOK, mockReview())
}
