package developer

import (
	"github.com/trungvdn/ai-software-agents/domain/analysis"
	"github.com/trungvdn/ai-software-agents/domain/codepatch"
)

type Response struct {
	Analysis *analysis.Analysis

	Knowledge *KnowledgeContext

	CodeContext *CodeContext

	CodePatches []codepatch.CodePatch
}

type KnowledgeContext struct {
	Reflections []string

	HistoricalBugs []string
}

type CodeContext struct {
	Files []SourceFile
}

type SourceFile struct {
	Path string

	Content string
}
