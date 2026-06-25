package generate_requirement_package

import (
	"context"

	"golang.org/x/sync/errgroup"

	"github.com/trungvdn/ai-software-agents/internal/requirement/application/generate_epic"
	"github.com/trungvdn/ai-software-agents/internal/requirement/application/generate_requirement"
	"github.com/trungvdn/ai-software-agents/internal/requirement/application/generate_story"
	"github.com/trungvdn/ai-software-agents/internal/requirement/domain/requirement"
	"github.com/trungvdn/ai-software-agents/internal/requirement/domain/story"
)

type GenerateRequirementPackageUseCase struct {
	generateRequirementUseCase *generate_requirement.GenerateRequirementUseCase
	generateEpicUseCase        *generate_epic.GenerateEpicUseCase
	generateStoryUseCase       *generate_story.GenerateStoryUseCase
}

func NewGenerateRequirementPackageUseCase(
	generateRequirementUseCase *generate_requirement.GenerateRequirementUseCase,
	generateEpicUseCase *generate_epic.GenerateEpicUseCase,
	generateStoryUseCase *generate_story.GenerateStoryUseCase,
) *GenerateRequirementPackageUseCase {
	return &GenerateRequirementPackageUseCase{
		generateRequirementUseCase: generateRequirementUseCase,
		generateEpicUseCase:        generateEpicUseCase,
		generateStoryUseCase:       generateStoryUseCase,
	}
}

func (u *GenerateRequirementPackageUseCase) Execute(
	ctx context.Context,
	request GenerateRequirementPackageRequest,
) (*GenerateRequirementPackageResponse, error) {
	requirementResp, err :=
		u.generateRequirementUseCase.Generate(ctx, generate_requirement.GenerateRequirementRequest{Idea: request.Idea})
	if err != nil {
		return nil, err
	}
	epicResp, err :=
		u.generateEpicUseCase.Generate(ctx, generate_epic.GenerateEpicRequest{Requirement: requirementResp.Requirement})
	if err != nil {
		return nil, err
	}
	group, ctx := errgroup.WithContext(ctx)
	results := make([][]story.Story, len(epicResp.Epics))
	for i, epic := range epicResp.Epics {
		idx := i
		epicItem := epic
		group.Go(func() error {
			storyResp, err := u.generateStoryUseCase.Generate(ctx, generate_story.GenerateStoryRequest{Epic: epicItem})
			if err != nil {
				return err
			}
			results[idx] = storyResp.Stories
			return nil
		})
	}

	if err := group.Wait(); err != nil {
		return nil, err
	}

	epicAggregates := make([]requirement.EpicAggregate, 0, len(epicResp.Epics))
	for idx, epic := range epicResp.Epics {

		epicAggregates = append(epicAggregates, requirement.EpicAggregate{
			Epic:    epic,
			Stories: results[idx],
		})
	}

	aggregate := requirement.NewRequirementAggregate(requirementResp.Requirement, epicAggregates)

	return &GenerateRequirementPackageResponse{
		RequirementAggregate: aggregate,
	}, nil

}
