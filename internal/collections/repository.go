package collections

import (
	"context"

	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
)

type Repository interface {
	FindCollections(ctx context.Context, filter CollectionQuery, page scaffold.PageQuery) (scaffold.Page[Collection], error)
	FindCollection(ctx context.Context, collectionID string) (Collection, error)
	SaveCollection(ctx context.Context, collection Collection) error
	DeleteCollection(ctx context.Context, collectionID string) error
}
