package generate_requirement

import "context"

type RequirementGenerator interface {
	Generate(
		ctx context.Context,
		request GenerateRequirementRequest,
	) (*GenerateRequirementResponse, error)
}
