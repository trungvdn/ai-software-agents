package patchcandidate

type PatchCandidate struct {
	FilePath        string `json:"file_path"`
	Reason          string `json:"reason"`
	OriginalSnippet string `json:"original_snippet"`
	ProposedSnippet string `json:"proposed_snippet"`
}
