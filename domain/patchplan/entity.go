package patchplan

type PatchPlan struct {
	FilePath string `json:"file_path"`

	Reason string `json:"reason"`

	Changes []string `json:"changes"`
}
