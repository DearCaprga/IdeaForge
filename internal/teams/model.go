package teams

import "time"

type Status string

const (
	StatusRecruiting Status = "recruiting"
	StatusClosed     Status = "closed"
)

type ApplicationStatus string

const (
	ApplicationPending  ApplicationStatus = "pending"
	ApplicationAccepted ApplicationStatus = "accepted"
	ApplicationRejected ApplicationStatus = "rejected"
)

type Opening struct {
	ID        string   `json:"id"`
	Role      string   `json:"role"`
	Skills    []string `json:"skills"`
	Positions int      `json:"positions"`
}

type Member struct {
	UserID    string    `json:"user_id"`
	OpeningID string    `json:"opening_id,omitempty"`
	Role      string    `json:"role"`
	JoinedAt  time.Time `json:"joined_at"`
}

type Team struct {
	ID              string     `json:"id"`
	ProjectID       string     `json:"project_id"`
	OwnerID         string     `json:"owner_id"`
	Name            string     `json:"name"`
	Description     string     `json:"description"`
	Status          Status     `json:"status"`
	RecruitingUntil *time.Time `json:"recruiting_until,omitempty"`
	Openings        []Opening  `json:"openings"`
	Members         []Member   `json:"members"`
	CreatedAt       time.Time  `json:"created_at"`
}

type Application struct {
	ID          string            `json:"id"`
	TeamID      string            `json:"team_id"`
	ApplicantID string            `json:"applicant_id"`
	OpeningID   string            `json:"opening_id"`
	Message     string            `json:"message"`
	Status      ApplicationStatus `json:"status"`
	CreatedAt   time.Time         `json:"created_at"`
	ReviewedAt  *time.Time        `json:"reviewed_at,omitempty"`
}
