package developer

import (
	"github.com/trungvdn/ai-software-agents/domain/analysis"
	"github.com/trungvdn/ai-software-agents/domain/codepatch"
	"github.com/trungvdn/ai-software-agents/domain/patchcandidate"
	"github.com/trungvdn/ai-software-agents/internal/knowledge/application/retrieve_knowledge"
)

type Response struct {
	Analysis       *analysis.Analysis
	PatchCandidate []*patchcandidate.PatchCandidate
	Knowledge      *retrieve_knowledge.KnowledgeContextResponse
	CodeContext    *CodeContext
	CodePatches    []*codepatch.CodePatch
}

type SourceFile struct {
	Path    string
	Content string
}
