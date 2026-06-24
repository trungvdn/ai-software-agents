package retrieve_knowledge

import "github.com/trungvdn/ai-software-agents/shared/retrieval"

type Response struct {
	Reflections    []*retrieval.SearchResult
	HistoricalBugs []*retrieval.SearchResult
}
