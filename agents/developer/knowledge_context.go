package developer

import (
	"github.com/trungvdn/ai-software-agents/shared/retrieval"
)

type KnowledgeContext struct {
	Reflections    []retrieval.SearchResult
	HistoricalBugs []retrieval.SearchResult
}
