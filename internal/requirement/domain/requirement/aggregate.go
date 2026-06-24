package requirement

import (
	"context"

	"github.com/trungvdn/ai-software-agents/internal/requirement/domain/epic"
	"github.com/trungvdn/ai-software-agents/internal/requirement/domain/story"
)

type RequirementAggregate struct {
	Requirement Requirement
	Epics       []epic.Epic
	Stories     []story.Story
}

type RequirementAggregateRepository interface {
	SaveAggregate(
		ctx context.Context,
		aggregate RequirementAggregate,
	) error

	GetAggregate(
		ctx context.Context,
		id string,
	) (*RequirementAggregate, error)
}

func NewRequirementAggregate(requirement Requirement, epics []epic.Epic, stories []story.Story) *RequirementAggregate {
	return &RequirementAggregate{
		Requirement: requirement,
		Epics:       epics,
		Stories:     stories,
	}
}
