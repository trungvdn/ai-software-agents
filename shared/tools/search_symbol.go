package tools

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type SymbolMatch struct {
	File string `json:"file"`
	Line int    `json:"line"`
}

type SearchSymbolTool struct {
	rootPath string
}

func NewSearchSymbolTool(
	rootPath string,
) *SearchSymbolTool {
	return &SearchSymbolTool{
		rootPath: rootPath,
	}
}

func (t *SearchSymbolTool) Search(
	symbol string,
) ([]SymbolMatch, error) {
	var matches []SymbolMatch

	ignoredDirs := map[string]bool{
		".git":         true,
		"vendor":       true,
		"node_modules": true,
		".idea":        true,
		".vscode":      true,
	}

	ignoredExts := map[string]bool{
		".png": true,
		".jpg": true,
		".pdf": true,
		".exe": true,
	}

	err := filepath.Walk(t.rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // continue walking
		}

		// Skip ignored directories
		if info.IsDir() {
			if ignoredDirs[info.Name()] {
				return filepath.SkipDir
			}
			return nil
		}

		// Skip ignored file extensions
		ext := filepath.Ext(path)
		if ignoredExts[ext] {
			return nil
		}

		// Only search in text files (source code files)
		if !isTextFile(ext) {
			return nil
		}

		// Search for symbol in file
		file, err := os.Open(path)
		if err != nil {
			return nil
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		lineNum := 1
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, symbol) {
				// Get relative path
				relPath, err := filepath.Rel(t.rootPath, path)
				if err != nil {
					relPath = path
				}
				matches = append(matches, SymbolMatch{
					File: relPath,
					Line: lineNum,
				})
			}
			lineNum++
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error searching for symbol: %w", err)
	}
	return matches, nil
}

func isTextFile(ext string) bool {
	textExts := map[string]bool{
		".go":     true,
		".java":   true,
		".py":     true,
		".js":     true,
		".ts":     true,
		".tsx":    true,
		".jsx":    true,
		".cpp":    true,
		".c":      true,
		".h":      true,
		".hpp":    true,
		".rs":     true,
		".rb":     true,
		".php":    true,
		".sql":    true,
		".json":   true,
		".yaml":   true,
		".yml":    true,
		".xml":    true,
		".html":   true,
		".css":    true,
		".md":     true,
		".txt":    true,
		".sh":     true,
		".bat":    true,
		".gradle": true,
		".mod":    true,
		".sum":    true,
	}
	return textExts[ext]
}
