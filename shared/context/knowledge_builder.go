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
	builder.WriteString("=== CONTEXT ===\n\n")
	builder.WriteString("Relevant Knowledge:\n\n")
	builder.WriteString("[Reflection]\n")
	for i, res := range results {
		if res.Source == "reflection" {
			builder.WriteString(fmt.Sprintf("%d. %s\n\n", i+1, res.Content))
		}
	}
	builder.WriteString("[Historical Bug]\n")
	for i, res := range results {
		if res.Source == "historical_bug" {
			builder.WriteString(fmt.Sprintf("%d. %s\n\n", i+1, res.Content))
		}
	}
	return PromptContext{Content: builder.String()}, nil
}
