package solutions

import (
	"time"

	"github.com/DearCaprga/IdeaForge/internal/projects"
)

var samplePublishedAt = time.Date(2026, time.October, 22, 12, 0, 0, 0, time.UTC)

func mockSolution() Solution {
	return Solution{
		Content: Content{
			Title:        "Бот с подписками на группы",
			Summary:      "Бот сравнивает версии расписания и уведомляет об изменениях.",
			Approach:     "Сохранять последнее расписание группы и проверять изменения по таймеру.",
			Technologies: []string{"Python", "Telegram Bot API"},
			Steps:        []string{"Подключили тестовые расписания", "Добавили сравнение версий", "Настроили подписки"},
			Result:       "Бот сообщает о переносах и отменах по выбранной группе.",
			Links:        []projects.Link{{Label: "Репозиторий (пример)", URL: "https://example.com/timetable-bot"}},
			Attachments:  []projects.Link{},
			Limitations:  "Демонстрация работает с тестовым расписанием.",
		},
		ID:          "solution-1",
		ProjectID:   "project-1",
		Author:      projects.AuthorRef{ID: "user-1", DisplayName: "Алекс Морозов"},
		UsefulCount: 1,
		CreatedAt:   samplePublishedAt,
		UpdatedAt:   samplePublishedAt,
	}
}

func mockReview() Review {
	return Review{
		ID: "review-1", SolutionID: "solution-1",
		Author:  projects.AuthorRef{ID: "user-2", DisplayName: "Мария Соколова"},
		Clarity: 5, Functionality: 4, Originality: 4, Reproducibility: 5,
		Comment: "Инструкцию легко повторить.", CreatedAt: samplePublishedAt,
	}
}
