package accounts

import "time"

type RegisterRequest struct {
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type AuthResponse struct {
	User            PrivateProfile `json:"user"`
	AccessToken     string         `json:"access_token"`
	RefreshToken    string         `json:"refresh_token"`
	AccessExpiresAt time.Time      `json:"access_expires_at"`
}

type UpdateProfileRequest struct {
	DisplayName *string          `json:"display_name,omitempty"`
	AvatarURL   *string          `json:"avatar_url,omitempty"`
	Bio         *string          `json:"bio,omitempty"`
	Interests   *[]string        `json:"interests,omitempty"`
	Skills      *[]string        `json:"skills,omitempty"`
	Portfolio   *[]PortfolioLink `json:"portfolio,omitempty"`
}

type UpdateSettingsRequest struct {
	Theme *Theme `json:"theme,omitempty"`
}
