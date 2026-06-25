package mcp

import (
	"context"
	"encoding/json"
)

type Session interface {
	ListTools(ctx context.Context) ([]ToolInfo, error)

	CallTool(
		ctx context.Context,
		tool Tool,
		args any,
	) (json.RawMessage, error)

	Close() error
}
