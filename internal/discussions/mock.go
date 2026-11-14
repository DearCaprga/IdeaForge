package discussions

import (
	"time"

	"github.com/DearCaprga/IdeaForge/internal/projects"
)

var sampleCreatedAt = time.Date(2026, time.October, 24, 12, 0, 0, 0, time.UTC)

func mockProjectComment() Comment {
	return Comment{
		ID: "comment-1", ProjectID: "project-1",
		Author:     projects.AuthorRef{ID: "user-2", DisplayName: "Мария Соколова"},
		Body:       "Можно ли использовать тестовые данные вместо расписания колледжа?",
		Reactions:  []ReactionCount{{Kind: ReactionLike, Count: 1}},
		ReplyCount: 1, CreatedAt: sampleCreatedAt, UpdatedAt: sampleCreatedAt,
	}
}

func mockProjectReply() Comment {
	return Comment{
		ID: "comment-2", ProjectID: "project-1", ParentID: "comment-1",
		Author:    projects.AuthorRef{ID: "user-1", DisplayName: "Алекс Морозов"},
		Body:      "Да, подготовьте два примера расписания с одним переносом и одной отменой.",
		Reactions: []ReactionCount{{Kind: ReactionThanks, Count: 1}},
		IsUseful:  true, CreatedAt: sampleCreatedAt, UpdatedAt: sampleCreatedAt,
	}
}

func mockSolutionComment() Comment {
	return Comment{
		ID: "comment-3", SolutionID: "solution-1",
		Author:    projects.AuthorRef{ID: "user-2", DisplayName: "Мария Соколова"},
		Body:      "Получилось запустить демонстрацию с тестовым расписанием.",
		Reactions: []ReactionCount{},
		CreatedAt: sampleCreatedAt, UpdatedAt: sampleCreatedAt,
	}
}
