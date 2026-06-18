package historicalbug

import "github.com/trungvdn/ai-software-agents/shared/embedding"

type HistoricalBugRetriever struct {
	repo HistoricalBugRepository

	embedder embedding.Embedder
}

func NewHistoricalBugRetriever(
	repo HistoricalBugRepository,
	embedder embedding.Embedder,
) *HistoricalBugRetriever {
	return &HistoricalBugRetriever{
		repo:     repo,
		embedder: embedder,
	}
}

func (r *HistoricalBugRetriever) Retrieve(
	ctx context.Context,
	query string,
	topK int,
) ([]retrieval.SearchResult, error) {
	embedding, err := r.embedder.Embed(ctx, query)
	if err != nil {
		return nil, err
	}

	bugs, err := r.repo.SearchSimilar(ctx, embedding, topK)
	if err != nil {
		return nil, err
	}

	var results []retrieval.SearchResult
	for _, bug := range bugs {
		results = append(results, retrieval.SearchResult{
			ID:      bug.ID.String(),
			Content: bug.Title,
			Source:  "historical_bug",
		})
	}
	return results, nil
}
