package generate_story

import "context"

type StoryGenerator interface {
	Generate(
		ctx context.Context,
		request GenerateStoryRequest,
	) (*GenerateStoryResponse, error)
}
type GenerateStoryUseCase struct {
	storyGenerator StoryGenerator
}

func NewGenerateRequirementUseCase(storyGenerator StoryGenerator) *GenerateStoryUseCase {
	return &GenerateStoryUseCase{
		storyGenerator: storyGenerator,
	}
}

func (g *GenerateStoryUseCase) Generate(
	ctx context.Context,
	request GenerateStoryRequest,
) (*GenerateStoryResponse, error) {
	return g.Generate(ctx, request)
}
