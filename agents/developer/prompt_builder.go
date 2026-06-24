package developer

import (
	"fmt"
	"strings"

	"github.com/trungvdn/ai-software-agents/domain/analysis"
	"github.com/trungvdn/ai-software-agents/internal/knowledge/application/retrieve_knowledge"
)

type DeveloperPromptBuilder struct {
}

func NewDefaultPromptBuilder() *DeveloperPromptBuilder {
	return &DeveloperPromptBuilder{}
}

func (b *DeveloperPromptBuilder) Build(
	bugDescription string,
	knowledge *retrieve_knowledge.KnowledgeContextResponse,
	codeContext *CodeContext,
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

		Relevant Code:
		1.File: internal/user/user_service.go
		<content>
		2.File: internal/user/repository.go
		<content>

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

	prompt.WriteString("Relevant Code:\n")
	if len(codeContext.Files) == 0 {
		prompt.WriteString("No relevant code files found.\n\n")
	} else {
		for i, codeFile := range codeContext.Files {
			prompt.WriteString(fmt.Sprintf("%d. File: %s\n", i+1, codeFile.Path))
			prompt.WriteString("```go\n")
			prompt.WriteString(codeFile.Content)
			prompt.WriteString("\n```\n\n")
		}
	}

	prompt.WriteString("Based on the analysis, provide your response ONLY as a valid JSON object (no markdown, no extra text) with exactly this structure:\n")
	prompt.WriteString("{\n")
	prompt.WriteString("  \"root_cause\": \"root cause,\"\n")
	prompt.WriteString("  \"fix_strategy\": \"strategy,\"\n")
	prompt.WriteString("  \"confidence\": 0.87\n")
	prompt.WriteString("}\n\n")
	prompt.WriteString("Where:\n")
	prompt.WriteString("- root_cause: root case bug\n")
	prompt.WriteString("- fix_strategy: fix bug strategy \n")
	prompt.WriteString("- confidence: confidence of fix strategy\n")
	prompt.WriteString("- Ensure all strings are properly escaped\n")
	prompt.WriteString("- Return ONLY valid JSON, no other text\n")
	prompt.WriteString("Analyze the root cause and suggest a fix.")
	return prompt.String()
}

func (b *DeveloperPromptBuilder) BuildPatchCandidate(
	analysis *analysis.Analysis,
	codeContext *CodeContext) string {
	/*
				You are a senior software engineer.
				Root cause:
				Fix nil pointer in UserService

				Bug Strategy:

				Relevant Files:
				1.File: internal/user/user_service.go
				<content>
				2.File: internal/user/repository.go
				<content>

				Based on the analysis, provide your response ONLY as a valid JSON array (no markdown, no extra text) with exactly this structure:
				{
					"file_path":"...",
		  			"reason":"...",
		  			"original_snippet":"...",
		  			"proposed_snippet":"..."
				}
				Create patch plan fix.
	*/
	var prompt strings.Builder
	prompt.WriteString("You are a senior software engineer.\n\n")
	prompt.WriteString("Root cause:\n" + analysis.RootCause + "\n\n")
	prompt.WriteString("Bug Strategy:\n" + analysis.SuggestedFix + "\n\n")
	prompt.WriteString("Relevant Code:\n")
	if len(codeContext.Files) == 0 {
		prompt.WriteString("No relevant code files found.\n\n")
	} else {
		for i, codeFile := range codeContext.Files {
			prompt.WriteString(fmt.Sprintf("%d. File: %s\n", i+1, codeFile.Path))
			prompt.WriteString("```go\n")
			prompt.WriteString(codeFile.Content)
			prompt.WriteString("\n```\n\n")
		}
	}
	prompt.WriteString("Based on the analysis, provide your response ONLY as a valid JSON array (no markdown, no extra text) with exactly this structure:\n")
	prompt.WriteString("{\n")
	prompt.WriteString("  \"file_path\": \"file path to patch\",\n")
	prompt.WriteString("  \"reason\": \"reason for the patch\",\n")
	prompt.WriteString("  \"original_snippet\": \"original code snippet\",\n")
	prompt.WriteString("  \"proposed_snippet\": \"proposed code snippet\"\n")
	prompt.WriteString("}\n\n")
	prompt.WriteString("Where:\n")
	prompt.WriteString("- file_path: file path to patch\n")
	prompt.WriteString("- reason: reason for the patch\n")
	prompt.WriteString("- changes: list of code changes, each change should be a string describing a single code modification (e.g., add a line, modify a line, delete a line)\n")
	prompt.WriteString("- Ensure all strings are properly escaped\n")
	prompt.WriteString("- Return ONLY valid JSON, no other text\n")
	prompt.WriteString("Create patch plan fix.")

	return prompt.String()
}
