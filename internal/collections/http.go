package collections

import (
	"net/http"

	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup) {
	collections := api.Group("/collections")
	collections.GET("", listCollections)
	collections.POST("", createCollection)
	collections.GET("/:collectionId", getCollection)
	collections.PATCH("/:collectionId", updateCollection)
	collections.DELETE("/:collectionId", deleteCollection)
}

func listCollections(c *gin.Context) {
	c.JSON(http.StatusOK, scaffold.Page[Collection]{Items: []Collection{demoCollection}, Total: 1, Limit: 20})
}

func createCollection(c *gin.Context) {
	c.JSON(http.StatusCreated, demoCollection)
}

func getCollection(c *gin.Context) {
	c.JSON(http.StatusOK, demoCollection)
}

func updateCollection(c *gin.Context) {
	c.JSON(http.StatusOK, demoCollection)
}

func deleteCollection(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
