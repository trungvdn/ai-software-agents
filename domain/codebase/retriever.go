package codebase

import (
	"context"

	"github.com/trungvdn/ai-software-agents/shared/embedding"
	"github.com/trungvdn/ai-software-agents/shared/retrieval"
)

type CodeDocumentRetriever struct {
	repo     CodeDocumentRepository
	embedder embedding.Embedder
}

func NewCodeDocumentRetriever(
	repo CodeDocumentRepository,
	embedder embedding.Embedder,
) *CodeDocumentRetriever {
	return &CodeDocumentRetriever{
		repo:     repo,
		embedder: embedder,
	}
}

func (r *CodeDocumentRetriever) Retrieve(
	ctx context.Context,
	query string,
	topK int,
) ([]retrieval.SearchResult, error) {
	embedding, err := r.embedder.Embed(ctx, query)
	if err != nil {
		return nil, err
	}

	codebases, err := r.repo.SearchSimilar(ctx, embedding, topK)
	if err != nil {
		return nil, err
	}

	var results []retrieval.SearchResult
	for _, cb := range codebases {
		results = append(results, retrieval.SearchResult{
			Source:  "codebase",
			ID:      cb.ID.String(),
			Content: cb.Content,
		})
	}

	return results, nil
}
