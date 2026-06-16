package tools

import (
	"os"
	"path/filepath"
	"strings"
)

type SearchCodeTool struct {
	root string
}

func NewSearchCodeTool(root string) *SearchCodeTool {
	return &SearchCodeTool{root: root}
}

func (t *SearchCodeTool) Search(keyword string) ([]string, error) {

	var matches []string

	err := filepath.Walk(
		t.root,
		func(path string, info os.FileInfo, err error) error {

			if err != nil {
				return nil
			}

			if info.IsDir() {
				return nil
			}

			if !strings.HasSuffix(path, ".go") {
				return nil
			}

			b, err := os.ReadFile(path)
			if err != nil {
				return nil
			}

			if strings.Contains(
				strings.ToLower(string(b)),
				strings.ToLower(keyword),
			) {
				matches = append(matches, path)
			}

			return nil
		},
	)

	return matches, err
}
