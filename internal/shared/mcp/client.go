package mcp

import "context"

type Client interface {
	Call(
		ctx context.Context,
		request Request,
		options CallOptions,
	) (*Response, error)
}
