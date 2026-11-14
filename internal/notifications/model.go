package notifications

import "time"

type Kind string

const (
	KindCommentReply        Kind = "comment_reply"
	KindNewSolution         Kind = "new_solution"
	KindTeamApplication     Kind = "team_application"
	KindTeamDecision        Kind = "team_decision"
	KindProjectUpdate       Kind = "project_update"
	KindMatchingProject     Kind = "matching_project"
	KindPublicationDecision Kind = "publication_decision"
	KindProjectReminder     Kind = "project_reminder"
)

type References struct {
	ProjectID     string `json:"project_id,omitempty"`
	SolutionID    string `json:"solution_id,omitempty"`
	CommentID     string `json:"comment_id,omitempty"`
	TeamID        string `json:"team_id,omitempty"`
	ApplicationID string `json:"application_id,omitempty"`
}

type Notification struct {
	ID         string     `json:"id"`
	UserID     string     `json:"-"`
	Kind       Kind       `json:"kind"`
	Text       string     `json:"text"`
	References References `json:"references"`
	CreatedAt  time.Time  `json:"created_at"`
	ReadAt     *time.Time `json:"read_at,omitempty"`
}

type Frequency string

const (
	FrequencyImmediate Frequency = "immediate"
	FrequencyDaily     Frequency = "daily"
	FrequencyWeekly    Frequency = "weekly"
)

type Settings struct {
	EnabledKinds []Kind    `json:"enabled_kinds"`
	Frequency    Frequency `json:"frequency"`
}
