// Package web provides the page structure for the IdeaForge scaffold.
package web

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed templates/*.html static/styles.css
var files embed.FS

type pageView struct {
	Title      string
	ResourceID string
}

type screen struct {
	Path     string
	Title    string
	Template string
	IDParam  string
}

var screens = []screen{
	{Path: "/", Title: "Главная", Template: "home.html"},
	{Path: "/catalog", Title: "Каталог проектов", Template: "catalog.html"},
	{Path: "/projects/new", Title: "Создать проект", Template: "project-form.html"},
	{Path: "/projects/:projectId", Title: "Проект", Template: "project.html", IDParam: "projectId"},
	{Path: "/projects/:projectId/edit", Title: "Редактировать проект", Template: "project-form.html", IDParam: "projectId"},
	{Path: "/projects/:projectId/preview", Title: "Предпросмотр проекта", Template: "project-preview.html", IDParam: "projectId"},
	{Path: "/projects/:projectId/work", Title: "Этапы работы", Template: "work.html", IDParam: "projectId"},
	{Path: "/projects/:projectId/solutions/new", Title: "Опубликовать решение", Template: "solution-form.html", IDParam: "projectId"},
	{Path: "/solutions/:solutionId", Title: "Решение", Template: "solution.html", IDParam: "solutionId"},
	{Path: "/solutions/:solutionId/edit", Title: "Редактировать решение", Template: "solution-form.html", IDParam: "solutionId"},
	{Path: "/login", Title: "Войти", Template: "login.html"},
	{Path: "/register", Title: "Регистрация", Template: "register.html"},
	{Path: "/me/projects", Title: "Мои проекты", Template: "my-projects.html"},
	{Path: "/me", Title: "Мой профиль", Template: "profile.html"},
	{Path: "/users/:userId", Title: "Профиль участника", Template: "profile.html", IDParam: "userId"},
	{Path: "/me/settings", Title: "Настройки", Template: "settings.html"},
	{Path: "/me/notifications", Title: "Уведомления", Template: "notifications.html"},
	{Path: "/me/teams", Title: "Мои команды", Template: "my-teams.html"},
	{Path: "/projects/:projectId/teams", Title: "Поиск команды", Template: "teams.html", IDParam: "projectId"},
	{Path: "/teams/:teamId", Title: "Команда", Template: "team.html", IDParam: "teamId"},
	{Path: "/community", Title: "Сообщество", Template: "community.html"},
	{Path: "/collections/:collectionId", Title: "Подборка проектов", Template: "collection.html", IDParam: "collectionId"},
	{Path: "/admin", Title: "Модерация и управление", Template: "admin.html"},
}

func RegisterRoutes(router *gin.Engine) {
	router.SetHTMLTemplate(template.Must(template.ParseFS(files, "templates/*.html")))
	router.StaticFileFS("/assets/styles.css", "static/styles.css", http.FS(files))
	for _, view := range screens {
		router.GET(view.Path, renderScreen(view))
	}
}

func renderScreen(view screen) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, view.Template, pageView{
			Title:      view.Title,
			ResourceID: c.Param(view.IDParam),
		})
	}
}
