package bugfix

type PromptBuilder struct {
}

func (b *PromptBuilder) Build(
	bugDescription string,
	context string,
) string {
	/*
		You are a senior software engineer.
		Bug Description:
		Fix nil pointer in UserService

		Relevant Reflections:
		...

		Analyze the root cause and suggest a fix.
	*/
	var prompt strings.Builder
	prompt.WriteString("You are a senior software engineer.\n\n")
	prompt.WriteString("Bug Description:\n" + bugDescription + "\n\n")
	prompt.WriteString("Relevant Reflections:\n" + context + "\n\n")
	prompt.WriteString("Analyze the root cause and suggest a fix.")
	return prompt.String()
}
