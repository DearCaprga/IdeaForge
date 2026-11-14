package teams

import (
	"context"

	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
)

type Service interface {
	List(ctx context.Context, query ListQuery) (scaffold.Page[Team], error)
	ListMyTeams(ctx context.Context, userID string, query scaffold.PageQuery) (scaffold.Page[Team], error)
	Get(ctx context.Context, teamID string) (Team, error)
	Create(ctx context.Context, userID string, projectID string, request CreateRequest) (Team, error)
	Update(ctx context.Context, userID string, teamID string, request UpdateRequest) (Team, error)
	ListApplications(ctx context.Context, userID string, teamID string, query ApplicationsQuery) (scaffold.Page[Application], error)
	Apply(ctx context.Context, userID string, teamID string, request ApplyRequest) (Application, error)
	ReviewApplication(ctx context.Context, userID string, teamID string, applicationID string, request ReviewRequest) (Application, error)
}

type StubService struct{}

var _ Service = StubService{}

func (StubService) List(ctx context.Context, query ListQuery) (scaffold.Page[Team], error) {
	return scaffold.Page[Team]{}, scaffold.ErrNotImplemented
}

func (StubService) ListMyTeams(ctx context.Context, userID string, query scaffold.PageQuery) (scaffold.Page[Team], error) {
	return scaffold.Page[Team]{}, scaffold.ErrNotImplemented
}

func (StubService) Get(ctx context.Context, teamID string) (Team, error) {
	return Team{}, scaffold.ErrNotImplemented
}

func (StubService) Create(ctx context.Context, userID string, projectID string, request CreateRequest) (Team, error) {
	return Team{}, scaffold.ErrNotImplemented
}

func (StubService) Update(ctx context.Context, userID string, teamID string, request UpdateRequest) (Team, error) {
	return Team{}, scaffold.ErrNotImplemented
}

func (StubService) ListApplications(ctx context.Context, userID string, teamID string, query ApplicationsQuery) (scaffold.Page[Application], error) {
	return scaffold.Page[Application]{}, scaffold.ErrNotImplemented
}

func (StubService) Apply(ctx context.Context, userID string, teamID string, request ApplyRequest) (Application, error) {
	return Application{}, scaffold.ErrNotImplemented
}

func (StubService) ReviewApplication(ctx context.Context, userID string, teamID string, applicationID string, request ReviewRequest) (Application, error) {
	return Application{}, scaffold.ErrNotImplemented
}
