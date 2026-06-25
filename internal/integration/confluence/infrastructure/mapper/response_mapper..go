package mapper

import (
	"github.com/trungvdn/ai-software-agents/internal/integration/confluence/domain"
	"github.com/trungvdn/ai-software-agents/internal/integration/confluence/infrastructure/mcp"
)

func MapResponseToPage(
	response mcp.CreatePageResponse,
) domain.Page {
	return domain.Page{
		ID:      response.ID,
		URL:     response.URL,
		Version: response.Version,
	}
}
