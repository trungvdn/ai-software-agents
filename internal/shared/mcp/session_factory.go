package mcp

import (
	"context"
	"fmt"
	"net/http"
	"os/exec"

	mcpSDK "github.com/modelcontextprotocol/go-sdk/mcp"
)

type SessionFactory struct {
	cfg Config
}

func NewSessionFactory(cfg Config) *SessionFactory {
	return &SessionFactory{cfg: cfg}
}

func (f *SessionFactory) Create(ctx context.Context) (Session, error) {
	if err := validateConfig(f.cfg); err != nil {
		return nil, err
	}

	var transport mcpSDK.Transport
	var err error

	switch f.cfg.Transport {
	case TransportRemote:
		if f.cfg.ServerURL == "" {
			return nil, fmt.Errorf("server URL is required for remote transport")
		}
		transport = &mcpSDK.StreamableClientTransport{Endpoint: f.cfg.ServerURL, HTTPClient: &http.Client{Timeout: f.cfg.ConnectTimeout}}
	case TransportStdio:
		if f.cfg.Command == "" {
			return nil, fmt.Errorf("command is required for stdio transport")
		}
		cmd := exec.CommandContext(ctx, f.cfg.Command, f.cfg.Args...)
		for key, value := range f.cfg.Env {
			cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", key, value))
		}
		transport = &mcpSDK.CommandTransport{Command: cmd}
	default:
		return nil, fmt.Errorf("unsupported transport type: %s", f.cfg.Transport)
	}

	client := mcpSDK.NewClient(&mcpSDK.Implementation{Name: "ai-software-agents", Version: "1.0.0"}, nil)
	clientSession, err := client.Connect(ctx, transport, &mcpSDK.ClientSessionOptions{})
	if err != nil {
		return nil, fmt.Errorf("connect to MCP server: %w", err)
	}

	return &sdkSession{session: clientSession}, nil
}

func validateConfig(cfg Config) error {
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
