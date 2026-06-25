package mcp

import "errors"

var (
	ErrToolNotFound           = errors.New("tool not found")
	ErrCallTimeout            = errors.New("call timeout")
	ErrUnexpectedResponseType = errors.New("unexpected response type")
)
