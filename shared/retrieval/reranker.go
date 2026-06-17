package retrieval

import (
	"context"
)

type ReRanker interface {
	ReRank(
		ctx context.Context,
		query string,
		docs []SearchResult,
	) ([]SearchResult, error)
}
