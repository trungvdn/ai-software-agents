package developer

import (
	"context"
	"errors"
	"log"

	"github.com/trungvdn/ai-software-agents/domain/developer"
	"github.com/trungvdn/ai-software-agents/shared/llm"
)

var ErrFeatureNotImplemented = errors.New("Feature not implemented yet")

type DeveloperAgent struct {
	requirementAnalyzer RequirementAnalyzer
	knowledgeRetriever  KnowledgeRetriever
	codeRetriever       CodeRetriever
	promptBuilder       *DeveloperPromptBuilder
	diffGenerator       DiffGenerator
	patchApplier        PatchApplier

	llm llm.Client
}

func NewDeveloperAgent(
	knowledgeRetriever KnowledgeRetriever,
	codeRetriever CodeRetriever,
	prompt *DeveloperPromptBuilder,
	diffGenerator DiffGenerator,
	patchApplier PatchApplier,
	llm llm.Client,
) *DeveloperAgent {
	return &DeveloperAgent{
		knowledgeRetriever: knowledgeRetriever,
		codeRetriever:      codeRetriever,
		promptBuilder:      prompt,
		diffGenerator:      diffGenerator,
		patchApplier:       patchApplier,
		llm:                llm,
	}
}

func (a *DeveloperAgent) ExecuteBugFix(ctx context.Context, task *developer.DevelopmentTask) (*Response, error) {
	// Step 1: Retrieve knowledge context (reflections, historical bugs) from the knowledge base
	knowledgeContext, err := a.knowledgeRetriever.Retrieve(ctx, task.Description, 10)
	if err != nil {
		return nil, err
	}

	log.Printf("Retrieved %d reflections and %d historical bugs for the bug description", len(knowledgeContext.Reflections), len(knowledgeContext.HistoricalBugs))

	// Step 2: Retrieve relevant code context
	codeContext, err := a.codeRetriever.Retrieve(&RetrievalQuery{
		Query: task.Description,
	})
	if err != nil {
		return nil, err
	}

	log.Printf("Retrieved %d relevant source files for the bug description", len(codeContext.Files))

	// Step 3: Build developer prompt
	prompt := a.promptBuilder.Build(task.Description, knowledgeContext, codeContext)

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
		log.Printf("Patch Candidate %d:\nFile: %s\nOriginal Snippet:\n%s\nModified Snippet:\n%s", i+1, candidate.FilePath, candidate.OriginalSnippet, candidate.ProposedSnippet)
	}

	// Step 8: Generate code patches based on patch candidate
	codePatches, err := a.diffGenerator.Generate(patchCandidate, codeContext)
	if err != nil {
		return nil, err
	}

	for i, patch := range codePatches {
		log.Printf("Generated Code Patch %d:\nFile: %s\nDiff:\n%s", i+1, patch.FilePath, patch.Diff)
	}

	// err = a.patchApplier.Apply(patchCandidate)
	// if err != nil {
	// 	return nil, err
	// }

	return &Response{
		Knowledge:      knowledgeContext,
		CodeContext:    codeContext,
		Analysis:       analysis,
		PatchCandidate: patchCandidate,
		CodePatches:    codePatches,
	}, nil
}

func (a *DeveloperAgent) ExecuteFeature(ctx context.Context, task *developer.DevelopmentTask) (*Response, error) {
	requirementAnalysis, err := a.requirementAnalyzer.Analyze(ctx, task)
	if err != nil {
		log.Printf("Error analyzing requirement: %v", err)
	}
	log.Printf("Requirement Analysis: %+v", requirementAnalysis)

	retrievalQuery := &RetrievalQuery{
		Query:            task.Description,
		CandidateSymbols: requirementAnalysis.CandidateSymbols,
	}
	codeContext, err := a.codeRetriever.Retrieve(retrievalQuery)
	log.Printf("Retrieved %d files for the feature description", len(codeContext.Files))
	if err != nil {
		log.Printf("Error retrieving code context: %v", err)
	}
	return &Response{
		Knowledge:      nil,
		CodeContext:    codeContext,
		PatchCandidate: nil,
		CodePatches:    nil,
	}, nil

}

func (a *DeveloperAgent) ExecuteTest(ctx context.Context, task *developer.DevelopmentTask) (*Response, error) {
	return nil, ErrFeatureNotImplemented
}
