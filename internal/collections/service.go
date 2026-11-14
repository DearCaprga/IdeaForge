package collections

import (
	"context"

	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) List(ctx context.Context, filter CollectionQuery, page scaffold.PageQuery) (scaffold.Page[Collection], error) {
	return scaffold.Page[Collection]{}, scaffold.ErrNotImplemented
}

func (s *Service) Get(ctx context.Context, collectionID string) (Collection, error) {
	return Collection{}, scaffold.ErrNotImplemented
}

func (s *Service) Create(ctx context.Context, authorID string, input CollectionInput) (Collection, error) {
	return Collection{}, scaffold.ErrNotImplemented
}

func (s *Service) Update(ctx context.Context, actorID, collectionID string, input UpdateCollectionInput) (Collection, error) {
	return Collection{}, scaffold.ErrNotImplemented
}

func (s *Service) Delete(ctx context.Context, actorID, collectionID string) error {
	return scaffold.ErrNotImplemented
}
