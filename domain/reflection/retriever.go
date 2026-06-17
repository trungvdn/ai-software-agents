package reflection

import (
	"context"
	"fmt"

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

func (r *ReflectionRetriever) RetrieveSimilar(
	ctx context.Context,
	text string,
	limit int,
) ([]retrieval.SearchResult, error) {
	embedding, err := r.Embedder.Embed(ctx, text)
	if err != nil {
		return nil, err
	}

	reflections, err := r.ReflectionRepository.SearchSimilar(ctx, embedding, limit)
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
			Metadata: map[string]string{
				"importanceScore": fmt.Sprint(ref.Reflection.ImportanceScore),
				"usageCount":      fmt.Sprint(ref.Reflection.UsageCount),
			},
		})
	}

	return results, nil
}
