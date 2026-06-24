package generate_story

import "context"

type GenerateStoryUseCase struct {
	storyGenerator StoryGenerator
}

func NewGenerateStoryUseCase(storyGenerator StoryGenerator) *GenerateStoryUseCase {
	return &GenerateStoryUseCase{
		storyGenerator: storyGenerator,
	}
}

func (g *GenerateStoryUseCase) Generate(
	ctx context.Context,
	request GenerateStoryRequest,
) (*GenerateStoryResponse, error) {
	return g.storyGenerator.Generate(ctx, request)
}
