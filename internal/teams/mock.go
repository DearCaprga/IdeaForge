package teams

import (
	"time"

	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
)

func mockTeam() Team {
	createdAt := time.Date(2026, time.November, 1, 10, 0, 0, 0, time.UTC)
	recruitingUntil := time.Date(2026, time.December, 1, 10, 0, 0, 0, time.UTC)

	return Team{
		ID:              "team-1",
		ProjectID:       "project-1",
		OwnerID:         "user-1",
		Name:            "Команда бота расписания",
		Description:     "Ищем разработчика для интеграции с расписанием.",
		Status:          StatusRecruiting,
		RecruitingUntil: &recruitingUntil,
		Openings: []Opening{{
			ID:        "opening-1",
			Role:      "Разработчик",
			Skills:    []string{"Python", "HTTP API"},
			Positions: 1,
		}},
		Members: []Member{{
			UserID:   "user-1",
			Role:     "Координатор",
			JoinedAt: createdAt,
		}},
		CreatedAt: createdAt,
	}
}

func mockTeams() scaffold.Page[Team] {
	return scaffold.Page[Team]{Items: []Team{mockTeam()}, Total: 1, Limit: 20}
}

func mockApplication() Application {
	return Application{
		ID:          "application-1",
		TeamID:      "team-1",
		ApplicantID: "user-2",
		OpeningID:   "opening-1",
		Message:     "Могу подключить API расписания.",
		Status:      ApplicationPending,
		CreatedAt:   time.Date(2026, time.November, 2, 10, 0, 0, 0, time.UTC),
	}
}

func mockApplications() scaffold.Page[Application] {
	return scaffold.Page[Application]{Items: []Application{mockApplication()}, Total: 1, Limit: 20}
}

func mockReviewedApplication() Application {
	application := mockApplication()
	reviewedAt := time.Date(2026, time.November, 3, 10, 0, 0, 0, time.UTC)
	application.Status = ApplicationAccepted
	application.ReviewedAt = &reviewedAt
	return application
}
