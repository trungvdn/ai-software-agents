package mcp

import (
	"github.com/trungvdn/ai-software-agents/internal/integration/confluence/domain"
)

func mapResponseToPage(
	response CreatePageResponse,
) domain.Page {
	return domain.Page{
		ID:      response.ID,
		URL:     response.URL,
		Version: response.Version,
	}
}
