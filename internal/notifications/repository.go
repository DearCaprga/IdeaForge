package notifications

import (
	"context"
	"time"

	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
)

type Repository interface {
	Create(ctx context.Context, notification Notification) error
	List(ctx context.Context, userID string, query ListQuery) (scaffold.Page[Notification], error)
	CountUnread(ctx context.Context, userID string) (int, error)
	MarkRead(ctx context.Context, userID string, notificationID string, readAt time.Time) (Notification, error)
	GetSettings(ctx context.Context, userID string) (Settings, error)
	UpdateSettings(ctx context.Context, userID string, changes UpdateSettingsRequest) (Settings, error)
}
