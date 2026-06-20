package developer

import (
	"fmt"
	"strings"
)

type DeveloperPromptBuilder struct {
}

func (b *DeveloperPromptBuilder) Build(
	bugDescription string,
	knowledge *KnowledgeContext,
) string {
	/*
		You are a senior software engineer.
		Bug Description:
		Fix nil pointer in UserService

		Relevant Knowledge:
		[Reflection]
		1. ...
		[Historical Bug]
		1. ...

		Analyze the root cause and suggest a fix.
	*/
	var prompt strings.Builder
	prompt.WriteString("You are a senior software engineer.\n\n")
	prompt.WriteString("Bug Description:\n" + bugDescription + "\n\n")
	prompt.WriteString("Relevant Knowledge:\n")
	prompt.WriteString("[Reflection]\n")
	if len(knowledge.Reflections) == 0 {
		prompt.WriteString("No relevant reflections found.\n\n")
	} else {
		for i, reflection := range knowledge.Reflections {
			prompt.WriteString(fmt.Sprintf("%d. %s\n\n", i+1, reflection.Content))
		}
	}
	prompt.WriteString("[Historical Bug]\n")
	if len(knowledge.HistoricalBugs) == 0 {
		prompt.WriteString("No relevant historical bugs found.\n\n")
	} else {
		for i, historicalBug := range knowledge.HistoricalBugs {
			prompt.WriteString(fmt.Sprintf("%d. %s\n\n", i+1, historicalBug.Content))
		}
	}
	prompt.WriteString("Analyze the root cause and suggest a fix.")
	return prompt.String()
}
