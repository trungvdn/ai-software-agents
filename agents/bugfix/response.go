package bugfix

type Response struct {
	Analysis string
}

func (r *Response) String(rootCause string, SuggestedFix string) string {
	return "Root Cause: " + rootCause + "\nSuggested Fix: " + SuggestedFix
}
