package developer

import (
	"regexp"
	"strings"

	"github.com/trungvdn/ai-software-agents/shared/tools"
)

type CodeRetriever interface {
	Retrieve(
		bug string,
	) (*CodeContext, error)
}

type DefaultCodeRetriever struct {
	searchSymbolTool *tools.SearchSymbolTool
	readFileTool     *tools.ReadFileTool
}

func NewDefaultCodeRetriever(
	searchSymbolTool *tools.SearchSymbolTool,
	readFileTool *tools.ReadFileTool,
) *DefaultCodeRetriever {
	return &DefaultCodeRetriever{
		searchSymbolTool: searchSymbolTool,
		readFileTool:     readFileTool,
	}
}

func (r *DefaultCodeRetriever) Retrieve(
	bug string,
) (*CodeContext, error) {
	// Step 0: Extract relevant code symbols based on the bug description
	targets := extractTargets(bug)
	if len(targets) == 0 {
		return &CodeContext{Files: []*tools.FileContent{}}, nil
	}

	// Step 1: Use the search symbol tool to find relevant files
	const MaxFiles = 5
	fileMatches := make(map[string]*tools.SymbolMatch) // Use map to deduplicate files

	for _, target := range targets {
		// Stop if we've already found enough files
		if len(fileMatches) >= MaxFiles {
			break
		}

		matches, err := r.searchSymbolTool.Search(target)
		if err != nil {
			continue // Skip if search fails for a target
		}

		// Add matches up to MaxFiles total
		for _, match := range matches {
			if len(fileMatches) >= MaxFiles {
				break
			}
			if _, exists := fileMatches[match.File]; !exists {
				fileMatches[match.File] = match
			}
		}
	}

	// Step 2: Read file contents for all matched files
	fileContents := make([]*tools.FileContent, 0, len(fileMatches))
	for file := range fileMatches {
		content, err := r.readFileTool.Read(file)
		if err != nil {
			continue // Skip files that cannot be read
		}
		fileContents = append(fileContents, content)
	}

	return &CodeContext{
		Files: fileContents,
	}, nil
}

// extractTargets extracts potential symbols/identifiers from the bug description
func extractTargets(bug string) []string {
	targets := []string{}

	// Pattern 1: CamelCase words (type/struct/interface names) - prioritized
	camelCasePattern := regexp.MustCompile(`\b([A-Z][a-z]+(?:[A-Z][a-z]+)*)\b`)
	camelMatches := camelCasePattern.FindAllString(bug, -1)
	targets = append(targets, camelMatches...)

	// Pattern 2: UPPER_CASE words (constants)
	upperCasePattern := regexp.MustCompile(`\b([A-Z_]+)\b`)
	upperMatches := upperCasePattern.FindAllString(bug, -1)
	for _, match := range upperMatches {
		if len(match) > 2 && match != "AND" && match != "OR" && match != "THE" {
			targets = append(targets, match)
		}
	}

	// Remove duplicates and empty strings
	uniqueTargets := make(map[string]bool)
	for _, target := range targets {
		target = strings.TrimSpace(target)
		if len(target) > 0 {
			uniqueTargets[target] = true
		}
	}

	// Convert back to slice and limit to top 10 targets
	result := make([]string, 0, len(uniqueTargets))
	for target := range uniqueTargets {
		result = append(result, target)
	}

	if len(result) > 10 {
		result = result[:10]
	}

	return result
}
