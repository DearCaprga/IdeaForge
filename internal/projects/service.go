package projects

import (
	"context"

	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
)

type Service struct{}

func (Service) List(ctx context.Context, query CatalogQuery) (scaffold.Page[ProjectSummary], error) {
	return scaffold.Page[ProjectSummary]{}, scaffold.ErrNotImplemented
}

func (Service) ListPublishedByAuthor(ctx context.Context, authorID string, query scaffold.PageQuery) (scaffold.Page[ProjectSummary], error) {
	return scaffold.Page[ProjectSummary]{}, scaffold.ErrNotImplemented
}

// Get uses an empty viewer ID for guest access.
func (Service) Get(ctx context.Context, viewerID, projectID string) (Project, error) {
	return Project{}, scaffold.ErrNotImplemented
}

func (Service) GetCatalogOptions(ctx context.Context) (CatalogOptions, error) {
	return CatalogOptions{}, scaffold.ErrNotImplemented
}

func (Service) GetHome(ctx context.Context) (Home, error) {
	return Home{}, scaffold.ErrNotImplemented
}

func (Service) Suggest(ctx context.Context, query SuggestionQuery) ([]SearchSuggestion, error) {
	return nil, scaffold.ErrNotImplemented
}

func (Service) CreateDraft(ctx context.Context, authorID string, input DraftInput) (Project, error) {
	return Project{}, scaffold.ErrNotImplemented
}

func (Service) UpdateDraft(ctx context.Context, authorID, projectID string, input DraftInput) (Project, error) {
	return Project{}, scaffold.ErrNotImplemented
}

func (Service) DeleteDraft(ctx context.Context, authorID, projectID string) error {
	return scaffold.ErrNotImplemented
}

func (Service) Preview(ctx context.Context, authorID, projectID string) (PublicationPreview, error) {
	return PublicationPreview{}, scaffold.ErrNotImplemented
}

func (Service) Publish(ctx context.Context, authorID, projectID string) (Project, error) {
	return Project{}, scaffold.ErrNotImplemented
}

func (Service) ListMine(ctx context.Context, userID string, query MyProjectsQuery) (scaffold.Page[MyProject], error) {
	return scaffold.Page[MyProject]{}, scaffold.ErrNotImplemented
}

func (Service) AddFavorite(ctx context.Context, userID, projectID string) (Favorite, error) {
	return Favorite{}, scaffold.ErrNotImplemented
}

func (Service) RemoveFavorite(ctx context.Context, userID, projectID string) error {
	return scaffold.ErrNotImplemented
}

func (Service) GetWork(ctx context.Context, userID, projectID string) (ProjectWork, error) {
	return ProjectWork{}, scaffold.ErrNotImplemented
}

func (Service) StartWork(ctx context.Context, userID, projectID string) (ProjectWork, error) {
	return ProjectWork{}, scaffold.ErrNotImplemented
}

func (Service) UpdateWork(ctx context.Context, userID, projectID string, input UpdateWorkInput) (ProjectWork, error) {
	return ProjectWork{}, scaffold.ErrNotImplemented
}

func (Service) UpdateStage(ctx context.Context, userID, projectID, stageID string, input UpdateStageInput) (ProjectWork, error) {
	return ProjectWork{}, scaffold.ErrNotImplemented
}
