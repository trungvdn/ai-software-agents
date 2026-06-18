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
