package mcp

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type SessionFactory struct {
	cfg Config
}

func NewSessionFactory(
	cfg Config,
) *SessionFactory {
	return &SessionFactory{
		cfg: cfg,
	}
}

func (f *SessionFactory) Create(
	ctx context.Context,
) (*mcp.ClientSession, error) {
	return &mcp.ClientSession{}, nil
}
