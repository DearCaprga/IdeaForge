package sources

import (
	"context"

	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
)

type Repository interface {
	ListSources(ctx context.Context, query SourceQuery) (scaffold.Page[Source], error)
	GetSource(ctx context.Context, sourceID string) (Source, error)
	InsertSource(ctx context.Context, input SourceInput) (Source, error)
	UpdateSource(ctx context.Context, sourceID string, input UpdateSourceInput) (Source, error)
	SaveVerification(ctx context.Context, moderatorID, sourceID string, input VerificationInput) (Source, error)
	InsertImport(ctx context.Context, importJob Import) (Import, error)
	ListImports(ctx context.Context, query ImportQuery) (scaffold.Page[Import], error)
	GetImport(ctx context.Context, importID string) (Import, error)
	UpdateImport(ctx context.Context, importJob Import) (Import, error)
}

type Provider interface {
	ListRecords(ctx context.Context, source Source, input StartImportInput) ([]SourceRecord, error)
}

type CandidateProcessor interface {
	ExtractProblem(ctx context.Context, source Source, record SourceRecord) (Candidate, error)
	RemovePersonalData(ctx context.Context, candidate Candidate) (Candidate, error)
	FindDuplicates(ctx context.Context, candidate Candidate) ([]string, error)
}

type ProjectPublisher interface {
	CreateDraft(ctx context.Context, moderatorID string, candidate Candidate) (projectID string, err error)
	PublishDraft(ctx context.Context, moderatorID, projectID string) error
}
