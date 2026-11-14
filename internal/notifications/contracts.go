package notifications

import "github.com/DearCaprga/IdeaForge/internal/platform/scaffold"

type ListQuery struct {
	scaffold.PageQuery
	UnreadOnly bool `form:"unread_only"`
	Kind       Kind `form:"kind"`
}

type ListResponse struct {
	scaffold.Page[Notification]
	UnreadCount int `json:"unread_count"`
}

type UpdateSettingsRequest struct {
	EnabledKinds *[]Kind    `json:"enabled_kinds,omitempty"`
	Frequency    *Frequency `json:"frequency,omitempty"`
}

type NewNotification struct {
	UserID     string
	Kind       Kind
	Text       string
	References References
}
