package accounts

import "time"

func mockProfile() Profile {
	return Profile{
		ID:          "user-1",
		DisplayName: "Алекс Морозов",
		Bio:         "Изучаю разработку ботов и работу с API.",
		Interests:   []string{"Боты", "Образование"},
		Skills:      []string{"Python", "HTTP API"},
		Portfolio:   []PortfolioLink{},
		ProjectIDs:  []string{"project-1"},
		SolutionIDs: []string{"solution-1"},
		Stats:       ProfileStats{CompletedProjects: 1, Comments: 2, Teams: 1},
		Achievements: []Achievement{{
			Code:      "first_project",
			Title:     "Первый проект",
			AwardedAt: time.Date(2026, time.November, 1, 10, 0, 0, 0, time.UTC),
		}},
	}
}

func mockPrivateProfile() PrivateProfile {
	return PrivateProfile{
		Profile: mockProfile(),
		Email:   "alex@example.test",
		Roles:   []Role{RoleUser},
	}
}

func mockAuthResponse() AuthResponse {
	return AuthResponse{
		User:            mockPrivateProfile(),
		AccessToken:     "mock-access-token",
		RefreshToken:    "mock-refresh-token",
		AccessExpiresAt: time.Date(2099, time.January, 1, 12, 0, 0, 0, time.UTC),
	}
}
