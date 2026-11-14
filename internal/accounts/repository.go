package accounts

import (
	"context"
	"time"
)

type Repository interface {
	CreateAccount(ctx context.Context, account Account, profile Profile, settings Settings) error
	FindAccountByID(ctx context.Context, userID string) (Account, error)
	FindAccountByEmail(ctx context.Context, email string) (Account, error)
	FindProfile(ctx context.Context, userID string) (Profile, error)
	UpdateProfile(ctx context.Context, userID string, changes UpdateProfileRequest) (Profile, error)
	FindSettings(ctx context.Context, userID string) (Settings, error)
	UpdateSettings(ctx context.Context, userID string, changes UpdateSettingsRequest) (Settings, error)
}

type SessionRepository interface {
	Create(ctx context.Context, session Session) error
	FindByID(ctx context.Context, sessionID string) (Session, error)
	FindByRefreshTokenHash(ctx context.Context, refreshTokenHash string) (Session, error)
	Rotate(ctx context.Context, previous Session, replacement Session) error
	Revoke(ctx context.Context, sessionID string, revokedAt time.Time) error
}
