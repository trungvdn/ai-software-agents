package developer

import (
	"encoding/json"
	"fmt"

	"github.com/trungvdn/ai-software-agents/domain/patchcandidate"
	"github.com/trungvdn/ai-software-agents/shared/utils"
)

func ParsePatchCadidate(
	response string,
) (*patchcandidate.PatchCandidate, error) {
	// Strip markdown code blocks if present (e.g., ```json ... ```)
	jsonStr := utils.StripCodeFences(response)

	// Parse the LLM response as JSON
	var patchPlan patchcandidate.PatchCandidate
	if err := json.Unmarshal([]byte(jsonStr), &patchPlan); err != nil {
		return nil, fmt.Errorf("failed to parse LLM response as JSON: %w, response: %s", err, response)
	}
	return &patchPlan, nil
}
