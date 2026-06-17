package context

import (
	"context"
	"fmt"
	"strings"

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
) (PromptContext, error) {
	var builder strings.Builder
	builder.WriteString("=== CONTEXT ===\n\n")
	builder.WriteString("Relevant Reflections:\n\n")
	for i, res := range results {
		builder.WriteString(fmt.Sprintf("%d. %s\n\n", i+1, res.Content))
	}
	return PromptContext{Content: builder.String()}, nil
}
