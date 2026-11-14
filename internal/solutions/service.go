package solutions

import (
	"context"

	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
)

type Service struct{}

func (Service) ListForProject(ctx context.Context, projectID string, query scaffold.PageQuery) (scaffold.Page[Solution], error) {
	return scaffold.Page[Solution]{}, scaffold.ErrNotImplemented
}

func (Service) ListRecent(ctx context.Context, query scaffold.PageQuery) (scaffold.Page[Solution], error) {
	return scaffold.Page[Solution]{}, scaffold.ErrNotImplemented
}

func (Service) ListByAuthor(ctx context.Context, authorID string, query scaffold.PageQuery) (scaffold.Page[Solution], error) {
	return scaffold.Page[Solution]{}, scaffold.ErrNotImplemented
}

func (Service) Get(ctx context.Context, solutionID string) (Solution, error) {
	return Solution{}, scaffold.ErrNotImplemented
}

func (Service) Publish(ctx context.Context, authorID, projectID string, input CreateInput) (Solution, error) {
	return Solution{}, scaffold.ErrNotImplemented
}

func (Service) Update(ctx context.Context, editorID, solutionID string, input UpdateInput) (Solution, error) {
	return Solution{}, scaffold.ErrNotImplemented
}

func (Service) ListRevisions(ctx context.Context, solutionID string, query scaffold.PageQuery) (scaffold.Page[Revision], error) {
	return scaffold.Page[Revision]{}, scaffold.ErrNotImplemented
}

func (Service) AddUsefulMark(ctx context.Context, userID, solutionID string) (UsefulMark, error) {
	return UsefulMark{}, scaffold.ErrNotImplemented
}

func (Service) RemoveUsefulMark(ctx context.Context, userID, solutionID string) error {
	return scaffold.ErrNotImplemented
}

func (Service) ListReviews(ctx context.Context, solutionID string, query scaffold.PageQuery) (scaffold.Page[Review], error) {
	return scaffold.Page[Review]{}, scaffold.ErrNotImplemented
}

func (Service) SubmitReview(ctx context.Context, authorID, solutionID string, input ReviewInput) (Review, error) {
	return Review{}, scaffold.ErrNotImplemented
}
