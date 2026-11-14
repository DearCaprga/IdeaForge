package discussions

import (
	"net/http"

	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup) {
	api.GET("/projects/:projectId/comments", mockListProjectComments)
	api.POST("/projects/:projectId/comments", mockCreateProjectComment)
	api.GET("/solutions/:solutionId/comments", mockListSolutionComments)
	api.POST("/solutions/:solutionId/comments", mockCreateSolutionComment)
	api.GET("/comments", mockListProjectComments)

	item := api.Group("/comments/:commentId")
	item.PATCH("", mockEditComment)
	item.DELETE("", mockDeleteComment)
	item.GET("/replies", mockListReplies)
	item.POST("/replies", mockReply)
	item.PUT("/reactions/:kind", mockAddReaction)
	item.DELETE("/reactions/:kind", mockRemoveReaction)
	item.PUT("/useful", mockSetUseful)
}

func mockListProjectComments(c *gin.Context) {
	c.JSON(http.StatusOK, scaffold.Page[Comment]{Items: []Comment{mockProjectComment()}, Total: 1, Limit: 20})
}

func mockListSolutionComments(c *gin.Context) {
	c.JSON(http.StatusOK, scaffold.Page[Comment]{Items: []Comment{mockSolutionComment()}, Total: 1, Limit: 20})
}

func mockCreateProjectComment(c *gin.Context) {
	c.JSON(http.StatusCreated, mockProjectComment())
}

func mockCreateSolutionComment(c *gin.Context) {
	c.JSON(http.StatusCreated, mockSolutionComment())
}

func mockEditComment(c *gin.Context) {
	c.JSON(http.StatusOK, mockProjectComment())
}

func mockDeleteComment(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func mockListReplies(c *gin.Context) {
	c.JSON(http.StatusOK, scaffold.Page[Comment]{Items: []Comment{mockProjectReply()}, Total: 1, Limit: 20})
}

func mockReply(c *gin.Context) {
	c.JSON(http.StatusCreated, mockProjectReply())
}

func mockAddReaction(c *gin.Context) {
	c.JSON(http.StatusOK, Reaction{CommentID: "comment-1", UserID: "user-1", Kind: ReactionLike, CreatedAt: sampleCreatedAt})
}

func mockRemoveReaction(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func mockSetUseful(c *gin.Context) {
	c.JSON(http.StatusOK, mockProjectReply())
}
