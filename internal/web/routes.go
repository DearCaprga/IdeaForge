// Package web registers placeholder pages and their navigation.
package web

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed templates/*.html
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
	{Path: "/", Title: "Главная", Template: "page.html"},
	{Path: "/search", Title: "Поиск", Template: "page.html"},
	{Path: "/catalog", Title: "Каталог проектов", Template: "catalog.html"},
	{Path: "/projects/new", Title: "Создать проект", Template: "project-form.html"},
	{Path: "/projects/:projectId", Title: "Проект", Template: "project.html", IDParam: "projectId"},
	{Path: "/projects/:projectId/edit", Title: "Редактировать проект", Template: "project-form.html", IDParam: "projectId"},
	{Path: "/projects/:projectId/preview", Title: "Предпросмотр проекта", Template: "page.html"},
	{Path: "/projects/:projectId/work", Title: "Этапы работы", Template: "work.html", IDParam: "projectId"},
	{Path: "/projects/:projectId/discussion", Title: "Обсуждение проекта", Template: "page.html"},
	{Path: "/projects/:projectId/solutions", Title: "Решения проекта", Template: "solutions.html", IDParam: "projectId"},
	{Path: "/projects/:projectId/solutions/new", Title: "Опубликовать решение", Template: "page.html"},
	{Path: "/solutions/:solutionId", Title: "Решение", Template: "solution.html", IDParam: "solutionId"},
	{Path: "/solutions/:solutionId/edit", Title: "Редактировать решение", Template: "page.html"},
	{Path: "/login", Title: "Войти", Template: "login.html"},
	{Path: "/register", Title: "Регистрация", Template: "register.html"},
	{Path: "/me/projects", Title: "Мои проекты", Template: "page.html"},
	{Path: "/me", Title: "Мой профиль", Template: "account.html"},
	{Path: "/users/:userId", Title: "Профиль участника", Template: "page.html"},
	{Path: "/me/settings", Title: "Настройки", Template: "page.html"},
	{Path: "/me/notifications", Title: "Уведомления", Template: "page.html"},
	{Path: "/me/teams", Title: "Мои команды", Template: "page.html"},
	{Path: "/projects/:projectId/teams", Title: "Поиск команды", Template: "teams.html"},
	{Path: "/teams/:teamId", Title: "Команда", Template: "page.html"},
	{Path: "/community", Title: "Сообщество", Template: "community.html"},
	{Path: "/collections/:collectionId", Title: "Подборка проектов", Template: "page.html"},
	{Path: "/admin", Title: "Модерация и управление", Template: "page.html"},
}

func RegisterRoutes(router *gin.Engine) {
	router.SetHTMLTemplate(template.Must(template.ParseFS(files, "templates/*.html")))
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
