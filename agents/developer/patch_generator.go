package developer

import (
	"fmt"
	"log"
	"strings"

	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/trungvdn/ai-software-agents/domain/codepatch"
	"github.com/trungvdn/ai-software-agents/domain/patchcandidate"
)

type DiffGenerator interface {
	Generate(
		candidates []*patchcandidate.PatchCandidate,
		codeContext *CodeContext,
	) ([]*codepatch.CodePatch, error)
}

type DefaultDiffGenerator struct{}

func NewDefaultDiffGenerator() *DefaultDiffGenerator {
	return &DefaultDiffGenerator{}
}

func (g *DefaultDiffGenerator) Generate(
	candidates []*patchcandidate.PatchCandidate,
	codeContext *CodeContext,
) ([]*codepatch.CodePatch, error) {
	// Validate the candidate against the code context
	// TODO: Implement validation logic to ensure the candidate's file path and snippets are consistent with the code context.
	// for _, file := range codeContext.Files {
	// 	if err := validateCandidate(candidate, file.Content); err != nil {
	// 		return nil, fmt.Errorf("validation failed for file %s: %w", file.Path, err)
	// 	}
	// }

	// Generate the code patch based on the candidate and code context
	codePatches := []*codepatch.CodePatch{}
	for _, c := range candidates {
		diff := generateDiff(c.OriginalSnippet, c.ProposedSnippet)
		log.Printf("Generated diff for candidate in file %s:\n%s", c.FilePath, diff)
		codePatches = append(codePatches, &codepatch.CodePatch{
			FilePath: c.FilePath,
			Diff:     diff,
		})
	}
	return codePatches, nil
}

func validateCandidate(
	candidates []*patchcandidate.PatchCandidate,
	fileContent string,
) error {
	// Validate that the candidate's file path matches the provided file content
	for _, c := range candidates {
		if !strings.Contains(
			fileContent,
			c.OriginalSnippet,
		) {
			return fmt.Errorf("candidate's original snippet does not match the provided file content")
		}
	}

	return nil
}

func generateDiff(
	originalSnippet string,
	modifiedSnippet string,
) string {
	dmp := diffmatchpatch.New()

	// Compute line-based diffs
	diffs := dmp.DiffMain(originalSnippet, modifiedSnippet, true)

	// Convert to unified diff format
	var result strings.Builder

	// Write unified diff header
	result.WriteString("--- original\n")
	result.WriteString("+++ modified\n")

	// Generate unified diff output
	for _, diff := range diffs {
		switch diff.Type {
		case diffmatchpatch.DiffDelete:
			// Removed lines
			lines := strings.Split(diff.Text, "\n")
			for _, line := range lines {
				if line != "" {
					result.WriteString("- " + line + "\n")
				}
			}
		case diffmatchpatch.DiffInsert:
			// Added lines
			lines := strings.Split(diff.Text, "\n")
			for _, line := range lines {
				if line != "" {
					result.WriteString("+ " + line + "\n")
				}
			}
		case diffmatchpatch.DiffEqual:
			// Context lines
			lines := strings.Split(diff.Text, "\n")
			for _, line := range lines {
				if line != "" {
					result.WriteString("  " + line + "\n")
				}
			}
		}
	}

	return result.String()
}
