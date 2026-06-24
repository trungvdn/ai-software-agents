package retrieve_historical_bug

import (
	"context"
	"fmt"

	"github.com/trungvdn/ai-software-agents/internal/knowledge/domain/historicalbug"
	"github.com/trungvdn/ai-software-agents/shared/embedding"
	"github.com/trungvdn/ai-software-agents/shared/retrieval"
)

type RetrieveHistoricalBugUseCase struct {
	repo     historicalbug.HistoricalBugRepository
	embedder embedding.Embedder
}

func NewRetrieveHistoricalBugUseCase(
	repo historicalbug.HistoricalBugRepository,
	embedder embedding.Embedder,
) *RetrieveHistoricalBugUseCase {
	return &RetrieveHistoricalBugUseCase{
		repo:     repo,
		embedder: embedder,
	}
}

func (r *RetrieveHistoricalBugUseCase) Retrieve(
	ctx context.Context,
	request HistoricalBugRequest,
) (*HistoricalBugResponse, error) {
	embedding, err := r.embedder.Embed(ctx, request.Query)
	if err != nil {
		return nil, err
	}

	bugs, err := r.repo.SearchSimilar(ctx, embedding, request.Limit)
	if err != nil {
		return nil, err
	}

	var results []retrieval.SearchResult
	for _, bug := range bugs {
		results = append(results, retrieval.SearchResult{
			ID: bug.ID.String(),
			Content: fmt.Sprintf(
				"Title: %s\nRootCause: %s\nFix: %s", bug.Title, bug.RootCause, bug.FixSummary),
			Source: "historical_bug",
		})
	}
	return &HistoricalBugResponse{HistoricalBugs: results}, nil
}
