package context

import (
	"context"
	"fmt"

	"github.com/trungvdn/ai-software-agents/shared/retrieval"
)

type ReflectionContextBuilder struct {
}

func NewReflectionContextBuilder() *ReflectionContextBuilder {
	return &ReflectionContextBuilder{}
}

func (b *ReflectionContextBuilder) Build(
	ctx context.Context,
	results []retrieval.SearchResult,
) (Context, error) {
	/*
			=== CONTEXT ===

		Relevant Reflections:

		1. Always check nil after repository call

		2. Validate returned object before access

		3. Always handle transaction rollback
	*/
	context := "=== CONTEXT ===\n\nRelevant Reflections:\n\n"
	for i, res := range results {
		context += fmt.Sprintf("%d. %s\n\n", i+1, res.Content)
	}
	return Context{Content: context}, nil
}
