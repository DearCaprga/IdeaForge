package discussions

import (
	"context"

	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
)

type Repository interface {
	List(ctx context.Context, target Target, query ListQuery) (scaffold.Page[Comment], error)
	ListRecent(ctx context.Context, query ListQuery) (scaffold.Page[Comment], error)
	ListReplies(ctx context.Context, parentID string, query ListQuery) (scaffold.Page[Comment], error)
	FindByID(ctx context.Context, commentID string) (Comment, error)
	Create(ctx context.Context, comment Comment) error
	Update(ctx context.Context, comment Comment) error
	Delete(ctx context.Context, commentID string) error
	SetUseful(ctx context.Context, commentID string, useful bool) error
	AddReaction(ctx context.Context, reaction Reaction) error
	RemoveReaction(ctx context.Context, userID, commentID string, kind ReactionKind) error
}
