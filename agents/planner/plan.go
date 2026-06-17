package planner

type Planner interface {
	Plan(
		ctx context.Context,
		bugDescription string,
		analysis string,
	) (*ChangePlan, error)
}
