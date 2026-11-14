package teams

import (
	"time"

	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
)

type ListQuery struct {
	scaffold.PageQuery
	ProjectID string   `form:"project_id"`
	OpenOnly  bool     `form:"open_only"`
	Search    string   `form:"q"`
	Skills    []string `form:"skills"`
}

type ApplicationsQuery struct {
	scaffold.PageQuery
	Status ApplicationStatus `form:"status"`
}

type OpeningInput struct {
	ID        string   `json:"id,omitempty"`
	Role      string   `json:"role"`
	Skills    []string `json:"skills"`
	Positions int      `json:"positions"`
}

type CreateRequest struct {
	Name            string         `json:"name"`
	Description     string         `json:"description"`
	RecruitingUntil *time.Time     `json:"recruiting_until,omitempty"`
	Openings        []OpeningInput `json:"openings"`
}

type UpdateRequest struct {
	Name            *string         `json:"name,omitempty"`
	Description     *string         `json:"description,omitempty"`
	Status          *Status         `json:"status,omitempty"`
	RecruitingUntil *time.Time      `json:"recruiting_until,omitempty"`
	Openings        *[]OpeningInput `json:"openings,omitempty"`
}

type ApplyRequest struct {
	OpeningID string `json:"opening_id"`
	Message   string `json:"message"`
}

type ReviewDecision string

const (
	ReviewAccept ReviewDecision = "accept"
	ReviewReject ReviewDecision = "reject"
)

type ReviewRequest struct {
	Decision ReviewDecision `json:"decision"`
}
