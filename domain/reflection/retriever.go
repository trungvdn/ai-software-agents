package reflection

import (
	"context"

	"github.com/trungvdn/ai-software-agents/shared/embedding"
	"github.com/trungvdn/ai-software-agents/shared/retrieval"
)

type ReflectionRetriever struct {
	ReflectionRepository ReflectionRepository
	Embedder             embedding.Embedder
}

func NewReflectionRetriever(
	repo ReflectionRepository,
	embedder embedding.Embedder,
) *ReflectionRetriever {
	return &ReflectionRetriever{
		ReflectionRepository: repo,
		Embedder:             embedder,
	}
}

func (r *ReflectionRetriever) Retrieve(
	ctx context.Context,
	query string,
	topK int,
) ([]retrieval.SearchResult, error) {
	embedding, err := r.Embedder.Embed(ctx, query)
	if err != nil {
		return nil, err
	}

	reflections, err := r.ReflectionRepository.SearchSimilar(ctx, embedding, topK)
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
