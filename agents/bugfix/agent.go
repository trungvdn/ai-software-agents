package bugfix

import (
	"github.com/trungvdn/ai-software-agents/agents/coder"
	"github.com/trungvdn/ai-software-agents/agents/planner"
	prompt_context "github.com/trungvdn/ai-software-agents/shared/context"
	"github.com/trungvdn/ai-software-agents/shared/llm"
	"github.com/trungvdn/ai-software-agents/shared/retrieval"
)

type BugFixAgent struct {
	retriever retrieval.Retriever

	reRanker retrieval.ReRanker

	contextBuilder prompt_context.Builder

	planner planner.Planner

	coder coder.Coder

	llm llm.Client
}

func NewBugFixAgent(
	retriever retrieval.Retriever,
	reRanker retrieval.ReRanker,
	contextBuilder prompt_context.Builder,
	llm llm.Client,
	planner planner.Planner,
	coder coder.Coder,
) BugFixAgent {
	return BugFixAgent{
		retriever:      retriever,
		reRanker:       reRanker,
		contextBuilder: contextBuilder,
		llm:            llm,
		planner:        planner,
		coder:          coder,
	}
}
