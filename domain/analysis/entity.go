package analysis

type Analysis struct {
	RootCause    string  `json:"root_cause"`
	SuggestedFix string  `json:"fix_strategy"`
	Confidence   float64 `json:"confidence"`
}
