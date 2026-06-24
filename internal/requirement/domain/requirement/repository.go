package requirement

import "context"

type RequirementRepository interface {
	Save(
		ctx context.Context,
		requirement Requirement,
	) error

	GetByID(
		ctx context.Context,
		id string,
	) (*Requirement, error)

	UpdateStatus(
		ctx context.Context,
		id string,
		status Status,
	) error
}
