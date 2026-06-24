package llm

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/trungvdn/ai-software-agents/internal/requirement/application/generate_story"
	"github.com/trungvdn/ai-software-agents/internal/requirement/domain/story"
	"github.com/trungvdn/ai-software-agents/internal/requirement/infrastructure/llm/dto"
	"github.com/trungvdn/ai-software-agents/shared/llm"
	"github.com/trungvdn/ai-software-agents/shared/utils"
)

type OllamaStoryGenerator struct {
	client llm.Client
}

func NewOllamaStoryGenerator(
	client llm.Client,
) *OllamaStoryGenerator {
	return &OllamaStoryGenerator{
		client: client,
	}
}

func (o *OllamaStoryGenerator) Generate(
	ctx context.Context,
	request generate_story.GenerateStoryRequest,
) (*generate_story.GenerateStoryResponse, error) {
	/*
				You are a Senior Business Analyst.

		Your task is to generate User Stories from an Epic.

		Rules:

		1. Follow INVEST principles:
		   - Independent
		   - Negotiable
		   - Valuable
		   - Estimable
		   - Small
		   - Testable

		2. Use the format:

		As a <user>

		I want <goal>

		So that <benefit>

		3. Generate only business stories.
		4. Do not generate technical implementation tasks.
		5. Generate between 3 and 10 stories.

		Epic:

		Name:
		{{EPIC_NAME}}

		Description:
		{{EPIC_DESCRIPTION}}

		Return ONLY valid JSON.

		Schema:

		{
		  "stories": [
		    {
		      "title": "string",
		      "as_a": "string",
		      "i_want": "string",
		      "so_that": "string"
		    }
		  ]
		}
	*/

	prompt := strings.Builder{}
	prompt.WriteString("You are a Senior Business Analyst.\n\n")
	prompt.WriteString("Your task is to generate User Stories from an Epic.\n\n")
	prompt.WriteString("Rules:\n")
	prompt.WriteString("1. Follow INVEST principles:\n")
	prompt.WriteString("   - Independent\n")
	prompt.WriteString("   - Negotiable\n")
	prompt.WriteString("   - Valuable\n")
	prompt.WriteString("   - Estimable\n")
	prompt.WriteString("   - Small\n")
	prompt.WriteString("   - Testable\n")
	prompt.WriteString("2. Use the format:\n\n")
	prompt.WriteString("As a <user>\n\n")
	prompt.WriteString("I want <goal>\n\n")
	prompt.WriteString("So that <benefit>\n\n")
	prompt.WriteString("3. Generate only business stories.\n")
	prompt.WriteString("4. Do not generate technical implementation tasks.\n")
	prompt.WriteString("5. Generate between 3 and 10 stories.\n\n")
	prompt.WriteString("Epic:\n\n")
	prompt.WriteString("Name:\n")
	prompt.WriteString(request.Epic.Name)
	prompt.WriteString("\n\nDescription:\n")
	prompt.WriteString(request.Epic.Description)
	prompt.WriteString("\n\nReturn ONLY valid JSON.\n\n")
	prompt.WriteString("{\n")
	prompt.WriteString(`  "stories": [` + "\n")
	prompt.WriteString(`    {` + "\n")
	prompt.WriteString(`      "title": "string",` + "\n")
	prompt.WriteString(`      "as_a": "string",` + "\n")
	prompt.WriteString(`      "i_want": "string",` + "\n")
	prompt.WriteString(`      "so_that": "string"` + "\n")
	prompt.WriteString("    }\n")
	prompt.WriteString("  ]\n")
	prompt.WriteString("}\n")

	llmResponse, err := o.client.Chat(ctx, prompt.String())
	if err != nil {
		return nil, err
	}
	jsonStr := utils.StripCodeFences(llmResponse)

	// Parse the LLM response as JSON
	var response dto.StoryResponse
	if err := json.Unmarshal([]byte(jsonStr), &response); err != nil {
		return nil, fmt.Errorf("failed to parse LLM response as JSON: %w, response: %s", err, llmResponse)
	}
	log.Printf("OllamaStoryGenerator response: %+v", response)

	stories := make([]story.Story, len(response.Stories))
	for i, storyItem := range response.Stories {
		stories[i] = story.Story{
			Title:  storyItem.Title,
			AsA:    storyItem.AsA,
			IWant:  storyItem.IWant,
			SoThat: storyItem.SoThat,
		}
	}

	return &generate_story.GenerateStoryResponse{
		Stories: stories,
	}, nil
}
