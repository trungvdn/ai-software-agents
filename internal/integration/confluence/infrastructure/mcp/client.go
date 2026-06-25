package mcp

import (
	"context"
	"errors"
	"time"

	"github.com/trungvdn/ai-software-agents/internal/integration/confluence/domain"
	"github.com/trungvdn/ai-software-agents/internal/shared/mcp"
)

type MCPConfluenceClient struct {
	mcpClient mcp.Client
}

func NewMCPConfluenceClient() *MCPConfluenceClient {
	return &MCPConfluenceClient{}
}

func (c *MCPConfluenceClient) CreatePage(
	ctx context.Context,
	page domain.Page,
) (*domain.Page, error) {
	// Implement the logic to create a page in Confluence using the MCP client
	// You can use the request parameter to get the necessary information for creating the page
	// For example, you might need to call c.client.CreatePage(...) with the appropriate parameters
	req := MapPageToCreateRequest(page)
	resp, err := c.mcpClient.Call(
		ctx,
		mcp.Request{
			Tool:      mcp.ToolConfluenceCreatePage,
			Arguments: req,
		},
		mcp.CallOptions{
			Timeout: 30 * time.Second,
		},
	)
	if err != nil {
		return nil, err
	}
	createResp, ok := resp.Result.(CreatePageResponse)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	result := MapResponseToPage(createResp)
	result.Title = page.Title
	result.Content = page.Content
	result.ParentID = page.ParentID
	result.SpaceKey = page.SpaceKey
	return &result, nil
}

func (c *MCPConfluenceClient) UpdatePage(
	ctx context.Context,
	page domain.Page,
) error {
	return nil
}
