package requirement

import (
	"context"

	"github.com/trungvdn/ai-software-agents/internal/requirement/domain/epic"
	"github.com/trungvdn/ai-software-agents/internal/requirement/domain/story"
)

type EpicAggregate struct {
	epic.Epic
	Stories []story.Story
}

type RequirementAggregate struct {
	Requirement Requirement
	Epics       []EpicAggregate
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

func NewRequirementAggregate(requirement Requirement, epics []EpicAggregate) *RequirementAggregate {
	return &RequirementAggregate{
		Requirement: requirement,
		Epics:       epics,
	}
}
