package developer

import (
	"github.com/trungvdn/ai-software-agents/domain/historicalbug"
	"github.com/trungvdn/ai-software-agents/domain/reflection"
)

type KnowledgeBuilder interface {
	Build(
		reflections []reflection.Reflection,
		historicalBugs []historicalbug.HistoricalBug,
	) *KnowledgeContext
}

type DefaultKnowledgeBuilder struct {
}

func (b *DefaultKnowledgeBuilder) Build(
	reflections []reflection.Reflection,
	historicalBugs []historicalbug.HistoricalBug,
) *KnowledgeContext {

	return &KnowledgeContext{
		Reflections:    reflections,
		HistoricalBugs: historicalBugs,
	}
}
