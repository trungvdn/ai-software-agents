package publish_requirement

import "github.com/trungvdn/ai-software-agents/internal/requirement/domain/requirement"

type PublishRequirementRequest struct {
	RequirementAggregate *requirement.RequirementAggregate
}
