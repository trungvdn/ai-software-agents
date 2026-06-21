package developer

import (
	"context"

	"github.com/trungvdn/ai-software-agents/domain/historicalbug"
	"github.com/trungvdn/ai-software-agents/domain/reflection"
)

type KnowledgeRetriever interface {
	Retrieve(
		ctx context.Context,
		query string,
		limit int,
	) (*KnowledgeContext, error)
}

type DefaultKnowledgeRetriever struct {
	reflectionRetriever    *reflection.ReflectionRetriever
	historicalBugRetriever *historicalbug.HistoricalBugRetriever
}

func NewDefaultKnowledgeRetriever(
	reflectionRetriever *reflection.ReflectionRetriever,
	historicalBugRetriever *historicalbug.HistoricalBugRetriever) *DefaultKnowledgeRetriever {
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
