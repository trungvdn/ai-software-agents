package developer

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/trungvdn/ai-software-agents/domain/patchcandidate"
)

type PatchApplier interface {
	Apply(
		candidates []*patchcandidate.PatchCandidate,
	) error
}

type DefaultPatchApplier struct{}

func NewDefaultPatchApplier() *DefaultPatchApplier {
	return &DefaultPatchApplier{}
}

func (p *DefaultPatchApplier) Apply(
	candidates []*patchcandidate.PatchCandidate,
) error {
	// Implement the logic to apply the code patches to the source code files.
	// This could involve reading the files, applying the diffs, and writing back the changes.
	// For now, we'll just log the patches to be applied.
	for _, candidate := range candidates {
		log.Printf("Applying patch to file %s:\n%s", candidate.FilePath, candidate.ProposedSnippet)
		content := readFile(candidate.FilePath)
		fmt.Printf("Original content of file %s:\n%s", candidate.FilePath, content)

		updated :=
			strings.Replace(
				content,
				candidate.OriginalSnippet,
				candidate.ProposedSnippet,
				1,
			)
		log.Printf("Updated content for file %s:\n%s", candidate.FilePath, updated)
		os.WriteFile(candidate.FilePath, []byte(updated), 0644)
	}
	return nil
}

func readFile(filePath string) string {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file %s: %v", filePath, err)
	}
	return string(content)
}
