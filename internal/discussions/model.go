package discussions

import (
	"time"

	"github.com/DearCaprga/IdeaForge/internal/projects"
)

type TargetType string

const (
	TargetProject  TargetType = "project"
	TargetSolution TargetType = "solution"
)

type Target struct {
	Type TargetType `json:"type"`
	ID   string     `json:"id"`
}

type ReactionKind string

const (
	ReactionLike   ReactionKind = "like"
	ReactionThanks ReactionKind = "thanks"
)

type ReactionCount struct {
	Kind  ReactionKind `json:"kind"`
	Count int          `json:"count"`
}

type Comment struct {
	ID         string             `json:"id"`
	ProjectID  string             `json:"project_id,omitempty"`
	SolutionID string             `json:"solution_id,omitempty"`
	ParentID   string             `json:"parent_id,omitempty"`
	Author     projects.AuthorRef `json:"author"`
	Body       string             `json:"body"`
	Reactions  []ReactionCount    `json:"reactions"`
	ReplyCount int                `json:"reply_count"`
	IsUseful   bool               `json:"is_useful"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
}

type Reaction struct {
	CommentID string       `json:"comment_id"`
	UserID    string       `json:"user_id"`
	Kind      ReactionKind `json:"kind"`
	CreatedAt time.Time    `json:"created_at"`
}
