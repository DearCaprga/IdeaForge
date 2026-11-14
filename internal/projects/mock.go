package projects

import "time"

var samplePublishedAt = time.Date(2026, time.October, 20, 12, 0, 0, 0, time.UTC)

func mockProject() Project {
	publishedAt := samplePublishedAt
	sourcePublishedAt := time.Date(2026, time.October, 10, 12, 0, 0, 0, time.UTC)
	return Project{
		ProjectSummary: ProjectSummary{
			ID:               "project-1",
			Title:            "Telegram-бот для уведомлений об изменении расписания",
			Summary:          "Бот сообщает студенту о переносе или отмене занятия.",
			Category:         CategoryDevelopment,
			Topic:            "Telegram-боты",
			Tags:             []string{"Python", "Telegram", "Расписание"},
			Difficulty:       DifficultyBeginner,
			EstimatedDaysMin: 14,
			EstimatedDaysMax: 21,
			Status:           StatusSolutionsPublished,
			Author:           AuthorRef{ID: "user-1", DisplayName: "Алекс Морозов"},
			Source: &Source{
				ID:                 "source-1",
				URL:                "https://example.org/cases/schedule",
				Type:               "hackathon",
				PublishedAt:        &sourcePublishedAt,
				Status:             SourceVerified,
				ProblemDescription: "Студенты узнают о переносе занятий с задержкой.",
			},
			SolutionCount: 1,
			TeamAvailable: true,
			PublishedAt:   &publishedAt,
		},
		Description:     "Сделайте бота, который отслеживает изменения в учебном расписании и сообщает о них подписчикам.",
		Problem:         "Студенты могут пропустить перенос или отмену занятия.",
		Relevance:       "Подписка на изменения избавит студентов от ручной проверки расписания.",
		Goal:            "Отправлять студентам уведомления об изменениях расписания по выбранной группе.",
		TargetAudience:  "Студенты",
		ExpectedResult:  "Работающий бот, краткая инструкция и репозиторий проекта.",
		SuccessCriteria: []string{"Уведомления приходят корректно", "Настройки пользователя сохраняются"},
		Skills:          []string{"Основы Python", "Работа с API", "Хранение данных"},
		Technologies:    []string{"Python", "Telegram Bot API"},
		Resources:       []Link{{Label: "Документация Telegram Bot API", URL: "https://core.telegram.org/bots/api"}},
		Attachments:     []Link{},
		Stages: []Stage{
			{ID: "stage-1", Title: "Требования", Description: "Опишите сценарии подписки и уведомлений.", Position: 1},
			{ID: "stage-2", Title: "Макет диалога", Description: "Подготовьте команды и сообщения бота.", Position: 2},
			{ID: "stage-3", Title: "Разработка", Description: "Подключите источник расписания и настройте отправку уведомлений.", Position: 3},
			{ID: "stage-4", Title: "Тестирование", Description: "Проверьте изменения расписания на тестовых данных.", Position: 4},
			{ID: "stage-5", Title: "Демонстрация", Description: "Подготовьте инструкцию и покажите уведомление.", Position: 5},
		},
		Development: &DevelopmentDetails{
			Direction: "automation", Languages: []string{"Python"}, Frameworks: []string{}, Alternatives: []Link{}, ResultType: "bot",
		},
		PublicationStatus: PublicationPublished,
		AcceptsSolutions:  true,
		CreatedAt:         publishedAt,
		UpdatedAt:         publishedAt,
	}
}

func mockDraft() Project {
	project := mockProject()
	project.ID = "project-draft-1"
	project.PublicationStatus = PublicationDraft
	project.Status = StatusOpen
	project.PublishedAt = nil
	project.Source = nil
	project.SolutionCount = 0
	return project
}

func mockWorkStarted() ProjectWork {
	return ProjectWork{
		ID: "work-1", ProjectID: "project-1", UserID: "user-1",
		Status: WorkInProgress, NextStageID: "stage-1",
		StartedAt: samplePublishedAt, LastActivityAt: samplePublishedAt,
		Stages: []StageProgress{{StageID: "stage-1"}, {StageID: "stage-2"}, {StageID: "stage-3"}, {StageID: "stage-4"}, {StageID: "stage-5"}},
	}
}

