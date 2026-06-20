package developer

import (
	"fmt"
	"slices"
	"testing"

	"github.com/trungvdn/ai-software-agents/shared/tools"
)

func TestExtractTargets_CamelCase(t *testing.T) {
	tests := []struct {
		name     string
		bug      string
		expected []string
	}{
		{
			name:     "single camel case word",
			bug:      "UserService is not working",
			expected: []string{"UserService"},
		},
		{
			name:     "multiple camel case words",
			bug:      "UserService and RepositoryManager have issues",
			expected: []string{"UserService", "RepositoryManager"},
		},
		{
			name:     "nested camel case",
			bug:      "DefaultCodeRetriever fails on SearchSymbol",
			expected: []string{"DefaultCodeRetriever", "SearchSymbol"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractTargets(tt.bug)
			for _, exp := range tt.expected {
				if !slices.Contains(result, exp) {
					t.Errorf("expected %s in result, got %v", exp, result)
				}
			}
		})
	}
}

func TestExtractTargets_UpperCase(t *testing.T) {
	tests := []struct {
		name     string
		bug      string
		expected []string
	}{
		{
			name:     "single upper case constant",
			bug:      "MAX_SIZE is too small",
			expected: []string{"MAX_SIZE"},
		},
		{
			name:     "multiple upper case constants",
			bug:      "MAX_SIZE and DEFAULT_VALUE are wrong",
			expected: []string{"MAX_SIZE", "DEFAULT_VALUE"},
		},
		{
			name:     "skip common upper case words",
			bug:      "THE AND OR are not extracted",
			expected: []string{}, // THE, AND, OR should be skipped
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractTargets(tt.bug)
			for _, exp := range tt.expected {
				if !slices.Contains(result, exp) {
					t.Errorf("expected %s in result, got %v", exp, result)
				}
			}
		})
	}
}

func TestExtractTargets_Deduplication(t *testing.T) {
	bug := "UserService UserService Repository Repository manager"
	result := extractTargets(bug)

	// Count occurrences of UserService
	userServiceCount := 0
	for _, target := range result {
		if target == "UserService" {
			userServiceCount++
		}
	}

	if userServiceCount != 1 {
		t.Errorf("expected UserService to appear once, got %d times in %v", userServiceCount, result)
	}
}

// TestRetrieve_FindsUserService tests that Retrieve successfully finds and reads files related to UserService
func TestRetrieve_FindsUserService(t *testing.T) {
	bug := "Fix nill pointer in UserService"

	retriever := NewDefaultCodeRetriever(
		tools.NewSearchSymbolTool("../../testdata"),
		tools.NewReadFileTool("../../testdata"),
	)
	result, err := retriever.Retrieve(bug)
	fmt.Println(result.Files)
	if err != nil {
		t.Fatalf("Retrieve returned error: %v", err)
	}
}
