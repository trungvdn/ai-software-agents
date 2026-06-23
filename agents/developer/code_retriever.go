package developer

import (
	"log"
	"regexp"
	"strings"

	"github.com/trungvdn/ai-software-agents/shared/tools"
)

type CodeRetriever interface {
	Retrieve(
		query *RetrievalQuery,
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

type RetrievalQuery struct {
	Query            string
	CandidateSymbols []string
}

func (r *DefaultCodeRetriever) Retrieve(
	query *RetrievalQuery,
) (*CodeContext, error) {
	// Step 0: Extract relevant code symbols based on the bug description

	targets := query.CandidateSymbols
	if len(targets) == 0 {
		targets = extractTargets(query.Query)
	}
	log.Printf(
		"CodeRetriever: query=%s targets=%v",
		query.Query,
		targets,
	)

	if len(targets) == 0 {
		return &CodeContext{
			Files: []*tools.FileContent{},
		}, nil
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
			log.Printf("CodeRetriever: Search for target '%s' failed: %v", target, err)
			continue // Skip if search fails for a target
		}
		log.Printf("CodeRetriever: Found %d matches for target '%s'", len(matches), target)

		// Add matches up to MaxFiles total
		for _, match := range matches {
			if len(fileMatches) >= MaxFiles {
				break
			}
			if _, exists := fileMatches[match.File]; !exists {
				fileMatches[match.File] = match
				log.Printf("CodeRetriever: Added file match: %s (line %d)", match.File, match.Line)
			}
		}
	}

	log.Printf("CodeRetriever: Total unique files found: %d", len(fileMatches))

	// Step 2: Read file contents for all matched files
	fileContents := make([]*tools.FileContent, 0, len(fileMatches))
	for file := range fileMatches {
		content, err := r.readFileTool.Read(file)
		if err != nil {
			log.Printf("CodeRetriever: Failed to read file '%s': %v", file, err)
			continue // Skip files that cannot be read
		}
		log.Printf("CodeRetriever: Successfully read file: %s", file)
		fileContents = append(fileContents, content)
	}

	log.Printf("CodeRetriever: Returning %d file contents", len(fileContents))
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
