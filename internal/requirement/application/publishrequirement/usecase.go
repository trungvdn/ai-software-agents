package publish_requirement

import (
	"context"
)

type Publisher interface {
	Publish(
		ctx context.Context,
		request PublishRequirementRequest,
	) error
}

type PublishRequirementUseCase struct {
	publisher Publisher
}

func NewPublishRequirementUseCase(publisher Publisher) *PublishRequirementUseCase {
	return &PublishRequirementUseCase{
		publisher: publisher,
	}
}

func (p *PublishRequirementUseCase) Publish(
	ctx context.Context,
	request PublishRequirementRequest,
) error {
	return p.publisher.Publish(ctx, request)
}
