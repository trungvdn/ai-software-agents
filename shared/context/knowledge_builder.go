package context

import (
	"context"
	"fmt"
	"strings"

	"github.com/trungvdn/ai-software-agents/shared/retrieval"
)

type KnowledgeContextBuilder struct {
}

func NewKnowledgeContextBuilder() *KnowledgeContextBuilder {
	return &KnowledgeContextBuilder{}
}

func (b *KnowledgeContextBuilder) Build(
	ctx context.Context,
	results []retrieval.SearchResult,
) (PromptContext, error) {
	/*
		=== CONTEXT ===
		Relevant Knowledge:

		[reflection]
		Always check nil after repository call

		[historical_bug]
		Title: Null Pointer Exception in UserService
		RootCause: Repository returned nil
		Fix: Add nil check
	*/
	var builder strings.Builder
	builder.WriteString("[Reflection]\n")
	grouped := make(map[string][]retrieval.SearchResult)

	for _, result := range results {
		grouped[result.Source] = append(grouped[result.Source], result)
	}

	for reflectionIndex, res := range grouped["reflection"] {
		builder.WriteString(fmt.Sprintf("%d. %s\n\n", reflectionIndex+1, res.Content))
	}
	builder.WriteString("[Historical Bug]\n")
	for historicalBugIndex, res := range grouped["historical_bug"] {
		builder.WriteString(fmt.Sprintf("%d. %s\n\n", historicalBugIndex+1, res.Content))
	}

	return PromptContext{Content: builder.String()}, nil
}
