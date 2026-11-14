package accounts

import "time"

type Role string

const (
	RoleGuest        Role = "guest"
	RoleUser         Role = "user"
	RoleMentor       Role = "mentor"
	RoleOrganization Role = "organization"
	RoleModerator    Role = "moderator"
)

type Theme string

const (
	ThemeSystem Theme = "system"
	ThemeLight  Theme = "light"
	ThemeDark   Theme = "dark"
)

type Account struct {
	ID           string
	Email        string
	PasswordHash string `json:"-"`
	Roles        []Role
	CreatedAt    time.Time
}

type Session struct {
	ID               string
	UserID           string
	RefreshTokenHash string `json:"-"`
	ExpiresAt        time.Time
	RevokedAt        *time.Time
}

type Principal struct {
	UserID    string
	SessionID string
	Roles     []Role
}

type PortfolioLink struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}

type Organization struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
}

type Achievement struct {
	Code      string    `json:"code"`
	Title     string    `json:"title"`
	AwardedAt time.Time `json:"awarded_at"`
}

type ProfileStats struct {
	CompletedProjects int `json:"completed_projects"`
	Comments          int `json:"comments"`
	Teams             int `json:"teams"`
}

type Profile struct {
	ID           string          `json:"id"`
	DisplayName  string          `json:"display_name"`
	AvatarURL    string          `json:"avatar_url,omitempty"`
	Bio          string          `json:"bio"`
	Interests    []string        `json:"interests"`
	Skills       []string        `json:"skills"`
	Portfolio    []PortfolioLink `json:"portfolio"`
	Organization *Organization   `json:"organization,omitempty"`
	ProjectIDs   []string        `json:"project_ids"`
	SolutionIDs  []string        `json:"solution_ids"`
	Stats        ProfileStats    `json:"stats"`
	Achievements []Achievement   `json:"achievements"`
}

type PrivateProfile struct {
	Profile
	Email string `json:"email"`
	Roles []Role `json:"roles"`
}

type Settings struct {
	Theme Theme `json:"theme"`
}
