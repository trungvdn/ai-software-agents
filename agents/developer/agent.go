package developer

import (
	"context"
	"log"

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

	log.Printf("Retrieved %d reflections and %d historical bugs for the bug description", len(knowledgeContext.Reflections), len(knowledgeContext.HistoricalBugs))

	// Step 2: Retrieve relevant code context
	codeContext, err := a.codeRetriever.Retrieve(bug)
	if err != nil {
		return nil, err
	}

	log.Printf("Retrieved %d relevant source files for the bug description", len(codeContext.Files))

	// Step 3: Build developer prompt
	prompt := a.promptBuilder.Build(bug, knowledgeContext, codeContext)

	log.Printf("Constructed prompt for LLM:\n%s", prompt)

	// Step 4: Reasoning fix suggestion
	llmResponse, err := a.llm.Chat(ctx, prompt)
	if err != nil {
		return nil, err
	}

	log.Printf("LLM Response for analysis:\n%s", llmResponse)

	// Step 4: Parse Response
	analysis, err := ParseAnalysis(llmResponse)
	if err != nil {
		return nil, err
	}

	log.Printf("LLM Analysis Suggested Fix:\n%v", analysis)

	// Step 5: Build patch plan prompt
	patchCandidatePrompt := a.promptBuilder.BuildPatchCandidate(
		analysis, codeContext,
	)

	// Step 6: Reasoning build patch plan
	llmPatchCandidateResponse, err := a.llm.Chat(ctx, patchCandidatePrompt)
	if err != nil {
		return nil, err
	}

	// Step 7: Parse Response PatchCandidate
	patchCandidate, err := ParsePatchCandidate(llmPatchCandidateResponse)
	if err != nil {
		return nil, err
	}

	for i, candidate := range patchCandidate {
		log.Printf("Patch Candidate %d:\nFile: %s\nOriginalSnippet: %d\nProposedSnippet:\n%s\n  %s\nReason:", i+1, candidate.FilePath, candidate.OriginalSnippet, candidate.ProposedSnippet, candidate.Reason)
	}

	//Step 8: Generate code patches based on patch plan
	return &Response{
		Knowledge:      knowledgeContext,
		CodeContext:    codeContext,
		Analysis:       analysis,
		PatchCandidate: patchCandidate[0],
	}, nil
}
