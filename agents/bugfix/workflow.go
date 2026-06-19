package bugfix

import (
	"context"
	"fmt"
	"log"
)

/* Bug Description

      ↓

Retrieve

      ↓

ReRank

      ↓

Build Context

      ↓

Build Prompt

      ↓

LLM

      ↓

Response
*/

func (a *BugFixAgent) FixBug(
	ctx context.Context,
	bugDescription string,
) (*Response, error) {
	// Retrieve relevant context
	results, err := a.retriever.Retrieve(ctx, bugDescription, 10)
	if err != nil {
		log.Printf("Error retrieving context: %v", err)
		return nil, err
	}

	// Re-rank results based on relevance to the bug description
	resultsRerank, err := a.reRanker.ReRank(ctx, bugDescription, results)
	if err != nil {
		log.Printf("Error re-ranking results: %v", err)
		return nil, err
	}

	// Build context for LLM
	promptContext, err := a.contextBuilder.Build(ctx, resultsRerank)
	if err != nil {
		log.Printf("Error building context: %v", err)
		return nil, err
	}

	// Build prompt for LLM
	promptBuilder := &PromptBuilder{}
	prompt := promptBuilder.Build(
		bugDescription,
		promptContext.Content,
	)

	fmt.Println("Prompt for LLM:", prompt)

	// Get response from LLM
	responseLLM, err := a.llm.Chat(
		ctx,
		prompt,
	)
	if err != nil {
		log.Printf("Error getting response from LLM: %v", err)
		return nil, err
	}

	// Parse LLM response (for simplicity, we just return the raw response here)
	response := &Response{
		Analysis: responseLLM,
	}
	fmt.Println("LLM Response:", responseLLM)

	// Change planning based on LLM response

	plan, err := a.planner.Plan(ctx, bugDescription, responseLLM)
	if err != nil {
		log.Printf("Error planning changes: %v", err)
		return nil, err
	}
	response.Plan = plan

	// Code generation based on change plan
	codePatches, err := a.coder.GeneratePatches(ctx, bugDescription, plan)
	if err != nil {
		log.Printf("Error generating code: %v", err)
		return nil, err
	}

	fmt.Println("Generated Code Patches:", codePatches)

	response.CodePatches = codePatches

	return response, nil
}
