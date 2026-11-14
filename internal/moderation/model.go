// Package moderation defines publication review, reports, and administration.
package moderation

import (
	"time"

	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
)

type SubjectKind string

const (
	SubjectProject  SubjectKind = "project"
	SubjectSolution SubjectKind = "solution"
	SubjectComment  SubjectKind = "comment"
	SubjectSource   SubjectKind = "source"
)

type Subject struct {
	Kind SubjectKind `json:"kind"`
	ID   string      `json:"id"`
}

type ReviewStatus string

const (
	ReviewPending          ReviewStatus = "pending"
	ReviewApproved         ReviewStatus = "approved"
	ReviewChangesRequested ReviewStatus = "changes_requested"
	ReviewRejected         ReviewStatus = "rejected"
)

type QueueItem struct {
	ID          string       `json:"id"`
	Subject     Subject      `json:"subject"`
	Title       string       `json:"title"`
	SubmitterID string       `json:"submitter_id"`
	Status      ReviewStatus `json:"status"`
	SubmittedAt time.Time    `json:"submitted_at"`
}

type SubmitForReviewInput struct {
	Subject Subject
	Title   string
}

type QueueQuery struct {
	scaffold.PageQuery
	Kind   SubjectKind  `form:"kind"`
	Status ReviewStatus `form:"status"`
}

type CheckKind string

const (
	CheckDuplicate CheckKind = "duplicate"
	CheckSource    CheckKind = "source"
)

type CheckStatus string

const (
	CheckPending CheckStatus = "pending"
	CheckPassed  CheckStatus = "passed"
	CheckFlagged CheckStatus = "flagged"
)

type Check struct {
	Kind              CheckKind   `json:"kind"`
	Status            CheckStatus `json:"status"`
	Details           string      `json:"details"`
	RelatedProjectIDs []string    `json:"related_project_ids"`
	SourceURL         string      `json:"source_url,omitempty"`
	CheckedAt         *time.Time  `json:"checked_at,omitempty"`
}

type ReviewDecision struct {
	Status ReviewStatus `json:"status"`
	Reason string       `json:"reason"`
}

type ReportReason string

const (
	ReportOutdated   ReportReason = "outdated"
	ReportUnreliable ReportReason = "unreliable"
	ReportDuplicate  ReportReason = "duplicate"
	ReportAbuse      ReportReason = "abuse"
	ReportOther      ReportReason = "other"
)

type ReportStatus string

const (
	ReportOpen      ReportStatus = "open"
	ReportResolved  ReportStatus = "resolved"
	ReportDismissed ReportStatus = "dismissed"
)

type CreateReportInput struct {
	Subject     Subject      `json:"subject"`
	Reason      ReportReason `json:"reason"`
	Description string       `json:"description"`
}

type Report struct {
	ID          string       `json:"id"`
	ReporterID  string       `json:"reporter_id"`
	Subject     Subject      `json:"subject"`
	Reason      ReportReason `json:"reason"`
	Description string       `json:"description"`
	Status      ReportStatus `json:"status"`
	CreatedAt   time.Time    `json:"created_at"`
}

type ReportQuery struct {
	scaffold.PageQuery
	Status ReportStatus `form:"status"`
	Kind   SubjectKind  `form:"kind"`
}

type ReportDecision struct {
	Status ReportStatus `json:"status"`
	Reason string       `json:"reason"`
}

type AuditAction string

const (
	AuditReviewDecision AuditAction = "review_decision"
	AuditReportDecision AuditAction = "report_decision"
	AuditSourceCheck    AuditAction = "source_check"
	AuditTaxonomyEdit   AuditAction = "taxonomy_edit"
)

type AuditEntry struct {
	ID          string      `json:"id"`
	ModeratorID string      `json:"moderator_id"`
	Action      AuditAction `json:"action"`
	TargetID    string      `json:"target_id"`
	Reason      string      `json:"reason"`
	CreatedAt   time.Time   `json:"created_at"`
}

type AuditQuery struct {
	scaffold.PageQuery
	ModeratorID string      `form:"moderator_id"`
	Action      AuditAction `form:"action"`
	From        *time.Time  `form:"from"`
	To          *time.Time  `form:"to"`
}

type CategoryInput struct {
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ParentID    string `json:"parent_id,omitempty"`
}

type UpdateCategoryInput struct {
	Slug        *string `json:"slug,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ParentID    *string `json:"parent_id,omitempty"`
}

type Category struct {
	ID string `json:"id"`
	CategoryInput
}

type TagInput struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type UpdateTagInput struct {
	Slug *string `json:"slug,omitempty"`
	Name *string `json:"name,omitempty"`
}

type Tag struct {
	ID string `json:"id"`
	TagInput
}

type UsageQuery struct {
	From time.Time `form:"from" json:"from"`
	To   time.Time `form:"to" json:"to"`
}

type UsageMetrics struct {
	UsageQuery
	PublishedProjects  int `json:"published_projects"`
	PublishedSolutions int `json:"published_solutions"`
	ActiveAuthors      int `json:"active_authors"`
}
