package notifications

import (
	"context"

	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
)

type Service interface {
	Notify(ctx context.Context, input NewNotification) error
	List(ctx context.Context, userID string, query ListQuery) (ListResponse, error)
	MarkRead(ctx context.Context, userID string, notificationID string) (Notification, error)
	GetSettings(ctx context.Context, userID string) (Settings, error)
	UpdateSettings(ctx context.Context, userID string, request UpdateSettingsRequest) (Settings, error)
}

type StubService struct{}

var _ Service = StubService{}

func (StubService) Notify(ctx context.Context, input NewNotification) error {
	return scaffold.ErrNotImplemented
}

func (StubService) List(ctx context.Context, userID string, query ListQuery) (ListResponse, error) {
	return ListResponse{}, scaffold.ErrNotImplemented
}

func (StubService) MarkRead(ctx context.Context, userID string, notificationID string) (Notification, error) {
	return Notification{}, scaffold.ErrNotImplemented
}

func (StubService) GetSettings(ctx context.Context, userID string) (Settings, error) {
	return Settings{}, scaffold.ErrNotImplemented
}

func (StubService) UpdateSettings(ctx context.Context, userID string, request UpdateSettingsRequest) (Settings, error) {
	return Settings{}, scaffold.ErrNotImplemented
}
