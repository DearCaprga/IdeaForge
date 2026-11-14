// Package sources defines provenance and the review workflow for external projects.
package sources

import (
	"time"

	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
)

type Kind string

const (
	KindManual    Kind = "manual"
	KindReview    Kind = "review"
	KindForum     Kind = "forum"
	KindHackathon Kind = "hackathon"
	KindDataset   Kind = "dataset"
)

type Verification string

const (
	VerificationPending  Verification = "pending"
	VerificationVerified Verification = "verified"
	VerificationRejected Verification = "rejected"
)

type SourceInput struct {
	Name            string `json:"name"`
	Kind            Kind   `json:"kind"`
	URL             string `json:"url"`
	PermissionURL   string `json:"permission_url,omitempty"`
	PermissionNotes string `json:"permission_notes,omitempty"`
}

type UpdateSourceInput struct {
	Name            *string `json:"name,omitempty"`
	Kind            *Kind   `json:"kind,omitempty"`
	URL             *string `json:"url,omitempty"`
	PermissionURL   *string `json:"permission_url,omitempty"`
	PermissionNotes *string `json:"permission_notes,omitempty"`
}

type Source struct {
	ID string `json:"id"`
	SourceInput
	Verification  Verification `json:"verification"`
	ImportAllowed bool         `json:"import_allowed"`
	CreatedAt     time.Time    `json:"created_at"`
	CheckedAt     *time.Time   `json:"checked_at,omitempty"`
}

type SourceQuery struct {
	scaffold.PageQuery
	Kind         Kind         `form:"kind"`
	Verification Verification `form:"verification"`
}

type VerificationInput struct {
	Verification  Verification `json:"verification"`
	ImportAllowed bool         `json:"import_allowed"`
	Reason        string       `json:"reason"`
}

type Provenance struct {
	SourceID      string     `json:"source_id"`
	URL           string     `json:"url"`
	PublishedAt   *time.Time `json:"published_at,omitempty"`
	ProblemClaim  string     `json:"problem_claim"`
	VerifiedFacts []string   `json:"verified_facts"`
}

// SourceRecord belongs to the import pipeline and may contain unprocessed text.
type SourceRecord struct {
	ExternalID  string
	URL         string
	PublishedAt *time.Time
	Text        string
}

type ImportStatus string

const (
	ImportQueued        ImportStatus = "queued"
	ImportProcessing    ImportStatus = "processing"
	ImportPendingReview ImportStatus = "pending_review"
	ImportApproved      ImportStatus = "approved"
	ImportRejected      ImportStatus = "rejected"
	ImportFailed        ImportStatus = "failed"
	ImportPublished     ImportStatus = "published"
)

type StartImportInput struct {
	PublicationURL string     `json:"publication_url,omitempty"`
	PublishedSince *time.Time `json:"published_since,omitempty"`
}

type CreateManualDraftInput struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	CategoryID  string     `json:"category_id"`
	Difficulty  string     `json:"difficulty"`
	Provenance  Provenance `json:"provenance"`
}

type Candidate struct {
	ID                  string     `json:"id"`
	Title               string     `json:"title"`
	Description         string     `json:"description"`
	CategoryID          string     `json:"category_id"`
	SuggestedDifficulty string     `json:"suggested_difficulty"`
	Provenance          Provenance `json:"provenance"`
	PrivacyReviewed     bool       `json:"privacy_reviewed"`
	DuplicateProjectIDs []string   `json:"duplicate_project_ids"`
	ProjectID           string     `json:"project_id,omitempty"`
}

type Import struct {
	ID          string        `json:"id"`
	SourceID    string        `json:"source_id"`
	Status      ImportStatus  `json:"status"`
	RequestedBy string        `json:"requested_by"`
	CreatedAt   time.Time     `json:"created_at"`
	Candidates  []Candidate   `json:"candidates"`
	Review      *ImportReview `json:"review,omitempty"`
}

type ImportQuery struct {
	scaffold.PageQuery
	SourceID string
	Status   ImportStatus `form:"status"`
}

type ImportDecision string

const (
	DecisionApprove ImportDecision = "approve"
	DecisionReject  ImportDecision = "reject"
)

type ReviewImportInput struct {
	Decision    ImportDecision `json:"decision"`
	Reason      string         `json:"reason"`
	AcceptedIDs []string       `json:"accepted_ids"`
}

type ImportReview struct {
	ReviewImportInput
	ReviewerID string    `json:"reviewer_id"`
	ReviewedAt time.Time `json:"reviewed_at"`
}
