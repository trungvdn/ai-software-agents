package tools

import (
	"testing"
)

func TestSearchCode(t *testing.T) {

	tool := NewSearchCodeTool("../../")

	results, err := tool.Search("GetUser")
	println(results[2])
	if err != nil {
		t.Fatal(err)
	}

	if len(results) == 0 {
		t.Fatal("expected result")
	}
}
