package retrieval

import (
	"context"
	"math"
	"sort"
)

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

	/* heuristic
		finalScore =
	    semantic*0.7 +
	    importance*0.2 +
	    usage*0.1
	*/
	for i, res := range results {
		semanticScore := res.Score
		importanceScore := res.Metadata.ImportanceScore
		usageScore :=
			math.Min(
				float64(res.Metadata.UsageCount)/100.0,
				1.0,
			)
		finalScore := semanticScore*0.7 + importanceScore*0.2 + usageScore*0.1
		results[i].Score = finalScore
	}

	// Sort the results by the final score in descending order
	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})

	return results, nil
}
