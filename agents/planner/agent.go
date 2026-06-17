package planner

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/trungvdn/ai-software-agents/domain/changeplan"
	"github.com/trungvdn/ai-software-agents/shared/llm"
)

type LLMChangePlanner struct {
	llm            llm.Client
	contextBuilder prompt_context.Builder
}

// LLMResponse represents the structured response from the LLM
type LLMResponse struct {
	AffectedFiles []string `json:"affected_files"`
	Steps         []string `json:"steps"`
}

func NewLLMChangePlanner(llm llm.Client, contextBuilder prompt_context.Builder) *LLMChangePlanner {
	return &LLMChangePlanner{
		llm:            llm,
		contextBuilder: contextBuilder,
	}
}

func (p *LLMChangePlanner) Plan(
	ctx context.Context,
	bugDescription string,
	analysis string,
) (*changeplan.ChangePlan, error) {

	// Build a structured prompt that guides the LLM to output valid JSON
	promptBuilder := &PlannerPromptBuilder{}
	prompt := promptBuilder.Build(
		bugDescription,
		analysis,
	)

	response, err := p.llm.Chat(ctx, prompt)
	if err != nil {
		return nil, fmt.Errorf("failed to get response from LLM: %w", err)
	}

	// Parse the LLM response as JSON
	var llmResp LLMResponse
	if err := json.Unmarshal([]byte(response), &llmResp); err != nil {
		return nil, fmt.Errorf("failed to parse LLM response as JSON: %w, response: %s", err, response)
	}

	// Validate the response contains required fields
	if len(llmResp.AffectedFiles) == 0 {
		// warning log
		log.Printf("Warning: LLM response contains no affected files. Response: %s", response)
		return nil, nil
	}
	if len(llmResp.Steps) == 0 {
		log.Printf("Warning: LLM response contains no implementation steps. Response: %s", response)
		return nil, nil
	}

	plan := &changeplan.ChangePlan{
		AffectedFiles: llmResp.AffectedFiles,
		Steps:         llmResp.Steps,
	}
	return plan, nil
}
