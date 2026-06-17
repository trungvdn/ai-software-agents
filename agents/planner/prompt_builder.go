package planner

import "strings"

type PlannerPromptBuilder struct {
}

func (b *PlannerPromptBuilder) Build(
	bugDescription string,
	analysis string,
) string {

	promptBuilder := strings.Builder{}
	promptBuilder.WriteString("You are a senior software engineer analyzing a bug fix.\n\n")
	promptBuilder.WriteString("Bug:\n" + bugDescription + "\n\n")
	promptBuilder.WriteString("Analysis:\n" + analysis + "\n\n")
	promptBuilder.WriteString("Based on the analysis, provide your response ONLY as a valid JSON object (no markdown, no extra text) with exactly this structure:\n")
	promptBuilder.WriteString("{\n")
	promptBuilder.WriteString("  \"affected_files\": [\"file1.go\", \"file2.go\"],\n")
	promptBuilder.WriteString("  \"steps\": [\"step 1\", \"step 2\", \"step 3\"]\n")
	promptBuilder.WriteString("}\n\n")
	promptBuilder.WriteString("Where:\n")
	promptBuilder.WriteString("- affected_files: List of Go source files that need to be modified\n")
	promptBuilder.WriteString("- steps: List of concrete implementation steps to fix the bug\n")
	promptBuilder.WriteString("- Ensure all strings are properly escaped\n")
	promptBuilder.WriteString("- Return ONLY valid JSON, no other text\n")

	prompt := promptBuilder.String()
	return prompt
}
