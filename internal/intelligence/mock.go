package intelligence

import "github.com/DearCaprga/IdeaForge/internal/platform/scaffold"

var fixtureMeta = SuggestionMeta{
	Label:             "Предложено системой",
	SuggestedBySystem: true,
}

var fixtureEditableMeta = SuggestionMeta{
	Label:             "Предложено системой",
	SuggestedBySystem: true,
	Editable:          true,
}

var fixtureRecommendations = scaffold.Page[Recommendation]{
	Items: []Recommendation{
		{
			ProjectID:      "project-1",
			Title:          "Telegram-бот для уведомлений об изменении расписания",
			Reason:         "Подходит для практики Python и работы с API.",
			SuggestionMeta: fixtureMeta,
		},
	},
	Total: 1,
	Limit: 20,
}

var fixtureAuthorAssistance = AuthorAssistance{
	SuggestionMeta: fixtureEditableMeta,
	Goal:           "Уведомлять студентов об изменении расписания.",
	ExpectedResult: "Telegram-бот, репозиторий и инструкция по запуску.",
	Skills:         []string{"Python", "Работа с API"},
	Stages: []ProposedStage{
		{Title: "Требования", Description: "Определить источники данных и типы уведомлений."},
		{Title: "Разработка", Description: "Подготовить прототип бота."},
		{Title: "Проверка", Description: "Проверить уведомления о переносе и отмене занятий."},
	},
}

var fixtureSimilarProjects = SimilarProjects{
	SuggestionMeta: fixtureMeta,
	Items: []SimilarProject{
		{
			ProjectID:         "project-1",
			Title:             "Telegram-бот для уведомлений об изменении расписания",
			Reason:            "Совпадают аудитория и ожидаемый результат.",
			PossibleDuplicate: true,
		},
	},
}

var fixtureDifficulty = DifficultySuggestion{
	SuggestionMeta: fixtureEditableMeta,
	Difficulty:     "beginner",
	Rationale:      "Для первой версии достаточно простого диалога и одного источника расписания.",
}

var fixtureRelevance = RelevanceAssessment{
	SuggestionMeta: fixtureEditableMeta,
	Summary:        "Предварительная оценка проблемы требует подтверждения.",
	Signals: []RelevanceSignal{
		{Description: "Пользователи сообщают о задержке оповещений о переносе занятий."},
	},
}
