package tools

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type FileContent struct {
	Path    string
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
	fullPath := filepath.Join(
		t.rootPath,
		cleanPath,
	)
	log.Printf("ReadFileTool: Reading file from fullPath: %s (rootPath=%s, cleanPath=%s)", fullPath, t.rootPath, cleanPath)
	b, err := os.ReadFile(fullPath)
	if err != nil {
		log.Printf("ReadFileTool: Failed to read file: %v", err)
		return nil, err
	}
	log.Printf("ReadFileTool: Successfully read file, size: %d bytes", len(b))

	return &FileContent{
		Path:    cleanPath,
		Content: string(b),
	}, nil
}
