package projects

import (
	"net/http"

	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup) {
	api.GET("/home", mockGetHome)
	api.GET("/search/suggestions", mockSuggest)
	api.GET("/catalog/options", mockGetCatalogOptions)

	collection := api.Group("/projects")
	collection.GET("", mockListProjects)
	collection.POST("", mockCreateDraft)
	api.GET("/users/:userId/projects", mockListProjects)

	item := collection.Group("/:projectId")
	item.GET("", mockGetProject)
	item.PATCH("", mockUpdateDraft)
	item.DELETE("", mockDeleteDraft)
	item.GET("/preview", mockPreview)
	item.POST("/publish", mockPublish)
	item.PUT("/favorite", mockAddFavorite)
	item.DELETE("/favorite", mockRemoveFavorite)
	item.GET("/work", mockGetWork)
	item.POST("/work", mockStartWork)
	item.PATCH("/work", mockUpdateWork)
	item.PATCH("/work/stages/:stageId", mockUpdateStage)

	api.GET("/me/projects", mockListMine)
}

func mockGetHome(c *gin.Context) {
	project := mockProject().ProjectSummary
	c.JSON(http.StatusOK, Home{
		Sections: []HomeSection{
			{Key: "popular", Title: "Популярные проекты", Projects: []ProjectSummary{project}},
			{Key: "beginner", Title: "Для начинающих", Projects: []ProjectSummary{project}},
			{Key: "real_cases", Title: "Реальные задачи", Projects: []ProjectSummary{project}},
			{Key: "looking_for_team", Title: "Ищут команду", Projects: []ProjectSummary{project}},
		},
		LatestSolutions: []RecentSolution{{
			ID: "solution-1", ProjectID: project.ID, Title: "Бот с подписками на группы", Author: project.Author,
		}},
	})
}

func mockSuggest(c *gin.Context) {
	project := mockProject()
	c.JSON(http.StatusOK, []SearchSuggestion{{Text: project.Title, ProjectID: project.ID}})
}

func mockGetCatalogOptions(c *gin.Context) {
	c.JSON(http.StatusOK, mockCatalogOptions())
}

func mockListProjects(c *gin.Context) {
	c.JSON(http.StatusOK, scaffold.Page[ProjectSummary]{
		Items: []ProjectSummary{mockProject().ProjectSummary}, Total: 1, Limit: 20, Offset: 0,
	})
}

func mockGetProject(c *gin.Context) {
	c.JSON(http.StatusOK, mockProject())
}

func mockCreateDraft(c *gin.Context) {
	c.JSON(http.StatusCreated, mockDraft())
}

func mockUpdateDraft(c *gin.Context) {
	c.JSON(http.StatusOK, mockDraft())
}

func mockDeleteDraft(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func mockPreview(c *gin.Context) {
	c.JSON(http.StatusOK, PublicationPreview{
		Project: mockDraft(), Warnings: []FieldWarning{}, PossibleDuplicates: []ProjectSummary{mockProject().ProjectSummary},
	})
}

func mockPublish(c *gin.Context) {
	draft := mockDraft()
	draft.PublicationStatus = PublicationUnderReview
	c.JSON(http.StatusAccepted, draft)
}

func mockAddFavorite(c *gin.Context) {
	c.JSON(http.StatusOK, Favorite{UserID: "user-1", ProjectID: "project-1", CreatedAt: samplePublishedAt})
}

func mockRemoveFavorite(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func mockGetWork(c *gin.Context) {
	c.JSON(http.StatusOK, mockWorkStarted())
}

func mockStartWork(c *gin.Context) {
	c.JSON(http.StatusCreated, mockWorkStarted())
}

func mockUpdateWork(c *gin.Context) {
	c.JSON(http.StatusOK, mockWorkCompleted())
}

func mockUpdateStage(c *gin.Context) {
	c.JSON(http.StatusOK, mockWorkProgress())
}

func mockListMine(c *gin.Context) {
	project := mockProject()
	work := mockWorkProgress()
	c.JSON(http.StatusOK, scaffold.Page[MyProject]{
		Items: []MyProject{{Project: project.ProjectSummary, PublicationStatus: project.PublicationStatus, Work: &work}},
		Total: 1, Limit: 20, Offset: 0,
	})
}
