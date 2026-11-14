package solutions

import "github.com/DearCaprga/IdeaForge/internal/projects"

type CreateInput struct {
	Content
	TeamID string `json:"team_id,omitempty"`
}

type UpdateInput struct {
	Title        *string          `json:"title,omitempty"`
	Summary      *string          `json:"summary,omitempty"`
	Approach     *string          `json:"approach,omitempty"`
	Technologies *[]string        `json:"technologies,omitempty"`
	Steps        *[]string        `json:"steps,omitempty"`
	Result       *string          `json:"result,omitempty"`
	Links        *[]projects.Link `json:"links,omitempty"`
	Attachments  *[]projects.Link `json:"attachments,omitempty"`
	Limitations  *string          `json:"limitations,omitempty"`
}

type ReviewInput struct {
	Clarity         int    `json:"clarity"`
	Functionality   int    `json:"functionality"`
	Originality     int    `json:"originality"`
	Reproducibility int    `json:"reproducibility"`
	Comment         string `json:"comment"`
}
