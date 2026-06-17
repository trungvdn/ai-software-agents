package bugfix

import (
	"github.com/trungvdn/ai-software-agents/domain/changeplan"
)

type Response struct {
	Analysis     string
	RootCause    string
	SuggestedFix string

	Plan *changeplan.ChangePlan
}
