package mcp

import "context"

type Client interface {
	Call(
		ctx context.Context,
		request Request,
	) (*Response, error)
}

type MCPClient struct {
}
