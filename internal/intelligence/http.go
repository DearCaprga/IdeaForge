package intelligence

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes attaches the mock endpoints to the API group.
func RegisterRoutes(api *gin.RouterGroup) {
	api.GET("/me/recommendations", listRecommendations)
	tools := api.Group("/intelligence")
	tools.POST("/author-assistance", assistAuthor)
	tools.POST("/similar-projects", findSimilarProjects)
	tools.POST("/difficulty", suggestDifficulty)
	tools.POST("/relevance", assessRelevance)
}

func listRecommendations(c *gin.Context) {
	c.JSON(http.StatusOK, fixtureRecommendations)
}

func assistAuthor(c *gin.Context) {
	c.JSON(http.StatusOK, fixtureAuthorAssistance)
}

func findSimilarProjects(c *gin.Context) {
	c.JSON(http.StatusOK, fixtureSimilarProjects)
}

func suggestDifficulty(c *gin.Context) {
	c.JSON(http.StatusOK, fixtureDifficulty)
}

func assessRelevance(c *gin.Context) {
	c.JSON(http.StatusOK, fixtureRelevance)
}
