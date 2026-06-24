package retrieve_reflection

import (
	"context"

	"github.com/trungvdn/ai-software-agents/internal/knowledge/domain/reflection"
	"github.com/trungvdn/ai-software-agents/shared/embedding"
	"github.com/trungvdn/ai-software-agents/shared/retrieval"
)

type RetrieveReflectionUseCase struct {
	repo     reflection.ReflectionRepository
	embedder embedding.Embedder
}

func NewRetrieveReflectionUseCase(
	repo reflection.ReflectionRepository,
	embedder embedding.Embedder,
) *RetrieveReflectionUseCase {
	return &RetrieveReflectionUseCase{
		repo:     repo,
		embedder: embedder,
	}
}

func (r *RetrieveReflectionUseCase) Retrieve(
	ctx context.Context,
	request ReflectionRequest,
) (*ReflectionResponse, error) {
	embedding, err := r.embedder.Embed(ctx, request.Query)
	if err != nil {
		return nil, err
	}

	reflections, err := r.repo.SearchSimilar(ctx, embedding, request.Limit)
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

	return &ReflectionResponse{Reflections: results}, nil
}
