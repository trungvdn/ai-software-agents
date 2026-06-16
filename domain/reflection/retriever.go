package reflection

import (
	"context"

	"github.com/trungvdn/ai-software-agents/shared/embedding"
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
) ([]Reflection, error) {
	embedding, err := r.Embedder.Embed(ctx, text)
	if err != nil {
		return nil, err
	}

	return r.ReflectionRepository.SearchSimilar(ctx, embedding, limit)
}
