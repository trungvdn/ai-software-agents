package dto

type RequirementResponse struct {
	ProjectName string   `json:"project_name"`
	Vision      string   `json:"vision"`
	Goals       []string `json:"goals"`
}
