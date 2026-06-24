package retrieve_knowledge

import (
	"context"

	"github.com/trungvdn/ai-software-agents/shared/retrieval"
)

type ReflectionRetriever interface {
	Retrieve(
		ctx context.Context,
		query string,
		topK int,
	) ([]*retrieval.SearchResult, error)
}

type HistoricalBugRetriever interface {
	Retrieve(
		ctx context.Context,
		query string,
		topK int,
	) ([]*retrieval.SearchResult, error)
}

type RetrieveKnowledgeUseCase struct {
	retrieveHistoricalBugUseCase HistoricalBugRetriever
	retrieveReflectionUseCase    ReflectionRetriever
}

func NewRetrieveKnowledgeUseCase(
	retrievehistoricalbug HistoricalBugRetriever,
	retrievereflection ReflectionRetriever,
) *RetrieveKnowledgeUseCase {
	return &RetrieveKnowledgeUseCase{
		retrieveHistoricalBugUseCase: retrievehistoricalbug,
		retrieveReflectionUseCase:    retrievereflection,
	}
}

func (r *RetrieveKnowledgeUseCase) Retrieve(
	ctx context.Context,
	query string,
	limit int,
) (*Response, error) {
	reflections, err := r.retrieveReflectionUseCase.Retrieve(ctx, query, limit)
	if err != nil {
		return nil, err
	}
	historicalBugs, err := r.retrieveHistoricalBugUseCase.Retrieve(ctx, query, limit)
	if err != nil {
		return nil, err
	}

	return &Response{
		Reflections:    reflections,
		HistoricalBugs: historicalBugs,
	}, nil
}
