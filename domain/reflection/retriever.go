package reflection

import (
	"context"

	"github.com/trungvdn/ai-software-agents/shared/embedding"
	"github.com/trungvdn/ai-software-agents/shared/retrieval"
)

type ReflectionRetriever struct {
	repo     ReflectionRepository
	embedder embedding.Embedder
}

func NewReflectionRetriever(
	repo ReflectionRepository,
	embedder embedding.Embedder,
) *ReflectionRetriever {
	return &ReflectionRetriever{
		repo:     repo,
		embedder: embedder,
	}
}

func (r *ReflectionRetriever) Retrieve(
	ctx context.Context,
	query string,
	topK int,
) ([]retrieval.SearchResult, error) {
	embedding, err := r.embedder.Embed(ctx, query)
	if err != nil {
		return nil, err
	}

	reflections, err := r.repo.SearchSimilar(ctx, embedding, topK)
	if err != nil {
		return nil, err
	}

	var results []retrieval.SearchResult
	for _, ref := range reflections {
		results = append(results, retrieval.SearchResult{
			ID:      ref.Reflection.ID.String(),
			Content: ref.Reflection.Content,
			Score:   ref.Similarity,
			Source:  "reflection",
			Metadata: retrieval.SearchMetadata{
				ImportanceScore: ref.Reflection.ImportanceScore,
				UsageCount:      ref.Reflection.UsageCount,
			},
		})
	}

	return results, nil
}
