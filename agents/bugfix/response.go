package bugfix

import (
	"github.com/trungvdn/ai-software-agents/domain/changeplan"
	"github.com/trungvdn/ai-software-agents/domain/codepatch"
)

type Response struct {
	Analysis     string
	RootCause    string
	SuggestedFix string

	Plan        *changeplan.ChangePlan
	CodePatches []codepatch.CodePatch
}
