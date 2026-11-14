// Package intelligence defines contracts for optional project recommendations and author assistance.
package intelligence

import "github.com/DearCaprga/IdeaForge/internal/platform/scaffold"

type RecommendationQuery struct {
	scaffold.PageQuery
	CategoryID string `form:"category_id"`
}

type ReaderContext struct {
	Interests       []string
	Skills          []string
	SavedProjectIDs []string
}

type ProjectBrief struct {
	ProjectID      string   `json:"project_id,omitempty"`
	Idea           string   `json:"idea"`
	CategoryID     string   `json:"category_id,omitempty"`
	Problem        string   `json:"problem,omitempty"`
	ExpectedResult string   `json:"expected_result,omitempty"`
	Technologies   []string `json:"technologies,omitempty"`
	DurationDays   int      `json:"duration_days,omitempty"`
	Requirements   []string `json:"requirements,omitempty"`
}

type SuggestionMeta struct {
	Label             string `json:"label"`
	SuggestedBySystem bool   `json:"suggested_by_system"`
	Editable          bool   `json:"editable"`
}

type Recommendation struct {
	ProjectID string `json:"project_id"`
	Title     string `json:"title"`
	Reason    string `json:"reason"`
	SuggestionMeta
}

type ProposedStage struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type AuthorAssistance struct {
	SuggestionMeta
	Goal           string          `json:"goal"`
	ExpectedResult string          `json:"expected_result"`
	Skills         []string        `json:"skills"`
	Stages         []ProposedStage `json:"stages"`
}

type SimilarProject struct {
	ProjectID         string `json:"project_id"`
	Title             string `json:"title"`
	Reason            string `json:"reason"`
	PossibleDuplicate bool   `json:"possible_duplicate"`
}

type SimilarProjects struct {
	SuggestionMeta
	Items []SimilarProject `json:"items"`
}

type DifficultySuggestion struct {
	SuggestionMeta
	Difficulty string `json:"difficulty"`
	Rationale  string `json:"rationale"`
}

type RelevanceSignal struct {
	Description string `json:"description"`
	EvidenceURL string `json:"evidence_url,omitempty"`
}

type RelevanceAssessment struct {
	SuggestionMeta
	Summary string            `json:"summary"`
	Signals []RelevanceSignal `json:"signals"`
}
