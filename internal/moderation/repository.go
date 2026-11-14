package moderation

import (
	"context"

	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
)

// Repository defines the storage boundary shared by review and catalog administration.
type Repository interface {
	InsertQueueItem(ctx context.Context, submitterID string, input SubmitForReviewInput) (QueueItem, error)
	ListQueue(ctx context.Context, query QueueQuery) (scaffold.Page[QueueItem], error)
	FindQueueItem(ctx context.Context, itemID string) (QueueItem, error)
	SaveChecks(ctx context.Context, itemID string, checks []Check) error
	SaveReviewDecision(ctx context.Context, moderatorID, itemID string, decision ReviewDecision) (QueueItem, error)
	InsertReport(ctx context.Context, reporterID string, input CreateReportInput) (Report, error)
	ListReports(ctx context.Context, query ReportQuery) (scaffold.Page[Report], error)
	SaveReportDecision(ctx context.Context, moderatorID, reportID string, decision ReportDecision) (Report, error)
	AppendAuditEntry(ctx context.Context, entry AuditEntry) error
	ListAuditEntries(ctx context.Context, query AuditQuery) (scaffold.Page[AuditEntry], error)
	ListCategories(ctx context.Context, query scaffold.PageQuery) (scaffold.Page[Category], error)
	InsertCategory(ctx context.Context, input CategoryInput) (Category, error)
	UpdateCategory(ctx context.Context, categoryID string, input UpdateCategoryInput) (Category, error)
	DeleteCategory(ctx context.Context, categoryID string) error
	ListTags(ctx context.Context, query scaffold.PageQuery) (scaffold.Page[Tag], error)
	InsertTag(ctx context.Context, input TagInput) (Tag, error)
	UpdateTag(ctx context.Context, tagID string, input UpdateTagInput) (Tag, error)
	DeleteTag(ctx context.Context, tagID string) error
	GetUsageMetrics(ctx context.Context, query UsageQuery) (UsageMetrics, error)
}

type Reviewer interface {
	CheckDuplicates(ctx context.Context, subject Subject) (Check, error)
	CheckSource(ctx context.Context, subject Subject) (Check, error)
}
