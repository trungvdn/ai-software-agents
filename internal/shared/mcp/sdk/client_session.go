package sdk

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type Session struct {
	session *mcp.ClientSession
}

func (s *Session) CallTool(
	ctx context.Context,
	tool string,
	args any,
) (json.RawMessage, error) {
	return nil, errors.New("not implemented")
}
