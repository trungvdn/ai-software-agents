package requirement

type Goal struct {
}

type Status struct {
}

type Requirement struct {
	ID          string
	ProjectName string
	Vision      string
	Goals       []Goal
	Status      Status
	Version     int
}
