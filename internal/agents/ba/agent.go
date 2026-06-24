package ba

import (
	"context"

	"github.com/trungvdn/ai-software-agents/internal/requirement/application/generate_requirement_package"
	"github.com/trungvdn/ai-software-agents/internal/requirement/application/publish_requirement"
)

type BAAgent struct {
	generateRequirementPackageUseCase *generate_requirement_package.GenerateRequirementPackageUseCase
	publishRequirementUseCase         *publish_requirement.PublishRequirementUseCase
}

func NewBAAgent(
	generateRequirementPackageUseCase *generate_requirement_package.GenerateRequirementPackageUseCase,
	publishRequirementUseCase *publish_requirement.PublishRequirementUseCase,
) *BAAgent {
	return &BAAgent{
		generateRequirementPackageUseCase: generateRequirementPackageUseCase,
		publishRequirementUseCase:         publishRequirementUseCase,
	}
}

func (a *BAAgent) Execute(
	ctx context.Context,
	request Request,
) (*Response, error) {
	// Generate the requirement package
	generateRequirementPackageResponse, err := a.generateRequirementPackageUseCase.Execute(ctx, generate_requirement_package.GenerateRequirementPackageRequest{
		Idea: request.Idea,
	})
	if err != nil {
		return nil, err
	}

	// Publish the requirement package
	err = a.publishRequirementUseCase.Publish(ctx, publish_requirement.PublishRequirementRequest{
		RequirementAggregate: generateRequirementPackageResponse.RequirementAggregate,
	})
	if err != nil {
		return nil, err
	}
	return &Response{
		RequirementAggregate: *generateRequirementPackageResponse.RequirementAggregate,
	}, nil
}
