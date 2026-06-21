package developer

import (
	"context"

	"github.com/trungvdn/ai-software-agents/shared/retrieval"
)

type KnowledgeRetriever interface {
	Retrieve(
		ctx context.Context,
		query string,
		limit int,
	) (*KnowledgeContext, error)
}

type DefaultKnowledgeRetriever struct {
	reflectionRetriever    retrieval.Retriever
	historicalBugRetriever retrieval.Retriever
}

func NewDefaultKnowledgeRetriever(
	reflectionRetriever retrieval.Retriever,
	historicalBugRetriever retrieval.Retriever) *DefaultKnowledgeRetriever {
	return &DefaultKnowledgeRetriever{
		reflectionRetriever:    reflectionRetriever,
		historicalBugRetriever: historicalBugRetriever,
	}

}

func (b *DefaultKnowledgeRetriever) Retrieve(
	ctx context.Context,
	query string,
	limit int,
) (*KnowledgeContext, error) {
	reflections, err := b.reflectionRetriever.Retrieve(ctx, query, limit)
	if err != nil {
		return nil, err
	}
	historicalBugs, err := b.historicalBugRetriever.Retrieve(ctx, query, limit)
	if err != nil {
		return nil, err
	}
	return &KnowledgeContext{
		Reflections:    reflections,
		HistoricalBugs: historicalBugs,
	}, nil
}
