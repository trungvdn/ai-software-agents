package mcp

import "context"

type MCPClient struct {
	adapter Adapter
}

func NewMCPClient(
	adapter Adapter,
) *MCPClient {
	return &MCPClient{
		adapter: adapter,
	}
}

func (c *MCPClient) Call(
	ctx context.Context,
	request Request,
	options CallOptions,
) (*Response, error) {
	return c.adapter.Call(
		ctx,
		request,
		options,
	)
}
