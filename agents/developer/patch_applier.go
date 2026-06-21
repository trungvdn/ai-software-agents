package developer

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
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
		content, err := readFile(candidate.FilePath)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %w", candidate.FilePath, err)
		}
		fmt.Printf("Original content of file %s:\n%s", candidate.FilePath, content)
		if !strings.Contains(content, candidate.OriginalSnippet) {
			return fmt.Errorf("original snippet not found in file %s. Cannot apply patch.", candidate.FilePath)
		}
		updated :=
			strings.Replace(
				content,
				candidate.OriginalSnippet,
				candidate.ProposedSnippet,
				1,
			)
		log.Printf("Updated content for file %s:\n%s", candidate.FilePath, updated)
		err = os.WriteFile(candidate.FilePath, []byte(updated), 0644)
		if err != nil {
			return fmt.Errorf("failed to write updated content to file %s: %w", candidate.FilePath, err)
		}
	}
	return nil
}

func readFile(filePath string) (string, error) {
	cleanPath := filepath.Clean(filePath)
	content, err := os.ReadFile(cleanPath)
	if err != nil {
		return "", fmt.Errorf("Error reading file %s: %v", filePath, err)
	}
	return string(content), nil
}
