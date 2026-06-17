package reflection

import (
	"context"

	"github.com/trungvdn/ai-software-agents/shared/embedding"
	"github.com/trungvdn/ai-software-agents/shared/rag"
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
) ([]rag.SearchResult, error) {
	embedding, err := r.Embedder.Embed(ctx, text)
	if err != nil {
		return nil, err
	}

	reflections, err := r.ReflectionRepository.SearchSimilar(ctx, embedding, limit)
	if err != nil {
		return nil, err
	}

	var results []rag.SearchResult
	for _, ref := range reflections {
		results = append(results, rag.SearchResult{
			ID:      ref.Reflection.ID.String(),
			Content: ref.Reflection.Content,
			Score:   ref.Similarity,
			Source:  "reflection",
		})
	}

	return results, nil
}
