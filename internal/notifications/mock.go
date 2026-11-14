package notifications

import (
	"time"

	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
)

func mockNotification() Notification {
	return Notification{
		ID:     "notification-1",
		UserID: "user-1",
		Kind:   KindNewSolution,
		Text:   "У проекта бота расписания появилось новое решение.",
		References: References{
			ProjectID:  "project-1",
			SolutionID: "solution-1",
		},
		CreatedAt: time.Date(2026, time.November, 3, 10, 0, 0, 0, time.UTC),
	}
}

func mockNotifications() ListResponse {
	return ListResponse{
		Page:        scaffold.Page[Notification]{Items: []Notification{mockNotification()}, Total: 1, Limit: 20},
		UnreadCount: 1,
	}
}

func mockReadNotification() Notification {
	notification := mockNotification()
	readAt := time.Date(2026, time.November, 3, 10, 5, 0, 0, time.UTC)
	notification.ReadAt = &readAt
	return notification
}

func mockSettings(frequency Frequency) Settings {
	return Settings{
		EnabledKinds: []Kind{
			KindCommentReply,
			KindNewSolution,
			KindTeamApplication,
			KindTeamDecision,
			KindProjectUpdate,
			KindMatchingProject,
			KindPublicationDecision,
			KindProjectReminder,
		},
		Frequency: frequency,
	}
}
