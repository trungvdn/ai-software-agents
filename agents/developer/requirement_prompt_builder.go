package developer

import (
	"strings"

	"github.com/trungvdn/ai-software-agents/domain/developer"
)

type RequirementPromptBuilder struct{}

func NewRequirementPromptBuilder() *RequirementPromptBuilder {
	return &RequirementPromptBuilder{}
}

func (b *RequirementPromptBuilder) Build(
	task *developer.DevelopmentTask,
) string {
	/*
		You are a senior software architect.

		Feature Request:
		Add forgot password feature

		Analyze:

		1. Goal
		2. Candidate Symbols
		3. Candidate Packages
		4. Technical Tasks
		5. Acceptance Criteria

		Return JSON only.
	*/
	var prompt strings.Builder
	prompt.WriteString("You are a senior software engineer.\n\n")
	prompt.WriteString("Feature Request:\n" + task.Description + "\n\n")
	prompt.WriteString("Analyze:\n")
	prompt.WriteString("1. Goal\n")
	prompt.WriteString("2. Candidate Symbols\n")
	prompt.WriteString("3. Candidate Packages\n")
	prompt.WriteString("4. Technical Tasks\n")
	prompt.WriteString("5. Acceptance Criteria\n\n")
	prompt.WriteString("Return JSON only.\n")
	return prompt.String()
}
