package intelligence

import (
	"context"

	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
)

type ReaderRepository interface {
	GetReaderContext(ctx context.Context, userID string) (ReaderContext, error)
}

type ProjectIndex interface {
	Recommend(ctx context.Context, reader ReaderContext, query RecommendationQuery) (scaffold.Page[Recommendation], error)
	FindSimilar(ctx context.Context, brief ProjectBrief) (SimilarProjects, error)
}

type AuthorAssistant interface {
	Assist(ctx context.Context, brief ProjectBrief) (AuthorAssistance, error)
	SuggestDifficulty(ctx context.Context, brief ProjectBrief) (DifficultySuggestion, error)
	AssessRelevance(ctx context.Context, brief ProjectBrief) (RelevanceAssessment, error)
}
