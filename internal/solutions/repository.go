package solutions

import (
	"context"

	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
)

type Repository interface {
	ListForProject(ctx context.Context, projectID string, query scaffold.PageQuery) (scaffold.Page[Solution], error)
	ListRecent(ctx context.Context, query scaffold.PageQuery) (scaffold.Page[Solution], error)
	ListByAuthor(ctx context.Context, authorID string, query scaffold.PageQuery) (scaffold.Page[Solution], error)
	FindByID(ctx context.Context, solutionID string) (Solution, error)
	CreateWithRevision(ctx context.Context, solution Solution, revision Revision) error
	UpdateWithRevision(ctx context.Context, solution Solution, revision Revision) error
	ListRevisions(ctx context.Context, solutionID string, query scaffold.PageQuery) (scaffold.Page[Revision], error)
}

type FeedbackRepository interface {
	AddUsefulMark(ctx context.Context, mark UsefulMark) error
	RemoveUsefulMark(ctx context.Context, userID, solutionID string) error
	ListReviews(ctx context.Context, solutionID string, query scaffold.PageQuery) (scaffold.Page[Review], error)
	SaveReview(ctx context.Context, review Review) error
}
