package retrieve_knowledge

import (
	"context"

	"github.com/trungvdn/ai-software-agents/internal/knowledge/application/retrieve_historical_bug"
	"github.com/trungvdn/ai-software-agents/internal/knowledge/application/retrieve_reflection"
)

type ReflectionRetriever interface {
	Retrieve(
		ctx context.Context,
		request retrieve_reflection.ReflectionRequest,
	) (*retrieve_reflection.ReflectionResponse, error)
}

type HistoricalBugRetriever interface {
	Retrieve(
		ctx context.Context,
		request retrieve_historical_bug.HistoricalBugRequest,
	) (*retrieve_historical_bug.HistoricalBugResponse, error)
}

type RetrieveKnowledgeUseCase struct {
	retrieveHistoricalBugUseCase HistoricalBugRetriever
	retrieveReflectionUseCase    ReflectionRetriever
}

func NewRetrieveKnowledgeUseCase(
	retrievereflection ReflectionRetriever,
	retrievehistoricalbug HistoricalBugRetriever,
) *RetrieveKnowledgeUseCase {
	return &RetrieveKnowledgeUseCase{
		retrieveHistoricalBugUseCase: retrievehistoricalbug,
		retrieveReflectionUseCase:    retrievereflection,
	}
}

func (r *RetrieveKnowledgeUseCase) Retrieve(
	ctx context.Context,
	request KnowledgeRequest,
) (*KnowledgeContextResponse, error) {
	reflections, err := r.retrieveReflectionUseCase.Retrieve(ctx, retrieve_reflection.ReflectionRequest{
		Query: request.Query,
		Limit: request.Limit,
	})
	if err != nil {
		return nil, err
	}
	historicalBugs, err := r.retrieveHistoricalBugUseCase.Retrieve(ctx,
		retrieve_historical_bug.HistoricalBugRequest{
			Query: request.Query,
			Limit: request.Limit,
		})
	if err != nil {
		return nil, err
	}

	return &KnowledgeContextResponse{
		Reflections:    reflections.Reflections,
		HistoricalBugs: historicalBugs.HistoricalBugs,
	}, nil
}
