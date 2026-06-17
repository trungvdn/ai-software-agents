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
	// Implementation for fixing the bug

	results, err := a.retriever.Retrieve(ctx, bugDescription, 10)
	if err != nil {
		log.Printf("Error retrieving context: %v", err)
		return nil, err
	}

	resultsRerank, err := a.reRanker.ReRank(ctx, bugDescription, results)
	if err != nil {
		log.Printf("Error re-ranking results: %v", err)
		return nil, err
	}

	promptContext, err := a.contextBuilder.Build(ctx, resultsRerank)
	if err != nil {
		log.Printf("Error building context: %v", err)
		return nil, err
	}

	responseLLM, err := a.llm.Chat(
		ctx,
		promptContext.Content,
	)
	if err != nil {
		log.Printf("Error getting response from LLM: %v", err)
		return nil, err
	}
	response := &Response{
		Analysis: responseLLM,
	}
	fmt.Println("LLM Response:", responseLLM)
	return response, nil
}
