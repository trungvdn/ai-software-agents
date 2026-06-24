package generate_epic

import (
	"context"
)

type EpicGenerator interface {
	Generate(
		ctx context.Context,
		request GenerateEpicRequest,
	) (*GenerateEpicResponse, error)
}

func NewGenerateEpicUseCase(epicGenerator EpicGenerator) *GenerateEpicUseCase {
	return &GenerateEpicUseCase{
		epicGenerator: epicGenerator,
	}
}

func (g *GenerateEpicUseCase) Generate(
	ctx context.Context,
	request GenerateEpicRequest,
) (*GenerateEpicResponse, error) {
	return g.epicGenerator.Generate(ctx, request)
}
