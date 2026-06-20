package developer

import (
	"github.com/trungvdn/ai-software-agents/shared/retrieval"
)

type KnowledgeBuilder interface {
	Build(
		searchResults []retrieval.SearchResult,
	) *KnowledgeContext
}

type DefaultKnowledgeBuilder struct {
}

func (b *DefaultKnowledgeBuilder) Build(
	searchResults []retrieval.SearchResult,
) *KnowledgeContext {

	return &KnowledgeContext{
		Results: searchResults,
	}
}
