package llm

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	generate_requirement "github.com/trungvdn/ai-software-agents/internal/requirement/application/generaterequirement"
	"github.com/trungvdn/ai-software-agents/internal/requirement/domain/requirement"
	"github.com/trungvdn/ai-software-agents/shared/llm"
	"github.com/trungvdn/ai-software-agents/shared/utils"
)

type OllamaRequirementGenerator struct {
	client llm.Client
}

func NewOllamaRequirementGenerator(
	client llm.Client,
) *OllamaRequirementGenerator {
	return &OllamaRequirementGenerator{
		client: client,
	}
}

func (o *OllamaRequirementGenerator) Generate(
	ctx context.Context,
	request generate_requirement.GenerateRequirementRequest,
) (*generate_requirement.GenerateRequirementResponse, error) {
	/*
		You are a senior business analyst.
		Generate:
		- Project Name
		- Vision
		- Goals
		Idea:

		{{IDEA}}

		Return ONLY valid JSON:
		{
			"project_name": "InvestPilot",
			"vision": "...",
			"goals": [
			"..."
			]
		}

	*/
	var prompt strings.Builder
	prompt.WriteString("You are a senior business analyst.\n\n")
	prompt.WriteString("Generate:\n")
	prompt.WriteString("- Project Name\n")
	prompt.WriteString("- Vision\n")
	prompt.WriteString("- Goals\n")
	prompt.WriteString("Idea:\n\n")
	prompt.WriteString(request.Idea)
	prompt.WriteString("\n\nReturn ONLY valid JSON:\n")
	prompt.WriteString("{\n")
	prompt.WriteString(`  "project_name": "InvestPilot",` + "\n")
	prompt.WriteString(`  "vision": "...",` + "\n")
	prompt.WriteString(`  "goals": [` + "\n")
	prompt.WriteString(`    "..."` + "\n")
	prompt.WriteString("  ]\n")
	prompt.WriteString("}\n")

	llmResponse, err := o.client.Chat(ctx, prompt.String())
	if err != nil {
		return nil, err
	}
	log.Printf(llmResponse)
	jsonStr := utils.StripCodeFences(llmResponse)

	// Parse the LLM response as JSON
	var requirement requirement.Requirement
	if err := json.Unmarshal([]byte(jsonStr), &requirement); err != nil {
		return nil, fmt.Errorf("failed to parse LLM response as JSON: %w, response: %s", err, llmResponse)
	}
	return &generate_requirement.GenerateRequirementResponse{
		Requirement: requirement,
	}, nil
}
