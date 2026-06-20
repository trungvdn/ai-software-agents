package developer

import (
	"github.com/trungvdn/ai-software-agents/domain/historicalbug"
	"github.com/trungvdn/ai-software-agents/domain/reflection"
)

type KnowledgeContext struct {
	Reflections    []reflection.Reflection
	HistoricalBugs []historicalbug.HistoricalBug
}
