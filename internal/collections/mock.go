package collections

import "time"

var demoCollection = Collection{
	ID:          "collection-1",
	AuthorID:    "user-1",
	Title:       "Первые проекты по разработке",
	Description: "Подборка для тех, кто начинает знакомство с API и автоматизацией.",
	ProjectIDs:  []string{"project-1"},
	UpdatedAt:   time.Date(2026, time.January, 15, 12, 0, 0, 0, time.UTC),
}
