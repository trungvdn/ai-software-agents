package bugfix

type LLMChangePlanner struct {
	llm llm.Client
}

func (p *LLMChangePlanner) Plan(
	ctx context.Context,
	bugDescription string,
	analysis string,
) (*ChangePlan, error) {
	/* You are a senior software engineer.

	Bug:
	Fix nil pointer in UserService

	Analysis:
	...

	List:

	1. Affected files

	2. Required changes*/
	promptBuilder := strings.Builder{}
	promptBuilder.WriteString("You are a senior software engineer.\n\n")
	promptBuilder.WriteString("Bug:\n" + bugDescription + "\n\n")
	promptBuilder.WriteString("Analysis:\n" + analysis + "\n\n")
	promptBuilder.WriteString("List:\n\n1. Affected files\n\n2. Required changes")
	prompt := promptBuilder.String()
	response, err := p.llm.Chat(ctx, prompt)
	if err != nil {
		return nil, fmt.Errorf("failed to get response from LLM: %w", err)
	}
	// Parse the LLM response to create a ChangePlan
	// This is a simplified example; you would need to implement proper parsing logic
	// Example response format:
	/*
	  "affected_files": [
	    "user_service.go",
	    "user_repository.go"
	  ],
	  "steps": [
	    "Add nil check",
	    "Return ErrUserNotFound"
	  ]
	}*/
	// Convert response string to a structured format (e.g., JSON) and extract affected files and steps from the
	responseData := struct {
		AffectedFiles []string `json:"affected_files"`
		Steps         []string `json:"steps"`
	}{}
	err = json.Unmarshal([]byte(response), &responseData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse LLM response: %w", err)
	}

	plan := &ChangePlan{
		AffectedFiles: responseData.AffectedFiles,
		Steps:         responseData.Steps,
	}
	return plan, nil
}
