package developer

import (
	"context"

	"github.com/trungvdn/ai-software-agents/shared/llm"
)

type DeveloperAgent struct {
	knowledgeRetriever KnowledgeRetriever
	codeRetriever      CodeRetriever
	promptBuilder      *DeveloperPromptBuilder

	llm llm.Client
}

func NewDeveloperAgent(
	knowledgeRetriever KnowledgeRetriever,
	codeRetriever CodeRetriever,
	prompt *DeveloperPromptBuilder,
	llm llm.Client,
) *DeveloperAgent {
	return &DeveloperAgent{
		knowledgeRetriever: knowledgeRetriever,
		codeRetriever:      codeRetriever,
		promptBuilder:      prompt,
		llm:                llm,
	}
}

func (a *DeveloperAgent) Execute(ctx context.Context, bug string) (*Response, error) {
	// Step 1: Retrieve knowledge context (reflections, historical bugs) from the knowledge base
	knowledgeContext, err := a.knowledgeRetriever.Retrieve(ctx, bug, 10)
	if err != nil {
		return nil, err
	}

	// Step 2: Retrieve relevant code context
	codeContext, err := a.codeRetriever.Retrieve(bug)
	if err != nil {
		return nil, err
	}

	// Step 3: Build developer prompt
	prompt := a.promptBuilder.Build(bug, knowledgeContext, codeContext)

	// Step 4: Reasoning fix suggestion
	llmResponse, err := a.llm.Chat(ctx, prompt)
	if err != nil {
		return nil, err
	}

	// Step 4: Parse Response
	analysis, err := ParseAnalysis(llmResponse)
	if err != nil {
		return nil, err
	}

	// Step 5: Build patch plan prompt
	patchCandidatePrompt := a.promptBuilder.BuildPatchCandidate(
		analysis, codeContext,
	)

	// Step 6: Reasoning build patch plan
	llmPatchPlanResponse, err := a.llm.Chat(ctx, patchCandidatePrompt)
	if err != nil {
		return nil, err
	}

	// Step 7: Parse Response PatchPlan
	patchPlan, err := ParsePatchCadidate(llmPatchPlanResponse)
	if err != nil {
		return nil, err
	}

	//Step 8: Generate code patches based on patch plan
	return &Response{
		Knowledge:   knowledgeContext,
		CodeContext: codeContext,
		Analysis:    analysis,
		PatchPlan:   patchPlan,
	}, nil
}
