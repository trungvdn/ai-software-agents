package developer

import (
	"github.com/trungvdn/ai-software-agents/shared/retrieval"
	"github.com/trungvdn/ai-software-agents/shared/tools"
)

type KnowledgeContext struct {
	Reflections    []retrieval.SearchResult
	HistoricalBugs []retrieval.SearchResult
	CodeFiles      []*tools.FileContent
}
