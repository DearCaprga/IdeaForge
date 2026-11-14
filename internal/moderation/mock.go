package moderation

import (
	"time"

	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
)

var fixtureDate = time.Date(2026, time.November, 10, 12, 0, 0, 0, time.UTC)

var fixtureQueueItem = QueueItem{
	ID:          "review-1",
	Subject:     Subject{Kind: SubjectProject, ID: "project-draft-1"},
	Title:       "Telegram-бот для уведомлений об изменении расписания",
	SubmitterID: "user-1",
	Status:      ReviewPending,
	SubmittedAt: fixtureDate,
}

var fixtureQueue = scaffold.Page[QueueItem]{
	Items: []QueueItem{fixtureQueueItem},
	Total: 1,
	Limit: 20,
}

var fixtureChecks = []Check{
	{
		Kind:              CheckDuplicate,
		Status:            CheckPassed,
		Details:           "В демонстрационных данных совпадений нет.",
		RelatedProjectIDs: []string{},
		CheckedAt:         &fixtureDate,
	},
	{
		Kind:              CheckSource,
		Status:            CheckPending,
		Details:           "Ожидает проверки модератором.",
		RelatedProjectIDs: []string{},
	},
}

var fixtureReport = Report{
	ID:          "report-1",
	ReporterID:  "user-1",
	Subject:     Subject{Kind: SubjectSource, ID: "source-1"},
	Reason:      ReportOutdated,
	Description: "Нужно проверить дату публикации источника.",
	Status:      ReportOpen,
	CreatedAt:   fixtureDate,
}

var fixtureReports = scaffold.Page[Report]{
	Items: []Report{fixtureReport},
	Total: 1,
	Limit: 20,
}

var fixtureAudit = scaffold.Page[AuditEntry]{
	Items: []AuditEntry{
		{
			ID:          "audit-1",
			ModeratorID: "moderator-1",
			Action:      AuditSourceCheck,
			TargetID:    "source-1",
			Reason:      "Источник добавлен в очередь на проверку.",
			CreatedAt:   fixtureDate,
		},
	},
	Total: 1,
	Limit: 20,
}

var fixtureCategory = Category{
	ID: "development",
	CategoryInput: CategoryInput{
		Slug:        "development",
		Name:        "Разработка",
		Description: "Приложения, боты, устройства и автоматизация.",
	},
}

var fixtureCategories = scaffold.Page[Category]{
	Items: []Category{
		{ID: "academic", CategoryInput: CategoryInput{Slug: "academic", Name: "Учебно-научные"}},
		{ID: "product_case", CategoryInput: CategoryInput{Slug: "product_case", Name: "Продуктовые кейсы"}},
		fixtureCategory,
	},
	Total: 3,
	Limit: 20,
}

var fixtureTag = Tag{ID: "tag-1", TagInput: TagInput{Slug: "telegram", Name: "Telegram"}}

var fixtureTags = scaffold.Page[Tag]{
	Items: []Tag{fixtureTag},
	Total: 1,
	Limit: 20,
}

var fixtureUsage = UsageMetrics{
	UsageQuery: UsageQuery{
		From: time.Date(2026, time.November, 1, 0, 0, 0, 0, time.UTC),
		To:   time.Date(2026, time.November, 14, 0, 0, 0, 0, time.UTC),
	},
	PublishedProjects:  3,
	PublishedSolutions: 1,
	ActiveAuthors:      2,
}
