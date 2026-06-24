package generate_story

import "github.com/trungvdn/ai-software-agents/internal/requirement/domain/epic"

type GenerateStoryRequest struct {
	Epics epic.Epic
}
