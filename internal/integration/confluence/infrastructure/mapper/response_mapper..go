package mapper

import (
	"github.com/trungvdn/ai-software-agents/internal/integration/confluence/domain"
	"github.com/trungvdn/ai-software-agents/internal/integration/confluence/infrastructure/mcp"
)

func MapResponseToPage(
	page domain.Page,
) mcp.CreatePageResponse {
	return mcp.CreatePageResponse{
		ID:      page.ID,
		Version: page.Version,
		URL:     page.URL,
	}
}
