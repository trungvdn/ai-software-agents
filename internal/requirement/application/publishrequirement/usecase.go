package publishrequirement

import (
	"context"

	"github.com/trungvdn/ai-software-agents/internal/requirement/domain/requirement"
)

type Publisher interface {
	Publish(
		ctx context.Context,
		requirement requirement.Requirement,
	) error
}

type PublishRequirementUseCase struct {
}

func NewPublishRequirementUseCase() *PublishRequirementUseCase {
	return &PublishRequirementUseCase{}
}

func (p *PublishRequirementUseCase) Publish(
	ctx context.Context,
	requirement requirement.Requirement,
) error {
	return nil
}
