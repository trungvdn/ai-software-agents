package developer

import (
	"github.com/trungvdn/ai-software-agents/domain/analysis"
	"github.com/trungvdn/ai-software-agents/domain/codepatch"
)

type Response struct {
	Analysis    *analysis.Analysis
	Knowledge   *KnowledgeContext
	CodeContext *CodeContext
	CodePatches []codepatch.CodePatch
}

type SourceFile struct {
	Path string

	Content string
}
