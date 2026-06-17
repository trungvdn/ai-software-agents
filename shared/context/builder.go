package context

import (
	"context"

	"github.com/trungvdn/ai-software-agents/shared/retrieval"
)

type Builder interface {
	Build(
		ctx context.Context,
		results []retrieval.SearchResult,
	) (Context, error)
}
