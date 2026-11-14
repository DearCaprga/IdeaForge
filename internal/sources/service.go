package sources

import (
	"context"

	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
)

type Service struct {
	repository Repository
	provider   Provider
	processor  CandidateProcessor
	publisher  ProjectPublisher
}

func NewService(repository Repository, provider Provider, processor CandidateProcessor, publisher ProjectPublisher) *Service {
	return &Service{repository: repository, provider: provider, processor: processor, publisher: publisher}
}

func (s *Service) ListSources(ctx context.Context, query SourceQuery) (scaffold.Page[Source], error) {
	return scaffold.Page[Source]{}, scaffold.ErrNotImplemented
}

func (s *Service) GetSource(ctx context.Context, sourceID string) (Source, error) {
	return Source{}, scaffold.ErrNotImplemented
}

func (s *Service) RegisterSource(ctx context.Context, moderatorID string, input SourceInput) (Source, error) {
	return Source{}, scaffold.ErrNotImplemented
}

func (s *Service) UpdateSource(ctx context.Context, moderatorID, sourceID string, input UpdateSourceInput) (Source, error) {
	return Source{}, scaffold.ErrNotImplemented
}

func (s *Service) VerifySource(ctx context.Context, moderatorID, sourceID string, input VerificationInput) (Source, error) {
	return Source{}, scaffold.ErrNotImplemented
}

func (s *Service) ListImports(ctx context.Context, query ImportQuery) (scaffold.Page[Import], error) {
	return scaffold.Page[Import]{}, scaffold.ErrNotImplemented
}

func (s *Service) StartImport(ctx context.Context, moderatorID, sourceID string, input StartImportInput) (Import, error) {
	return Import{}, scaffold.ErrNotImplemented
}

func (s *Service) GetImport(ctx context.Context, importID string) (Import, error) {
	return Import{}, scaffold.ErrNotImplemented
}

func (s *Service) ProcessImport(ctx context.Context, importID string) error {
	return scaffold.ErrNotImplemented
}

func (s *Service) CreateManualDraft(ctx context.Context, moderatorID, sourceID string, input CreateManualDraftInput) (Candidate, error) {
	return Candidate{}, scaffold.ErrNotImplemented
}

func (s *Service) ReviewImport(ctx context.Context, moderatorID, importID string, input ReviewImportInput) (Import, error) {
	return Import{}, scaffold.ErrNotImplemented
}

func (s *Service) PublishImport(ctx context.Context, moderatorID, importID string) (Import, error) {
	return Import{}, scaffold.ErrNotImplemented
}
