package intelligence

import (
	"context"

	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
)

type Service struct {
	readers   ReaderRepository
	projects  ProjectIndex
	assistant AuthorAssistant
}

func NewService(readers ReaderRepository, projects ProjectIndex, assistant AuthorAssistant) *Service {
	return &Service{readers: readers, projects: projects, assistant: assistant}
}

func (s *Service) ListRecommendations(ctx context.Context, userID string, query RecommendationQuery) (scaffold.Page[Recommendation], error) {
	return scaffold.Page[Recommendation]{}, scaffold.ErrNotImplemented
}

func (s *Service) AssistAuthor(ctx context.Context, brief ProjectBrief) (AuthorAssistance, error) {
	return AuthorAssistance{}, scaffold.ErrNotImplemented
}

func (s *Service) FindSimilarProjects(ctx context.Context, brief ProjectBrief) (SimilarProjects, error) {
	return SimilarProjects{}, scaffold.ErrNotImplemented
}

func (s *Service) SuggestDifficulty(ctx context.Context, brief ProjectBrief) (DifficultySuggestion, error) {
	return DifficultySuggestion{}, scaffold.ErrNotImplemented
}

func (s *Service) AssessRelevance(ctx context.Context, brief ProjectBrief) (RelevanceAssessment, error) {
	return RelevanceAssessment{}, scaffold.ErrNotImplemented
}
