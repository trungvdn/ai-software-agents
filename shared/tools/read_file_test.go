package tools

import (
	"testing"
)

func TestRead(t *testing.T) {

	tool := NewReadFileTool(
		"../../testdata",
	)

	content, err := tool.Read("user_service.go")
	if err != nil {
		t.Fatal(err)
	}

	if len(content.Content) == 0 {
		t.Errorf("expected file content, got empty string")
	}
}
