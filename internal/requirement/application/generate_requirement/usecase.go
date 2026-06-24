package generate_requirement

import (
	"context"
)

type RequirementGenerator interface {
	Generate(
		ctx context.Context,
		request GenerateRequirementRequest,
	) (*GenerateRequirementResponse, error)
}

type GenerateRequirementUseCase struct {
	generator RequirementGenerator
}

func NewGenerateRequirementUseCase(generator RequirementGenerator) *GenerateRequirementUseCase {
	return &GenerateRequirementUseCase{
		generator: generator,
	}
}

func (g *GenerateRequirementUseCase) Generate(
	ctx context.Context,
	request GenerateRequirementRequest,
) (*GenerateRequirementResponse, error) {
	return g.generator.Generate(ctx, request)
}
