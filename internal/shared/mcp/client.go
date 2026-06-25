package mcp

import (
	"context"
	"encoding/json"
)

type Client interface {
	Call(
		ctx context.Context,
		tool string,
		args any,
	) (json.RawMessage, error)
}
