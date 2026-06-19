package coder

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/trungvdn/ai-software-agents/domain/changeplan"
	"github.com/trungvdn/ai-software-agents/domain/codepatch"
	"github.com/trungvdn/ai-software-agents/shared/llm"
	"github.com/trungvdn/ai-software-agents/shared/utils"
)

type CoderAgent struct {
	llm llm.Client
}

func NewCoderAgent(
	llm llm.Client,
) *CoderAgent {
	return &CoderAgent{
		llm: llm,
	}
}

func (a *CoderAgent) GeneratePatches(
	ctx context.Context,
	bugDescription string,
	plan *changeplan.ChangePlan,
) ([]codepatch.CodePatch, error) {

	// Build context for code generation

	/*
		Change Plan:

		Affected Files:
		- user_service.go

		Steps:
		1. Add nil check
	*/

	coderContext := strings.Builder{}
	coderContext.WriteString("Change Plan:\n")
	coderContext.WriteString("\n")
	coderContext.WriteString("Affected Files:\n")
	for _, file := range plan.AffectedFiles {
		coderContext.WriteString(fmt.Sprintf("-%s\n", file))
	}
	coderContext.WriteString("\n")
	coderContext.WriteString("Steps:\n")
	for i, step := range plan.Steps {
		coderContext.WriteString(fmt.Sprintf("%d. %s\n", i+1, step))
	}

	// Build prompt for code generation
	promptBuilder := NewCoderPromptBuilder()
	prompt := promptBuilder.Build(
		bugDescription,
		coderContext.String(),
	)

	// Generate code patches using LLM
	response, err := a.llm.Chat(ctx, prompt)
	if err != nil {
		return nil, err
	}

	fmt.Println("LLM Response:", response)

	// Strip markdown code blocks if present (e.g., ```json ... ```)
	jsonStr := utils.StripCodeFences(response)

	// Parse the LLM response as JSON
	var patchResponse PatchResponse
	if err := json.Unmarshal([]byte(jsonStr), &patchResponse); err != nil {
		return nil, fmt.Errorf("failed to parse LLM response as JSON: %w, response: %s", err, response)
	}

	return patchResponse.Patches, nil
}
