package requirement

type RequirementAnalysis struct {
	Goal               string
	CandidateSymbols   []string
	CandidatePackages  []string
	TechnicalTasks     []string
	AcceptanceCriteria []string
	Confidence         float64
}
