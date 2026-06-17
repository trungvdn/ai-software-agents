package bugfix

type Response struct {
	Analysis     string
	RootCause    string
	SuggestedFix string

	Plan *ChangePlan
}
