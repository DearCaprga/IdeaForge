package sources

import (
	"time"

	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
)

var fixtureTime = time.Date(2026, time.November, 10, 13, 0, 0, 0, time.UTC)

var fixtureSourceInput = SourceInput{
	Name:            "Демонстрационный архив учебных кейсов",
	Kind:            KindHackathon,
	URL:             "https://example.org/cases",
	PermissionURL:   "https://example.org/terms",
	PermissionNotes: "Пример разрешения на импорт.",
}

var fixturePendingSource = Source{
	ID:           "source-1",
	SourceInput:  fixtureSourceInput,
	Verification: VerificationPending,
	CreatedAt:    fixtureTime,
}

var fixtureSource = Source{
	ID:            "source-1",
	SourceInput:   fixtureSourceInput,
	Verification:  VerificationVerified,
	ImportAllowed: true,
	CreatedAt:     fixtureTime,
	CheckedAt:     &fixtureTime,
}

var fixtureSources = scaffold.Page[Source]{
	Items: []Source{fixtureSource},
	Total: 1,
	Limit: 20,
}

var fixtureCandidate = Candidate{
	ID:                  "candidate-1",
	Title:               "Telegram-бот для уведомлений об изменении расписания",
	Description:         "Сообщать студентам о переносах и отменах занятий.",
	CategoryID:          "development",
	SuggestedDifficulty: "beginner",
	Provenance: Provenance{
		SourceID:      "source-1",
		URL:           "https://example.org/cases/schedule",
		PublishedAt:   &fixtureTime,
		ProblemClaim:  "Студенты узнают о переносе занятий с задержкой.",
		VerifiedFacts: []string{},
	},
	PrivacyReviewed:     true,
	DuplicateProjectIDs: []string{},
}

var fixtureDraft = Candidate{
	ID:                  fixtureCandidate.ID,
	Title:               fixtureCandidate.Title,
	Description:         fixtureCandidate.Description,
	CategoryID:          fixtureCandidate.CategoryID,
	SuggestedDifficulty: fixtureCandidate.SuggestedDifficulty,
	Provenance:          fixtureCandidate.Provenance,
	PrivacyReviewed:     fixtureCandidate.PrivacyReviewed,
	DuplicateProjectIDs: fixtureCandidate.DuplicateProjectIDs,
	ProjectID:           "project-1",
}

var fixtureQueuedImport = Import{
	ID:          "import-1",
	SourceID:    "source-1",
	Status:      ImportQueued,
	RequestedBy: "moderator-1",
	CreatedAt:   fixtureTime,
	Candidates:  []Candidate{},
}

var fixtureImport = Import{
	ID:          "import-1",
	SourceID:    "source-1",
	Status:      ImportPendingReview,
	RequestedBy: "moderator-1",
	CreatedAt:   fixtureTime,
	Candidates:  []Candidate{fixtureCandidate},
}

var fixtureImports = scaffold.Page[Import]{
	Items: []Import{fixtureImport},
	Total: 1,
	Limit: 20,
}

var fixtureReview = ImportReview{
	ReviewImportInput: ReviewImportInput{
		Decision:    DecisionApprove,
		Reason:      "Происхождение кейса проверено.",
		AcceptedIDs: []string{"candidate-1"},
	},
	ReviewerID: "moderator-1",
	ReviewedAt: fixtureTime,
}

var fixtureApprovedImport = Import{
	ID:          "import-1",
	SourceID:    "source-1",
	Status:      ImportApproved,
	RequestedBy: "moderator-1",
	CreatedAt:   fixtureTime,
	Candidates:  []Candidate{fixtureCandidate},
	Review:      &fixtureReview,
}

var fixturePublishedImport = Import{
	ID:          "import-1",
	SourceID:    "source-1",
	Status:      ImportPublished,
	RequestedBy: "moderator-1",
	CreatedAt:   fixtureTime,
	Candidates:  []Candidate{fixtureDraft},
	Review:      &fixtureReview,
}
