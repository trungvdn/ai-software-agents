package llm

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	generate_epic "github.com/trungvdn/ai-software-agents/internal/requirement/application/generateepic"
	"github.com/trungvdn/ai-software-agents/internal/requirement/domain/epic"
	"github.com/trungvdn/ai-software-agents/shared/llm"
	"github.com/trungvdn/ai-software-agents/shared/utils"
)

type OllamaEpicGenerator struct {
	client llm.Client
}

func NewOllamaEpicGenerator(
	client llm.Client,
) *OllamaEpicGenerator {
	return &OllamaEpicGenerator{
		client: client,
	}
}

func (o *OllamaEpicGenerator) Generate(
	ctx context.Context,
	request generate_epic.GenerateEpicRequest,
) (*generate_epic.GenerateEpicResponse, error) {
	/*
		You are a Senior Business Analyst.

		Your task is to create Epics from a Business Requirement.

		Rules:

		1. Generate only business capabilities.
		2. Do not generate technical tasks.
		3. Do not generate implementation details.
		4. Each Epic must represent a major business capability.
		5. Epic names must be concise.
		6. Avoid duplicate or overlapping Epics.
		7. Generate between 3 and 10 Epics.

		Business Requirement:

		Project Name:
		{{PROJECT_NAME}}

		Vision:
		{{VISION}}

		Goals:
		{{GOALS}}

		Return ONLY valid JSON.

		Schema:

		{
		  "epics": [
		    {
		      "name": "string",
		      "description": "string"
		    }
		  ]
		}
	*/
	prompt := strings.Builder{}
	prompt.WriteString("You are a Senior Business Analyst.\n\n")
	prompt.WriteString("Your task is to create Epics from a Business Requirement.\n\n")
	prompt.WriteString("Rules:\n")
	prompt.WriteString("1. Generate only business capabilities.\n")
	prompt.WriteString("2. Do not generate technical tasks.\n")
	prompt.WriteString("3. Do not generate implementation details.\n")
	prompt.WriteString("4. Each Epic must represent a major business capability.\n")
	prompt.WriteString("5. Epic names must be concise.\n")
	prompt.WriteString("6. Avoid duplicate or overlapping Epics.\n")
	prompt.WriteString("7. Generate between 3 and 10 Epics.\n\n")
	prompt.WriteString("Business Requirement:\n\n")
	prompt.WriteString("Project Name:\n")
	prompt.WriteString(request.Requirement.ProjectName)
	prompt.WriteString("\n\nVision:\n")
	prompt.WriteString(request.Requirement.Vision)
	prompt.WriteString("\n\nGoals:\n")
	for _, goal := range request.Requirement.Goals {
		prompt.WriteString("- " + goal.Description + "\n")
	}
	prompt.WriteString("\nReturn ONLY valid JSON.\n\n")
	prompt.WriteString("Schema:\n\n")
	prompt.WriteString("{\n")
	prompt.WriteString(`  "epics": [` + "\n")
	prompt.WriteString(`    {` + "\n")
	prompt.WriteString(`      "name": "string",` + "\n")
	prompt.WriteString(`      "description": "string"` + "\n")
	prompt.WriteString("    }\n")
	prompt.WriteString("  ]\n")
	prompt.WriteString("}\n")

	llmResponse, err := o.client.Chat(ctx, "")
	if err != nil {
		return nil, err
	}
	log.Printf(llmResponse)
	jsonStr := utils.StripCodeFences(llmResponse)

	// Parse the LLM response as JSON
	var epics []epic.Epic
	if err := json.Unmarshal([]byte(jsonStr), &epics); err != nil {
		return nil, fmt.Errorf("failed to parse LLM response as JSON: %w, response: %s", err, llmResponse)
	}

	return &generate_epic.GenerateEpicResponse{
		Epics: epics,
	}, nil
}
