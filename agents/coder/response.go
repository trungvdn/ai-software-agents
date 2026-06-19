package coder

import (
	"github.com/trungvdn/ai-software-agents/domain/codepatch"
)

type Response struct {
	Patches []codepatch.CodePatch
}

type PatchResponse struct {
	Patches []codepatch.CodePatch `json:"patch"`
}
