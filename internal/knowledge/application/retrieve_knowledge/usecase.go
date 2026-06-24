package retrieve_knowledge

import (
	"context"

	"github.com/trungvdn/ai-software-agents/internal/knowledge/application/retrieve_historical_bug"
	"github.com/trungvdn/ai-software-agents/internal/knowledge/application/retrieve_reflection"
)

type RetrieveKnowledgeUseCase struct {
	retrieveHistoricalBugUseCase retrieve_historical_bug.RetrieveHistoricalBugUseCase
	retrieveReflectionUseCase    retrieve_reflection.RetrieveReflectionUseCase
}

func NewRetrieveKnowledgeUseCase(
	retrievehistoricalbug retrieve_historical_bug.RetrieveHistoricalBugUseCase,
	retrievereflection retrieve_reflection.RetrieveReflectionUseCase,
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
) (*KnowledgeContext, error) {
	reflections, err := r.retrieveReflectionUseCase.Retrieve(ctx, query, limit)
	if err != nil {
		return nil, err
	}
	historicalBugs, err := r.retrieveHistoricalBugUseCase.Retrieve(ctx, query, limit)
	if err != nil {
		return nil, err
	}

	return &KnowledgeContext{
		Reflections:    reflections,
		HistoricalBugs: historicalBugs,
	}, nil
}