func mockWorkProgress() ProjectWork {
	work := mockWorkStarted()
	completedAt := samplePublishedAt
	work.Stages[0] = StageProgress{StageID: "stage-1", Completed: true, CompletedAt: &completedAt}
	work.ProgressPercent = 20
	work.NextStageID = "stage-2"
	return work
}

func mockWorkCompleted() ProjectWork {
	work := mockWorkStarted()
	completedAt := samplePublishedAt
	work.Stages = []StageProgress{
		{StageID: "stage-1", Completed: true, CompletedAt: &completedAt},
		{StageID: "stage-2", Completed: true, CompletedAt: &completedAt},
		{StageID: "stage-3", Completed: true, CompletedAt: &completedAt},
		{StageID: "stage-4", Completed: true, CompletedAt: &completedAt},
		{StageID: "stage-5", Completed: true, CompletedAt: &completedAt},
	}
	work.Status = WorkCompleted
	work.ProgressPercent = 100
	work.NextStageID = ""
	work.CompletedAt = &completedAt
	return work
}

func mockCatalogOptions() CatalogOptions {
	return CatalogOptions{
		Categories: []FilterOption{
			{Value: string(CategoryAcademic), Label: "Учебно-научные"},
			{Value: string(CategoryProductCase), Label: "Продуктовые кейсы"},
			{Value: string(CategoryDevelopment), Label: "Разработка"},
		},
		Difficulties: []FilterOption{
			{Value: string(DifficultyBeginner), Label: "Начальный"},
			{Value: string(DifficultyIntermediate), Label: "Средний"},
			{Value: string(DifficultyAdvanced), Label: "Продвинутый"},
			{Value: string(DifficultyExpert), Label: "Экспертный"},
		},
		Durations: []FilterOption{
			{Value: string(DurationWeek), Label: "До недели"},
			{Value: string(DurationMonth), Label: "До месяца"},
			{Value: string(DurationQuarter), Label: "1–3 месяца"},
			{Value: string(DurationLonger), Label: "Больше 3 месяцев"},
		},
		Statuses: []FilterOption{
			{Value: string(StatusOpen), Label: "Открыт"},
			{Value: string(StatusInProgress), Label: "В работе"},
			{Value: string(StatusSolutionsPublished), Label: "Решения опубликованы"},
		},
		SortOrders: []FilterOption{
			{Value: string(SortNewest), Label: "Новые"},
			{Value: string(SortPopular), Label: "Популярные"},
			{Value: string(SortDifficulty), Label: "По сложности"},
		},
		Topics:       []FilterOption{{Value: "Telegram-боты", Label: "Telegram-боты"}},
		Technologies: []FilterOption{{Value: "Python", Label: "Python"}, {Value: "Telegram Bot API", Label: "Telegram Bot API"}},
		Skills:       []FilterOption{{Value: "Работа с API", Label: "Работа с API"}},
		CategoryFields: []CategoryOptions{
			{Category: CategoryAcademic, Fields: map[string][]FilterOption{
				"subject": {{Value: "physics", Label: "Физика"}}, "subject_area": {{Value: "mechanics", Label: "Механика"}},
				"education_level": {{Value: "school", Label: "Школа"}}, "research_format": {{Value: "experiment", Label: "Эксперимент"}},
				"equipment": {{Value: "laboratory", Label: "Лабораторное оборудование"}},
			}},
			{Category: CategoryProductCase, Fields: map[string][]FilterOption{
				"industry": {{Value: "education", Label: "Образование"}}, "problem_source": {{Value: "interview", Label: "Интервью"}},
				"task_type": {{Value: "process_improvement", Label: "Улучшение процесса"}},
			}},
			{Category: CategoryDevelopment, Fields: map[string][]FilterOption{
				"direction": {{Value: "web", Label: "Веб"}, {Value: "mobile", Label: "Мобильная разработка"}, {Value: "ai", Label: "ИИ"}, {Value: "iot", Label: "Интернет вещей"}, {Value: "automation", Label: "Автоматизация"}},
				"languages": {{Value: "Python", Label: "Python"}}, "frameworks": {{Value: "aiogram", Label: "aiogram"}},
				"result_type": {{Value: "bot", Label: "Бот"}},
			}},
		},
	}
}
