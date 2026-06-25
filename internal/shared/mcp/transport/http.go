package transport

import (
	"context"

	"github.com/trungvdn/ai-software-agents/internal/shared/mcp"
)

type HTTPClient struct {
	adapter mcp.Adapter
}

func NewHTTPClient(adapter mcp.Adapter) *HTTPClient {
	return &HTTPClient{
		adapter: adapter,
	}
}

func (c *HTTPClient) Call(
	ctx context.Context,
	request mcp.Request,
	options mcp.CallOptions,
) (*mcp.Response, error) {
	return c.adapter.Call(ctx, request, options)
}
