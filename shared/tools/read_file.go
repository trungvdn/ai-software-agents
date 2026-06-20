package tools

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type FileContent struct {
	Path string

	Content string
}

type ReadFileTool struct {
	rootPath string
}

func NewReadFileTool(
	rootPath string,
) *ReadFileTool {
	return &ReadFileTool{
		rootPath: rootPath,
	}
}

func (t *ReadFileTool) Name() string {
	return "read_file"
}

func (t *ReadFileTool) Read(path string) (*FileContent, error) {
	cleanPath := filepath.Clean(path)
	if strings.Contains(
		cleanPath,
		"..",
	) {
		return nil, fmt.Errorf("invalid path: %s", path)
	}
	path = filepath.Join(
		t.rootPath,
		cleanPath,
	)
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return &FileContent{
		Path:    path,
		Content: string(b),
	}, nil
}
