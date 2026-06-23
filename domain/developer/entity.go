package developer

type TaskType string

const (
	TaskTypeBugFix  TaskType = "bug_fix"
	TaskTypeFeature TaskType = "feature"
	TaskTypeTest    TaskType = "test"
)

type DevelopmentTask struct {
	Type        TaskType
	Title       string
	Description string
}

type RequirementAnalysis struct {
	Goal               string
	CandidateSymbols   []string
	TechnicalTasks     []string
	AcceptanceCriteria []string
}
