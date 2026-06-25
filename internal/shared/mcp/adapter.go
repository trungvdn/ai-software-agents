package mcp

import "context"

type Adapter interface {
	Call(
		ctx context.Context,
		request Request,
		options CallOptions,
	) (*Response, error)
}
