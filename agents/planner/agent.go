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
	llm llm.Client
}

// LLMResponse represents the structured response from the LLM
type LLMResponse struct {
	AffectedFiles []string `json:"affected_files"`
	Steps         []string `json:"steps"`
}

func NewLLMChangePlanner(llm llm.Client) *LLMChangePlanner {
	return &LLMChangePlanner{
		llm: llm,
	}
}

func (p *LLMChangePlanner) Plan(
	ctx context.Context,
	bugDescription string,
	analysis string,
) (*changeplan.ChangePlan, error) {
	// Build a structured prompt that guides the LLM to output valid JSON
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
		return nil, fmt.Errorf("LLM response contains no affected files")
	}
	if len(llmResp.Steps) == 0 {
		return nil, fmt.Errorf("LLM response contains no implementation steps")
	}

	plan := &changeplan.ChangePlan{
		AffectedFiles: llmResp.AffectedFiles,
		Steps:         llmResp.Steps,
	}
	return plan, nil
}
