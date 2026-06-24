package requirement

import (
	"context"

	"github.com/trungvdn/ai-software-agents/internal/requirement/domain/epic"
	"github.com/trungvdn/ai-software-agents/internal/requirement/domain/story"
)

type Aggregate struct {
	Requirement Requirement
	Epics       []epic.Epic
	Stories     []story.Story
}

type AggregateRepository interface {
	SaveAggregate(
		ctx context.Context,
		aggregate Aggregate,
	) error

	GetAggregate(
		ctx context.Context,
		id string,
	) (*Aggregate, error)
}
