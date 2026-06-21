package developer

import (
	"github.com/trungvdn/ai-software-agents/domain/analysis"
	"github.com/trungvdn/ai-software-agents/domain/codepatch"
	"github.com/trungvdn/ai-software-agents/domain/patchcandidate"
)

type Response struct {
	Analysis    *analysis.Analysis
	PatchPlan   *patchcandidate.PatchCandidate
	Knowledge   *KnowledgeContext
	CodeContext *CodeContext
	CodePatches []codepatch.CodePatch
}

type SourceFile struct {
	Path string

	Content string
}
