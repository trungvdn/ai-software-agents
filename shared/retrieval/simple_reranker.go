package retrieval

import "context"

type SimpleReRanker struct {
}

func NewSimpleReRanker() *SimpleReRanker {
	return &SimpleReRanker{}
}

func (r *SimpleReRanker) ReRank(
	ctx context.Context,
	query string,
	results []SearchResult,
) ([]SearchResult, error) {

	// heuristic scoring

	return results, nil
}
