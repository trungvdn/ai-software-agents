package story

import "context"

type StoryRepository interface {
	Save(
		ctx context.Context,
		story Story,
	) error

	FindByEpicID(
		ctx context.Context,
		epicID string,
	) ([]Story, error)
}
