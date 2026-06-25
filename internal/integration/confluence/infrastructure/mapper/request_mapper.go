package mapper

import (
	"github.com/trungvdn/ai-software-agents/internal/integration/confluence/domain"
	"github.com/trungvdn/ai-software-agents/internal/integration/confluence/infrastructure/mcp"
)

func MapPageToCreateRequest(
	page domain.Page,
) mcp.CreatePageRequest {
	return mcp.CreatePageRequest{
		ParentID: page.ParentID,
		Title:    page.Title,
		Content:  page.Content,
		SpaceKey: page.SpaceKey,
	}
}
