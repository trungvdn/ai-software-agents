package developer

import (
	"context"

	developer_domain "github.com/trungvdn/ai-software-agents/domain/developer"
	"github.com/trungvdn/ai-software-agents/domain/requirement"
	"github.com/trungvdn/ai-software-agents/shared/llm"
)

type RequirementAnalyzer interface {
	Analyze(
		ctx context.Context,
		task *developer_domain.DevelopmentTask,
	) (*requirement.RequirementAnalysis, error)
}

type DefaultRequirementAnalyzer struct {
	llm llm.Client
}
