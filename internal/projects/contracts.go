package projects

import (
	"time"

	"github.com/DearCaprga/IdeaForge/internal/platform/scaffold"
)

type SortOrder string

const (
	SortNewest     SortOrder = "newest"
	SortPopular    SortOrder = "popular"
	SortDifficulty SortOrder = "difficulty"
)

type DurationRange string

const (
	DurationWeek    DurationRange = "week"
	DurationMonth   DurationRange = "month"
	DurationQuarter DurationRange = "quarter"
	DurationLonger  DurationRange = "longer"
)

type CatalogQuery struct {
	scaffold.PageQuery
	Query            string        `form:"q"`
	Category         Category      `form:"category"`
	Topic            string        `form:"topic"`
	Difficulty       Difficulty    `form:"difficulty"`
	Duration         DurationRange `form:"duration"`
	Status           Status        `form:"status"`
	Technologies     []string      `form:"technologies"`
	Skills           []string      `form:"skills"`
	TeamAvailable    *bool         `form:"team_available"`
	PublishedAfter   *time.Time    `form:"published_after"`
	PublishedBefore  *time.Time    `form:"published_before"`
	VerifiedSource   *bool         `form:"verified_source"`
	Sort             SortOrder     `form:"sort"`
	Subject          string        `form:"subject"`
	SubjectArea      string        `form:"subject_area"`
	EducationLevel   string        `form:"education_level"`
	ResearchFormat   string        `form:"research_format"`
	Equipment        []string      `form:"equipment"`
	Industry         string        `form:"industry"`
	ProblemSource    string        `form:"problem_source"`
	TaskType         string        `form:"task_type"`
	HasSolutions     *bool         `form:"has_solutions"`
	HasDemandSignals *bool         `form:"has_demand_signals"`
	Direction        string        `form:"direction"`
	Languages        []string      `form:"languages"`
	Frameworks       []string      `form:"frameworks"`
	HasAlternatives  *bool         `form:"has_alternatives"`
	ResultType       string        `form:"result_type"`
}

type FilterOption struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

type CategoryOptions struct {
	Category Category                  `json:"category"`
	Fields   map[string][]FilterOption `json:"fields"`
}

type CatalogOptions struct {
	Categories     []FilterOption    `json:"categories"`
	Difficulties   []FilterOption    `json:"difficulties"`
	Durations      []FilterOption    `json:"durations"`
	Statuses       []FilterOption    `json:"statuses"`
	SortOrders     []FilterOption    `json:"sort_orders"`
	Topics         []FilterOption    `json:"topics"`
	Technologies   []FilterOption    `json:"technologies"`
	Skills         []FilterOption    `json:"skills"`
	CategoryFields []CategoryOptions `json:"category_fields"`
}

// DraftInput используется для создания и частичного редактирования черновика.
type DraftInput struct {
	Title            *string             `json:"title,omitempty"`
	Summary          *string             `json:"summary,omitempty"`
	Category         *Category           `json:"category,omitempty"`
	Topic            *string             `json:"topic,omitempty"`
	Description      *string             `json:"description,omitempty"`
	Problem          *string             `json:"problem,omitempty"`
	Relevance        *string             `json:"relevance,omitempty"`
	Goal             *string             `json:"goal,omitempty"`
	TargetAudience   *string             `json:"target_audience,omitempty"`
	ExpectedResult   *string             `json:"expected_result,omitempty"`
	Difficulty       *Difficulty         `json:"difficulty,omitempty"`
	EstimatedDaysMin *int                `json:"estimated_days_min,omitempty"`
	EstimatedDaysMax *int                `json:"estimated_days_max,omitempty"`
	Tags             *[]string           `json:"tags,omitempty"`
	SuccessCriteria  *[]string           `json:"success_criteria,omitempty"`
	Skills           *[]string           `json:"skills,omitempty"`
	Technologies     *[]string           `json:"technologies,omitempty"`
	Resources        *[]Link             `json:"resources,omitempty"`
	Attachments      *[]Link             `json:"attachments,omitempty"`
	Stages           *[]StageInput       `json:"stages,omitempty"`
	Academic         *AcademicDetails    `json:"academic,omitempty"`
	ProductCase      *ProductCaseDetails `json:"product_case,omitempty"`
	Development      *DevelopmentDetails `json:"development,omitempty"`
	Source           *SourceInput        `json:"source,omitempty"`
	OrganizationID   *string             `json:"organization_id,omitempty"`
	CoverURL         *string             `json:"cover_url,omitempty"`
	TeamAvailable    *bool               `json:"team_available,omitempty"`
	AcceptsSolutions *bool               `json:"accepts_solutions,omitempty"`
}

type SourceInput struct {
	ID                 string     `json:"id,omitempty"`
	URL                string     `json:"url"`
	Type               string     `json:"type"`
	PublishedAt        *time.Time `json:"published_at,omitempty"`
	ProblemDescription string     `json:"problem_description,omitempty"`
}

type FieldWarning struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type PublicationPreview struct {
	Project            Project          `json:"project"`
	Warnings           []FieldWarning   `json:"warnings"`
	PossibleDuplicates []ProjectSummary `json:"possible_duplicates"`
}

type MyProjectsTab string

const (
	MyProjectsCreated   MyProjectsTab = "created"
	MyProjectsActive    MyProjectsTab = "in_progress"
	MyProjectsCompleted MyProjectsTab = "completed"
	MyProjectsFavorites MyProjectsTab = "favorites"
	MyProjectsDrafts    MyProjectsTab = "drafts"
)

type MyProjectsQuery struct {
	scaffold.PageQuery
	Tab MyProjectsTab `form:"tab"`
}

type UpdateStageInput struct {
	Completed bool `json:"completed"`
}

type HomeSection struct {
	Key      string           `json:"key"`
	Title    string           `json:"title"`
	Projects []ProjectSummary `json:"projects"`
}

type RecentSolution struct {
	ID        string    `json:"id"`
	ProjectID string    `json:"project_id"`
	Title     string    `json:"title"`
	Author    AuthorRef `json:"author"`
}

type Home struct {
	Sections        []HomeSection    `json:"sections"`
	LatestSolutions []RecentSolution `json:"latest_solutions"`
}

type SuggestionQuery struct {
	Query string `form:"q"`
	Limit int    `form:"limit"`
}

type SearchSuggestion struct {
	Text      string `json:"text"`
	ProjectID string `json:"project_id,omitempty"`
}

type UpdateWorkInput struct {
	Status WorkStatus `json:"status"`
}
