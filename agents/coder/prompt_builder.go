package coder

import (
	"strings"
)

type CoderPromptBuilder struct {
}

func NewCoderPromptBuilder() *CoderPromptBuilder {
	return &CoderPromptBuilder{}
}

func (b *CoderPromptBuilder) Build(
	bugDescription string,
	changePlan string,
) string {
	/*
		You are a senior Go engineer.
		Bug:
		Fix nil pointer in UserService

		Change Plan:
		1. Add nil check
		Affected Files:
		- user_service.go
		Steps:
		1. Add nil check

		Generate a code patch.

		Return JSON only.
	*/
	var prompt strings.Builder
	prompt.WriteString("You are a senior Go engineer.\n\n")
	prompt.WriteString("Bug:\n" + bugDescription + "\n\n")
	prompt.WriteString("Change Plan:\n" + changePlan + "\n\n")
	prompt.WriteString("Generate a code patch.\n\n")
	prompt.WriteString("Return JSON only.")
	return prompt.String()
}
