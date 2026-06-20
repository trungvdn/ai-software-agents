package developer

import (
	"context"
	"fmt"

	"github.com/trungvdn/ai-software-agents/shared/llm"
)

type DeveloperAgent struct {
	knowledgeRetriever KnowledgeRetriever
	codeRetriever      CodeRetriever

	llm llm.Client
}

func NewDeveloperAgent(
	knowledgeRetriever KnowledgeRetriever,
	llm llm.Client,
) *DeveloperAgent {
	return &DeveloperAgent{
		knowledgeRetriever: knowledgeRetriever,
		llm:                llm,
	}
}

func (a *DeveloperAgent) Execute(ctx context.Context, bug string) (*Response, error) {
	// Step 1: Retrieve knowledge context (reflections, historical bugs) from the knowledge base

	knowledgeContext, err := a.knowledgeRetriever.Retrieve(ctx, bug, 10)
	if err != nil {
		return nil, err
	}

	// Step 2: Use tools to retrieve relevant code context
	codeRetriever, err := a.codeRetriever.Retrieve(ctx, bug)

	// Step 3: Analyze the bug and generate a response using the LLM

	fmt.Println(codeRetriever.Files)

	// Step 4: Generate code patches based on the analysis and code context
	return &Response{
		Knowledge: knowledgeContext,
	}, nil
}
