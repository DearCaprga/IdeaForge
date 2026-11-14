package discussions

import "github.com/DearCaprga/IdeaForge/internal/platform/scaffold"

type SortOrder string

const (
	SortNewest SortOrder = "newest"
	SortOldest SortOrder = "oldest"
	SortUseful SortOrder = "useful"
)

type ListQuery struct {
	scaffold.PageQuery
	Sort SortOrder `form:"sort"`
}

type CommentInput struct {
	Body string `json:"body"`
}

type SetUsefulInput struct {
	IsUseful bool `json:"is_useful"`
}
