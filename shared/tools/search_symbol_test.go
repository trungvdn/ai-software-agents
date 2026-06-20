package tools

import (
	"path/filepath"
	"testing"
)

func TestSearch_WithTestdata(t *testing.T) {
	// Use actual testdata folder
	testdataPath := filepath.Join("..", "..", "testdata")

	tool := NewSearchSymbolTool(testdataPath)

	// Search for UserService
	matches, err := tool.Search("UserService")
	if err != nil {
		t.Fatalf("Search for UserService failed: %v", err)
	}

	if len(matches) == 0 {
		t.Errorf("expected to find UserService in testdata")
	}

	// Search for UserRepository
	matches, err = tool.Search("UserRepository")
	if err != nil {
		t.Fatalf("Search for UserRepository failed: %v", err)
	}
	if len(matches) == 0 {
		t.Errorf("expected to find UserRepository in testdata")
	}

	// Search for UserController
	matches, err = tool.Search("UserController")
	if err != nil {
		t.Fatalf("Search for UserController failed: %v", err)
	}
	if len(matches) == 0 {
		t.Errorf("expected to find UserController in testdata")
	}
}
