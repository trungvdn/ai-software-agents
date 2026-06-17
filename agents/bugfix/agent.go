package bugfix

import (
	prompt_context "github.com/trungvdn/ai-software-agents/shared/context"
	"github.com/trungvdn/ai-software-agents/shared/llm"
	"github.com/trungvdn/ai-software-agents/shared/retrieval"
)

type BugFixAgent struct {
	retriever retrieval.Retriever

	reRanker retrieval.ReRanker

	contextBuilder prompt_context.Builder

	llm llm.Client
}

func NewBugFixAgent(
	retriever retrieval.Retriever,
	reRanker retrieval.ReRanker,
	contextBuilder prompt_context.Builder,
	llm llm.Client,
) BugFixAgent {
	return BugFixAgent{
		retriever:      retriever,
		reRanker:       reRanker,
		contextBuilder: contextBuilder,
		llm:            llm,
	}
}
