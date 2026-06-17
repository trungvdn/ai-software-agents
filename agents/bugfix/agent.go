package bugfix

import (
	"github.com/trungvdn/ai-software-agents/planner/plan"
	prompt_context "github.com/trungvdn/ai-software-agents/shared/context"
	"github.com/trungvdn/ai-software-agents/shared/llm"
	"github.com/trungvdn/ai-software-agents/shared/retrieval"
)

type BugFixAgent struct {
	retriever retrieval.Retriever

	reRanker retrieval.ReRanker

	contextBuilder prompt_context.Builder

	planner plan.Planner

	llm llm.Client
}

func NewBugFixAgent(
	retriever retrieval.Retriever,
	reRanker retrieval.ReRanker,
	contextBuilder prompt_context.Builder,
	llm llm.Client,
	planner plan.Planner,
) BugFixAgent {
	return BugFixAgent{
		retriever:      retriever,
		reRanker:       reRanker,
		contextBuilder: contextBuilder,
		llm:            llm,
		planner:        planner,
	}
}
