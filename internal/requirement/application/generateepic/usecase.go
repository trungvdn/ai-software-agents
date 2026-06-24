package generate_epic

import (
	"context"
)

type GenerateEpicUseCase struct {
	epicGenerator EpicGenerator
}

type EpicGenerator interface {
	Generate(
		ctx context.Context,
		request GenerateEpicRequest,
	) (*GenerateEpicResponse, error)
}

func NewEpicGeneratorUseCase(epicGenerator EpicGenerator) *GenerateEpicUseCase {
	return &GenerateEpicUseCase{
		epicGenerator: epicGenerator,
	}
}

func (g *GenerateEpicUseCase) Generate(
	ctx context.Context,
	request GenerateEpicRequest,
) (*GenerateEpicResponse, error) {
	return g.Generate(ctx, request)
}
