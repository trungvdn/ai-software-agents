package mcp

import (
	"github.com/trungvdn/ai-software-agents/internal/integration/confluence/domain"
)

func MapPageToCreateRequest(
	page domain.Page,
) CreatePageRequest {
	return CreatePageRequest{
		ParentID: page.ParentID,
		Title:    page.Title,
		Content:  page.Content,
		SpaceKey: page.SpaceKey,
	}
}
