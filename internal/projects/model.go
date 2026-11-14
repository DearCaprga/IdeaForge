package projects

import "time"

type Category string

const (
	CategoryAcademic    Category = "academic"
	CategoryProductCase Category = "product_case"
	CategoryDevelopment Category = "development"
)

type Difficulty string

const (
	DifficultyBeginner     Difficulty = "beginner"
	DifficultyIntermediate Difficulty = "intermediate"
	DifficultyAdvanced     Difficulty = "advanced"
	DifficultyExpert       Difficulty = "expert"
)

type Status string

const (
	StatusOpen               Status = "open"
	StatusInProgress         Status = "in_progress"
	StatusSolutionsPublished Status = "solutions_published"
)

type PublicationStatus string

const (
	PublicationDraft         PublicationStatus = "draft"
	PublicationUnderReview   PublicationStatus = "under_review"
	PublicationPublished     PublicationStatus = "published"
	PublicationNeedsRevision PublicationStatus = "needs_revision"
)

type SourceStatus string

const (
	SourceUnverified SourceStatus = "unverified"
	SourceVerified   SourceStatus = "verified"
)

type AuthorRef struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
}

type OrganizationRef struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
}

type Link struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}

type Source struct {
	ID                 string       `json:"id,omitempty"`
	URL                string       `json:"url"`
	Type               string       `json:"type"`
	PublishedAt        *time.Time   `json:"published_at,omitempty"`
	Status             SourceStatus `json:"status"`
	ProblemDescription string       `json:"problem_description,omitempty"`
	VerifiedFacts      []string     `json:"verified_facts,omitempty"`
}

type Stage struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Position    int    `json:"position"`
}

type StageInput struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Position    int    `json:"position"`
}

type AcademicDetails struct {
	Subject        string   `json:"subject"`
	SubjectArea    string   `json:"subject_area"`
	EducationLevel string   `json:"education_level"`
	ResearchFormat string   `json:"research_format"`
	Equipment      []string `json:"equipment"`
}

type ProductCaseDetails struct {
	Industry          string `json:"industry"`
	ProblemSource     string `json:"problem_source"`
	TaskType          string `json:"task_type"`
	DemandSignals     []Link `json:"demand_signals"`
	NoveltyAssessment string `json:"novelty_assessment"`
}

type DevelopmentDetails struct {
	Direction    string   `json:"direction"`
	Languages    []string `json:"languages"`
	Frameworks   []string `json:"frameworks"`
	Alternatives []Link   `json:"alternatives"`
	ResultType   string   `json:"result_type"`
}

type ProjectSummary struct {
	ID               string           `json:"id"`
	Title            string           `json:"title"`
	Summary          string           `json:"summary"`
	Category         Category         `json:"category"`
	Topic            string           `json:"topic"`
	Tags             []string         `json:"tags"`
	Difficulty       Difficulty       `json:"difficulty"`
	EstimatedDaysMin int              `json:"estimated_days_min"`
	EstimatedDaysMax int              `json:"estimated_days_max"`
	Status           Status           `json:"status"`
	Author           AuthorRef        `json:"author"`
	Organization     *OrganizationRef `json:"organization,omitempty"`
	Source           *Source          `json:"source,omitempty"`
	CoverURL         string           `json:"cover_url,omitempty"`
	SolutionCount    int              `json:"solution_count"`
	TeamAvailable    bool             `json:"team_available"`
	PublishedAt      *time.Time       `json:"published_at,omitempty"`
}

type Project struct {
	ProjectSummary
	Description       string              `json:"description"`
	Problem           string              `json:"problem"`
	Relevance         string              `json:"relevance"`
	Goal              string              `json:"goal"`
	TargetAudience    string              `json:"target_audience"`
	ExpectedResult    string              `json:"expected_result"`
	SuccessCriteria   []string            `json:"success_criteria"`
	Skills            []string            `json:"skills"`
	Technologies      []string            `json:"technologies"`
	Resources         []Link              `json:"resources"`
	Attachments       []Link              `json:"attachments"`
	Stages            []Stage             `json:"stages"`
	Academic          *AcademicDetails    `json:"academic,omitempty"`
	ProductCase       *ProductCaseDetails `json:"product_case,omitempty"`
	Development       *DevelopmentDetails `json:"development,omitempty"`
	PublicationStatus PublicationStatus   `json:"publication_status"`
	AcceptsSolutions  bool                `json:"accepts_solutions"`
	Viewer            *ViewerState        `json:"viewer,omitempty"`
	CreatedAt         time.Time           `json:"created_at"`
	UpdatedAt         time.Time           `json:"updated_at"`
}

type ViewerState struct {
	IsFavorite bool        `json:"is_favorite"`
	WorkStatus *WorkStatus `json:"work_status,omitempty"`
}

type Favorite struct {
	UserID    string    `json:"user_id"`
	ProjectID string    `json:"project_id"`
	CreatedAt time.Time `json:"created_at"`
}

type WorkStatus string

const (
	WorkInProgress WorkStatus = "in_progress"
	WorkCompleted  WorkStatus = "completed"
)

type StageProgress struct {
	StageID     string     `json:"stage_id"`
	Completed   bool       `json:"completed"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

type ProjectWork struct {
	ID              string          `json:"id"`
	ProjectID       string          `json:"project_id"`
	UserID          string          `json:"user_id"`
	Status          WorkStatus      `json:"status"`
	Stages          []StageProgress `json:"stages"`
	ProgressPercent int             `json:"progress_percent"`
	NextStageID     string          `json:"next_stage_id,omitempty"`
	StartedAt       time.Time       `json:"started_at"`
	LastActivityAt  time.Time       `json:"last_activity_at"`
	CompletedAt     *time.Time      `json:"completed_at,omitempty"`
}

type MyProject struct {
	Project           ProjectSummary    `json:"project"`
	PublicationStatus PublicationStatus `json:"publication_status"`
	Work              *ProjectWork      `json:"work,omitempty"`
}
