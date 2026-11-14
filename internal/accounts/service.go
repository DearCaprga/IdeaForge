package accounts

import (
	"context"

	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
)

type Service interface {
	Authenticate(ctx context.Context, accessToken string) (Principal, error)
	Register(ctx context.Context, request RegisterRequest) (AuthResponse, error)
	Login(ctx context.Context, request LoginRequest) (AuthResponse, error)
	Refresh(ctx context.Context, request RefreshRequest) (AuthResponse, error)
	Logout(ctx context.Context, sessionID string) error
	GetPublicProfile(ctx context.Context, userID string) (Profile, error)
	GetPrivateProfile(ctx context.Context, userID string) (PrivateProfile, error)
	UpdateProfile(ctx context.Context, userID string, request UpdateProfileRequest) (PrivateProfile, error)
	GetSettings(ctx context.Context, userID string) (Settings, error)
	UpdateSettings(ctx context.Context, userID string, request UpdateSettingsRequest) (Settings, error)
}

type StubService struct{}

var _ Service = StubService{}

func (StubService) Authenticate(ctx context.Context, accessToken string) (Principal, error) {
	return Principal{}, scaffold.ErrNotImplemented
}

func (StubService) Register(ctx context.Context, request RegisterRequest) (AuthResponse, error) {
	return AuthResponse{}, scaffold.ErrNotImplemented
}

func (StubService) Login(ctx context.Context, request LoginRequest) (AuthResponse, error) {
	return AuthResponse{}, scaffold.ErrNotImplemented
}

func (StubService) Refresh(ctx context.Context, request RefreshRequest) (AuthResponse, error) {
	return AuthResponse{}, scaffold.ErrNotImplemented
}

func (StubService) Logout(ctx context.Context, sessionID string) error {
	return scaffold.ErrNotImplemented
}

func (StubService) GetPublicProfile(ctx context.Context, userID string) (Profile, error) {
	return Profile{}, scaffold.ErrNotImplemented
}

func (StubService) GetPrivateProfile(ctx context.Context, userID string) (PrivateProfile, error) {
	return PrivateProfile{}, scaffold.ErrNotImplemented
}

func (StubService) UpdateProfile(ctx context.Context, userID string, request UpdateProfileRequest) (PrivateProfile, error) {
	return PrivateProfile{}, scaffold.ErrNotImplemented
}

func (StubService) GetSettings(ctx context.Context, userID string) (Settings, error) {
	return Settings{}, scaffold.ErrNotImplemented
}

func (StubService) UpdateSettings(ctx context.Context, userID string, request UpdateSettingsRequest) (Settings, error) {
	return Settings{}, scaffold.ErrNotImplemented
}
