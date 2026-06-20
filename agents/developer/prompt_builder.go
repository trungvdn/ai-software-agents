package developer

import "strings"

type DeveloperPromptBuilder struct {
}

func (b *DeveloperPromptBuilder) Build(
	bugDescription string,
	analysis string,
) string {

	promptBuilder := strings.Builder{}
	prompt := promptBuilder.String()
	return prompt
}
