// Package collections defines curated project lists for the community and mentors.
package collections

import "time"

type Collection struct {
	ID          string    `json:"id"`
	AuthorID    string    `json:"author_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ProjectIDs  []string  `json:"project_ids"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CollectionInput struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	ProjectIDs  []string `json:"project_ids"`
}

type UpdateCollectionInput struct {
	Title       *string   `json:"title,omitempty"`
	Description *string   `json:"description,omitempty"`
	ProjectIDs  *[]string `json:"project_ids,omitempty"`
}

type CollectionQuery struct {
	AuthorID string `form:"author_id"`
}
