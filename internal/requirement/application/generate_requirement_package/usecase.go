package generate_requirement_package

import (
	"context"

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
	allStories := make([]story.Story, 0)
	for _, epic := range epicResp.Epics {
		storyResp, err := u.generateStoryUseCase.Generate(ctx, generate_story.GenerateStoryRequest{Epic: epic})
		if err != nil {
			return nil, err
		}
		allStories = append(allStories, storyResp.Stories...)
	}

	aggregate := requirement.NewRequirementAggregate(requirementResp.Requirement, epicResp.Epics, allStories)

	return &GenerateRequirementPackageResponse{
		RequirementAggregate: aggregate,
	}, nil

}
