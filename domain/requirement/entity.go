package requirement

type RequirementAnalysis struct {
	Goal               string
	CandidateSymbols   []string
	TechnicalTasks     []string
	AcceptanceCriteria []string
	Confidence         float64
}
