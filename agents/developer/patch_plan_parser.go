package developer

import (
	"encoding/json"
	"fmt"

	"github.com/trungvdn/ai-software-agents/domain/patchplan"
	"github.com/trungvdn/ai-software-agents/shared/utils"
)

func ParsePatchPlan(
	response string,
) (*patchplan.PatchPlan, error) {
	// Strip markdown code blocks if present (e.g., ```json ... ```)
	jsonStr := utils.StripCodeFences(response)

	// Parse the LLM response as JSON
	var patchPlan patchplan.PatchPlan
	if err := json.Unmarshal([]byte(jsonStr), &patchPlan); err != nil {
		return nil, fmt.Errorf("failed to parse LLM response as JSON: %w, response: %s", err, response)
	}
	return &patchPlan, nil
}
