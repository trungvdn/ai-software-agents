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

type AggregateRepository interface {
	SaveAggregate(
		ctx context.Context,
		aggregate RequirementAggregate,
	) error

	GetAggregate(
		ctx context.Context,
		id string,
	) (*RequirementAggregate, error)
}
