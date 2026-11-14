package teams

import (
	"context"

	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
)

type Repository interface {
	List(ctx context.Context, query ListQuery) (scaffold.Page[Team], error)
	ListByMember(ctx context.Context, userID string, query scaffold.PageQuery) (scaffold.Page[Team], error)
	Find(ctx context.Context, teamID string) (Team, error)
	Create(ctx context.Context, team Team) error
	Update(ctx context.Context, teamID string, changes UpdateRequest) (Team, error)
}

type ApplicationRepository interface {
	List(ctx context.Context, teamID string, query ApplicationsQuery) (scaffold.Page[Application], error)
	Find(ctx context.Context, teamID string, applicationID string) (Application, error)
	Create(ctx context.Context, application Application) error
	// Review saves the application and adds the member in the same operation when member is present.
	Review(ctx context.Context, application Application, member *Member) error
}
