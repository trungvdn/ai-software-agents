package planner

import (
	"context"
	"github.com/trungvdn/ai-software-agents/domain/changeplan"
)

type Planner interface {
	Plan(
		ctx context.Context,
		bugDescription string,
		analysis string,
	) (*changeplan.ChangePlan, error)
}
