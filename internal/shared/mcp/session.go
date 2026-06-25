package mcp

import (
	"context"
	"encoding/json"
)

type Session interface {
	ListTools(ctx context.Context) ([]Tool, error)

	CallTool(
		ctx context.Context,
		tool string,
		args any,
	) (json.RawMessage, error)

	Close() error
}
