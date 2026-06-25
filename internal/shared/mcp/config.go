package mcp

import (
	"fmt"
	"time"
)

type Config struct {
	Transport TransportType

	// Remote MCP
	ServerURL string

	// stdio MCP
	Command string
	Args    []string
	Env     map[string]string

	ConnectTimeout time.Duration
}

type TransportType string

const (
	TransportRemote TransportType = "remote"
	TransportStdio  TransportType = "stdio"
)

func (cfg Config) validateConfig() error {
	if cfg.Transport == "" {
		return fmt.Errorf("transport is required")
	}

	switch cfg.Transport {
	case TransportRemote:
		if cfg.ServerURL == "" {
			return fmt.Errorf("server URL is required for remote transport")
		}
	case TransportStdio:
		if cfg.Command == "" {
			return fmt.Errorf("command is required for stdio transport")
		}
	default:
		return fmt.Errorf("unsupported transport type: %s", cfg.Transport)
	}

	if cfg.ConnectTimeout < 0 {
		return fmt.Errorf("connect timeout must be non-negative")
	}

	return nil
}
