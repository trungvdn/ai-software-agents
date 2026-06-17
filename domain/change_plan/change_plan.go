package change_plan

type ChangePlanner interface {
	Plan(
		ctx context.Context,
		bugDescription string,
		analysis string,
	) (*ChangePlan, error)
}

type ChangePlan struct {
	AffectedFiles []string
	Steps         []string
}
