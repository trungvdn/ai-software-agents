package coder

import (
	"context"
	"github.com/trungvdn/ai-software-agents/domain/changeplan"
	"github.com/trungvdn/ai-software-agents/domain/codepatch"
)

type Coder interface {
	GeneratePatches(
		ctx context.Context,
		bugDescription string,
		plan *changeplan.ChangePlan,
	) ([]codepatch.CodePatch, error)
}
