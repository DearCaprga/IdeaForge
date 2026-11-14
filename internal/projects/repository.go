package projects

import (
	"context"

	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
)

type Repository interface {
	ListPublished(ctx context.Context, query CatalogQuery) (scaffold.Page[ProjectSummary], error)
	ListPublishedByAuthor(ctx context.Context, authorID string, query scaffold.PageQuery) (scaffold.Page[ProjectSummary], error)
	ListMyProjects(ctx context.Context, userID string, query MyProjectsQuery) (scaffold.Page[MyProject], error)
	FindByID(ctx context.Context, projectID string) (Project, error)
	Create(ctx context.Context, project Project) error
	Update(ctx context.Context, project Project) error
	DeleteDraft(ctx context.Context, authorID, projectID string) error
	ListCatalogOptions(ctx context.Context) (CatalogOptions, error)
}

type FavoriteRepository interface {
	Add(ctx context.Context, favorite Favorite) error
	Remove(ctx context.Context, userID, projectID string) error
}

type WorkRepository interface {
	Find(ctx context.Context, userID, projectID string) (ProjectWork, error)
	Create(ctx context.Context, work ProjectWork) error
	Update(ctx context.Context, work ProjectWork) error
}
