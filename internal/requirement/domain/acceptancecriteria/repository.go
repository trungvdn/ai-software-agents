package acceptance_criteria

import "context"

type AcceptanceCriteriaRepository interface {
	Save(
		ctx context.Context,
		acceptanceCriteria AcceptanceCriteria,
	) error

	FindByStoryID(
		ctx context.Context,
		storyID string,
	) (*AcceptanceCriteria, error)
}
