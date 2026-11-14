// Package app assembles the HTTP server and feature routes.
package app

import (
	"net/http"

	"github.com/DearCaprga/IdeaForge/internal/accounts"
	"github.com/DearCaprga/IdeaForge/internal/collections"
	"github.com/DearCaprga/IdeaForge/internal/discussions"
	"github.com/DearCaprga/IdeaForge/internal/intelligence"
	"github.com/DearCaprga/IdeaForge/internal/moderation"
	"github.com/DearCaprga/IdeaForge/internal/notifications"
	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
	"github.com/DearCaprga/IdeaForge/internal/projects"
	"github.com/DearCaprga/IdeaForge/internal/solutions"
	"github.com/DearCaprga/IdeaForge/internal/sources"
	"github.com/DearCaprga/IdeaForge/internal/teams"
	"github.com/DearCaprga/IdeaForge/internal/web"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.HandleMethodNotAllowed = true
	router.GET("/health", health)
	web.RegisterRoutes(router)

	api := router.Group("/api/v1")
	api.Use(markMockResponse)
	accounts.RegisterRoutes(api)
	projects.RegisterRoutes(api)
	solutions.RegisterRoutes(api)
	discussions.RegisterRoutes(api)
	teams.RegisterRoutes(api)
	notifications.RegisterRoutes(api)
	collections.RegisterRoutes(api)
	moderation.RegisterRoutes(api)
	sources.RegisterRoutes(api)
	intelligence.RegisterRoutes(api)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, scaffold.Problem{Code: "not_found", Message: "Маршрут не найден"})
	})
	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, scaffold.Problem{Code: "method_not_allowed", Message: "Метод недоступен"})
	})
	return router
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "okeyyy"})
}

func markMockResponse(c *gin.Context) {
	c.Header("X-IdeaForge-Mode", "mock")
	c.Next()
}
