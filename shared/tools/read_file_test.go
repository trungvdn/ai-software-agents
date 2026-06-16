package tools

import (
	"fmt"
	"testing"
)

func TestRead(t *testing.T) {

	tool := NewReadFileTool()

	content, err := tool.Read("../../testdata/user_service.go")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(content)

	if content == "" {
		t.Fatal("expected file content")
	}
}
