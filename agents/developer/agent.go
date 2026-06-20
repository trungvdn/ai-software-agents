package developer

import (
	"context"

	"github.com/trungvdn/ai-software-agents/domain/historicalbug"
	"github.com/trungvdn/ai-software-agents/domain/reflection"
	"github.com/trungvdn/ai-software-agents/shared/llm"
)

type DeveloperAgent struct {
	reflectionRetriever    reflection.ReflectionRetriever
	historicalBugRetriever historicalbug.HistoricalBugRetriever
	knowledgeBuilder       KnowledgeBuilder

	llm llm.Client
}

func NewDeveloperAgent(
	reflectionRetriever reflection.ReflectionRetriever,
	historicalBugRetriever historicalbug.HistoricalBugRetriever,
	knowledgeBuilder KnowledgeBuilder,
	llm llm.Client,
) *DeveloperAgent {
	return &DeveloperAgent{
		reflectionRetriever:    reflectionRetriever,
		historicalBugRetriever: historicalBugRetriever,
		knowledgeBuilder:       knowledgeBuilder,
		llm:                    llm,
	}
}

func (a *DeveloperAgent) Execute(ctx context.Context, bug string) (*Response, error) {
	// Step 1: Retrieve knowledge context (reflections, historical bugs) from the knowledge base

	reflectionResults, err := a.reflectionRetriever.Retrieve(ctx, bug, 5)
	if err != nil {
		return nil, err
	}
	historicalBugResults, err := a.historicalBugRetriever.Retrieve(ctx, bug, 5)
	if err != nil {
		return nil, err
	}

	searchResults := append(reflectionResults, historicalBugResults...)

	knowledge :=
		a.knowledgeBuilder.Build(
			searchResults,
		)

	// Step 2: Analyze the bug and generate a response using the LLM

	// Step 3: Use tools to retrieve relevant code context

	// Step 4: Generate code patches based on the analysis and code context
	return &Response{
		Knowledge: knowledge,
	}, nil
}
