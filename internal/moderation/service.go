package moderation

import (
	"context"

	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
)

type Service struct {
	repository Repository
	reviewer   Reviewer
}

func NewService(repository Repository, reviewer Reviewer) *Service {
	return &Service{repository: repository, reviewer: reviewer}
}

func (s *Service) SubmitForReview(ctx context.Context, submitterID string, input SubmitForReviewInput) (QueueItem, error) {
	return QueueItem{}, scaffold.ErrNotImplemented
}

func (s *Service) ListQueue(ctx context.Context, query QueueQuery) (scaffold.Page[QueueItem], error) {
	return scaffold.Page[QueueItem]{}, scaffold.ErrNotImplemented
}

func (s *Service) GetQueueItem(ctx context.Context, itemID string) (QueueItem, error) {
	return QueueItem{}, scaffold.ErrNotImplemented
}

func (s *Service) CheckItem(ctx context.Context, moderatorID, itemID string) ([]Check, error) {
	return nil, scaffold.ErrNotImplemented
}

func (s *Service) DecideItem(ctx context.Context, moderatorID, itemID string, decision ReviewDecision) (QueueItem, error) {
	return QueueItem{}, scaffold.ErrNotImplemented
}

func (s *Service) SubmitReport(ctx context.Context, reporterID string, input CreateReportInput) (Report, error) {
	return Report{}, scaffold.ErrNotImplemented
}

func (s *Service) ListReports(ctx context.Context, query ReportQuery) (scaffold.Page[Report], error) {
	return scaffold.Page[Report]{}, scaffold.ErrNotImplemented
}

func (s *Service) DecideReport(ctx context.Context, moderatorID, reportID string, decision ReportDecision) (Report, error) {
	return Report{}, scaffold.ErrNotImplemented
}

func (s *Service) ListAuditEntries(ctx context.Context, query AuditQuery) (scaffold.Page[AuditEntry], error) {
	return scaffold.Page[AuditEntry]{}, scaffold.ErrNotImplemented
}

func (s *Service) ListCategories(ctx context.Context, query scaffold.PageQuery) (scaffold.Page[Category], error) {
	return scaffold.Page[Category]{}, scaffold.ErrNotImplemented
}

func (s *Service) CreateCategory(ctx context.Context, moderatorID string, input CategoryInput) (Category, error) {
	return Category{}, scaffold.ErrNotImplemented
}

func (s *Service) UpdateCategory(ctx context.Context, moderatorID, categoryID string, input UpdateCategoryInput) (Category, error) {
	return Category{}, scaffold.ErrNotImplemented
}

func (s *Service) DeleteCategory(ctx context.Context, moderatorID, categoryID string) error {
	return scaffold.ErrNotImplemented
}

func (s *Service) ListTags(ctx context.Context, query scaffold.PageQuery) (scaffold.Page[Tag], error) {
	return scaffold.Page[Tag]{}, scaffold.ErrNotImplemented
}

func (s *Service) CreateTag(ctx context.Context, moderatorID string, input TagInput) (Tag, error) {
	return Tag{}, scaffold.ErrNotImplemented
}

func (s *Service) UpdateTag(ctx context.Context, moderatorID, tagID string, input UpdateTagInput) (Tag, error) {
	return Tag{}, scaffold.ErrNotImplemented
}

func (s *Service) DeleteTag(ctx context.Context, moderatorID, tagID string) error {
	return scaffold.ErrNotImplemented
}

func (s *Service) GetUsageMetrics(ctx context.Context, query UsageQuery) (UsageMetrics, error) {
	return UsageMetrics{}, scaffold.ErrNotImplemented
}
