package epic

import "context"

type EpicRepository interface {
	Save(
		ctx context.Context,
		epic Epic,
	) error

	FindByRequirementID(
		ctx context.Context,
		requirementID string,
	) ([]Epic, error)
}
