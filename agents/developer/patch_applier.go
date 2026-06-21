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

type DefaultPatchApplier struct {
	rootPath string
}

func NewDefaultPatchApplier(rootPath string) *DefaultPatchApplier {
	return &DefaultPatchApplier{
		rootPath: rootPath,
	}
}

func (p *DefaultPatchApplier) Apply(
	candidates []*patchcandidate.PatchCandidate,
) error {
	// Implement the logic to apply the code patches to the source code files.
	// This could involve reading the files, applying the diffs, and writing back the changes.
	// For now, we'll just log the patches to be applied.
	for _, candidate := range candidates {
		log.Printf("Applying patch to file %s:\n%s", candidate.FilePath, candidate.ProposedSnippet)
		content, err := readFile(p.rootPath, candidate.FilePath)
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

func readFile(rootPath string, filePath string) (string, error) {
	absRoot, _ := filepath.Abs(rootPath)
	absPath := filepath.Join(
		rootPath,
		filePath,
	)
	// Check if the absolute path is within the rootPath to prevent directory traversal
	if !strings.HasPrefix(absPath, absRoot) {
		return "", fmt.Errorf("file path %s is outside of the root path %s", filePath, rootPath)
	}
	content, err := os.ReadFile(absPath)
	if err != nil {
		return "", fmt.Errorf("Error reading file %s: %v", filePath, err)
	}
	return string(content), nil
}
