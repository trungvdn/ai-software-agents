package generate_story

import "context"

type StoryGenerator interface {
	Generate(
		ctx context.Context,
		request GenerateStoryRequest,
	) (*GenerateStoryResponse, error)
}
