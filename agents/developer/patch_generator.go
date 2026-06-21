package developer

import (
	"fmt"
	"log"
	"strings"

	"github.com/trungvdn/ai-software-agents/domain/codepatch"
	"github.com/trungvdn/ai-software-agents/domain/patchcandidate"
)

type PatchGenerator interface {
	Generate(
		candidate []*patchcandidate.PatchCandidate,
		codeContext *CodeContext,
	) ([]*codepatch.CodePatch, error)
}

type DefaultPatchGenerator struct{}

func NewDefaultPatchGenerator() *DefaultPatchGenerator {
	return &DefaultPatchGenerator{}
}

func (g *DefaultPatchGenerator) Generate(
	candidate []*patchcandidate.PatchCandidate,
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
	for _, c := range candidate {
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
	candidate []*patchcandidate.PatchCandidate,
	fileContent string,
) error {
	// Validate that the candidate's file path matches the provided file content
	for _, c := range candidate {
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
	// For simplicity, we will just return a string that represents the diff.
	// In a real-world scenario, you would use a proper diff library to generate the diff.
	// use lib github.com/sergi/go-diff for a more accurate diff representation.
	return fmt.Sprintf("Diff:\n- %s\n+ %s", originalSnippet, modifiedSnippet)
}
