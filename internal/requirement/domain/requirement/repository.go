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

	type Repository interface {
    SaveAggregate(
        ctx context.Context,
        aggregate Aggregate,
    ) error

    GetAggregate(
        ctx context.Context,
        id string,
    ) (*Aggregate, error)
}
}
