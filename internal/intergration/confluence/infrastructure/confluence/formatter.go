package confluence

import (
	"github.com/trungvdn/ai-software-agents/internal/intergration/confluence/domain"
	"github.com/trungvdn/ai-software-agents/internal/requirement/domain/requirement"
)

type Formatter interface {
	Format(
		aggregate *requirement.RequirementAggregate,
	) (*domain.Page, error)
}

type MarkdownFormatter struct{}
