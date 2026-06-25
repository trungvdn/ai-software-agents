package infrastructure

import (
	"github.com/trungvdn/ai-software-agents/internal/integration/confluence/domain"
	"github.com/trungvdn/ai-software-agents/internal/requirement/domain/requirement"
)

type Formatter interface {
	Format(
		aggregate *requirement.RequirementAggregate,
	) (*domain.Page, error)
}

type ConfluenceFormatter struct {
}

func NewConfluenceFormatter() *ConfluenceFormatter {
	return &ConfluenceFormatter{}
}

func (f *ConfluenceFormatter) Format(
	aggregate *requirement.RequirementAggregate,
) (*domain.Page, error) {
	// Implement the logic to format the requirement aggregate into a Confluence page
	// You can use the aggregate parameter to access the requirement, epics, and stories
	// For example, you might create a new domain.Page with the appropriate title and content based on the aggregate
	return &domain.Page{
		Title:   "",
		Content: "Formatted content based on the requirement aggregate",
	}, nil
}
