package discussions

import (
	"context"

	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
)

type Service struct{}

func (Service) List(ctx context.Context, target Target, query ListQuery) (scaffold.Page[Comment], error) {
	return scaffold.Page[Comment]{}, scaffold.ErrNotImplemented
}

func (Service) ListRecent(ctx context.Context, query ListQuery) (scaffold.Page[Comment], error) {
	return scaffold.Page[Comment]{}, scaffold.ErrNotImplemented
}

func (Service) ListReplies(ctx context.Context, parentID string, query ListQuery) (scaffold.Page[Comment], error) {
	return scaffold.Page[Comment]{}, scaffold.ErrNotImplemented
}

func (Service) Create(ctx context.Context, authorID string, target Target, input CommentInput) (Comment, error) {
	return Comment{}, scaffold.ErrNotImplemented
}

func (Service) Reply(ctx context.Context, authorID, parentID string, input CommentInput) (Comment, error) {
	return Comment{}, scaffold.ErrNotImplemented
}

func (Service) Edit(ctx context.Context, authorID, commentID string, input CommentInput) (Comment, error) {
	return Comment{}, scaffold.ErrNotImplemented
}

func (Service) Delete(ctx context.Context, actorID, commentID string) error {
	return scaffold.ErrNotImplemented
}

func (Service) AddReaction(ctx context.Context, userID, commentID string, kind ReactionKind) (Reaction, error) {
	return Reaction{}, scaffold.ErrNotImplemented
}

func (Service) RemoveReaction(ctx context.Context, userID, commentID string, kind ReactionKind) error {
	return scaffold.ErrNotImplemented
}

func (Service) SetUseful(ctx context.Context, actorID, commentID string, input SetUsefulInput) (Comment, error) {
	return Comment{}, scaffold.ErrNotImplemented
}
