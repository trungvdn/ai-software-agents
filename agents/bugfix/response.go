package bugfix

type Response struct {
	Analysis     string
	RootCause    string
	SuggestedFix string
}

func (r *Response) String(rootCause string, SuggestedFix string) string {
	return "Root Cause: " + rootCause + "\nSuggested Fix: " + SuggestedFix
}
