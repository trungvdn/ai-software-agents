package developer

import (
	"encoding/json"
	"fmt"

	"github.com/trungvdn/ai-software-agents/domain/analysis"
	"github.com/trungvdn/ai-software-agents/shared/utils"
)

func ParseAnalysis(
	response string,
) (*analysis.Analysis, error) {
	// Strip markdown code blocks if present (e.g., ```json ... ```)
	jsonStr := utils.StripCodeFences(response)

	// Parse the LLM response as JSON
	var analysis analysis.Analysis
	if err := json.Unmarshal([]byte(jsonStr), &analysis); err != nil {
		return nil, fmt.Errorf("failed to parse LLM response as JSON: %w, response: %s", err, response)
	}
	return &analysis, nil
}
