package developer

import (
	"encoding/json"
	"fmt"

	"github.com/trungvdn/ai-software-agents/domain/requirement"
	"github.com/trungvdn/ai-software-agents/shared/utils"
)

func ParseRequirementAnalysis(
	response string,
) (*requirement.RequirementAnalysis, error) {
	// Strip markdown code blocks if present (e.g., ```json ... ```)
	jsonStr := utils.StripCodeFences(response)

	// Parse the LLM response as JSON
	var requirementAnalysis requirement.RequirementAnalysis
	if err := json.Unmarshal([]byte(jsonStr), &requirementAnalysis); err != nil {
		return nil, fmt.Errorf("failed to parse LLM response as JSON: %w, response: %s", err, response)
	}
	return &requirementAnalysis, nil
}
