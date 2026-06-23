package developer

import (
	"context"
	"log"

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
	llm           llm.Client
	promptBuilder *RequirementPromptBuilder
}

func NewDefaultRequirementAnalyzer(
	llm llm.Client,
	promptBuilder *RequirementPromptBuilder,
) *DefaultRequirementAnalyzer {
	return &DefaultRequirementAnalyzer{
		llm:           llm,
		promptBuilder: promptBuilder,
	}
}

func (a *DefaultRequirementAnalyzer) Analyze(
	ctx context.Context,
	task *developer_domain.DevelopmentTask,
) (*requirement.RequirementAnalysis, error) {
	prompt := a.promptBuilder.Build(task)
	llmResponse, err := a.llm.Chat(ctx, prompt)
	if err != nil {
		return nil, err
	}

	requirementAnalysis, err := ParseRequirementAnalysis(llmResponse)
	if err != nil {
		return nil, err
	}
	log.Printf("Requirement analysis result: %+v", requirementAnalysis)
	return requirementAnalysis, nil
}
