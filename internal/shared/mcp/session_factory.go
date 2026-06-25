package mcp

import (
	"context"
	"errors"
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
) (Session, error) {
	return nil, errors.New(
		"not implemented",
	)
}
