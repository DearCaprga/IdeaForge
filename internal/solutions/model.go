package solutions

import (
	"time"

	"github.com/DearCaprga/IdeaForge/internal/projects"
)

type Content struct {
	Title        string          `json:"title"`
	Summary      string          `json:"summary"`
	Approach     string          `json:"approach"`
	Technologies []string        `json:"technologies"`
	Steps        []string        `json:"steps"`
	Result       string          `json:"result"`
	Links        []projects.Link `json:"links"`
	Attachments  []projects.Link `json:"attachments"`
	Limitations  string          `json:"limitations"`
}

type Solution struct {
	Content
	ID          string             `json:"id"`
	ProjectID   string             `json:"project_id"`
	Author      projects.AuthorRef `json:"author"`
	TeamID      string             `json:"team_id,omitempty"`
	UsefulCount int                `json:"useful_count"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

type Revision struct {
	ID          string    `json:"id"`
	SolutionID  string    `json:"solution_id"`
	EditorID    string    `json:"editor_id"`
	Content     Content   `json:"content"`
	PublishedAt time.Time `json:"published_at"`
}

type UsefulMark struct {
	SolutionID string    `json:"solution_id"`
	UserID     string    `json:"user_id"`
	CreatedAt  time.Time `json:"created_at"`
}

type Review struct {
	ID              string             `json:"id"`
	SolutionID      string             `json:"solution_id"`
	Author          projects.AuthorRef `json:"author"`
	Clarity         int                `json:"clarity"`
	Functionality   int                `json:"functionality"`
	Originality     int                `json:"originality"`
	Reproducibility int                `json:"reproducibility"`
	Comment         string             `json:"comment"`
	CreatedAt       time.Time          `json:"created_at"`
}
