package mcp

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"time"

	mcpSDK "github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/trungvdn/ai-software-agents/internal/oauth"
)

type SessionFactory struct {
	cfg Config
}

func NewSessionFactory(cfg Config) *SessionFactory {
	return &SessionFactory{cfg: cfg}
}

func (f *SessionFactory) Create(ctx context.Context) (Session, error) {
	if err := f.cfg.Validate(); err != nil {
		return nil, err
	}

	var transport mcpSDK.Transport
	var err error

	switch f.cfg.Transport {
	case TransportRemote:
		if f.cfg.Remote.ServerURL == "" {
			return nil, fmt.Errorf("server URL is required for remote transport")
		}
		if f.cfg.ConnectTimeout <= 0 {
			f.cfg.ConnectTimeout = 30 * time.Second
		}
		httpClient := &http.Client{Timeout: f.cfg.ConnectTimeout}
		handler, err := oauth.NewAuthorizationCodeHandler(oauth.OAuthConfig{
			ClientID:    "",
			RedirectURI: "http://localhost:8080/callback",
		})
		if err != nil {
			return nil, err
		}
		transport = &mcpSDK.StreamableClientTransport{
			Endpoint:     f.cfg.Remote.ServerURL,
			HTTPClient:   httpClient,
			OAuthHandler: handler,
		}
	case TransportStdio:
		if f.cfg.Stdio.Command == "" {
			return nil, fmt.Errorf("command is required for stdio transport")
		}
		cmd := exec.CommandContext(ctx, f.cfg.Stdio.Command, f.cfg.Stdio.Args...)
		cmd.Env = os.Environ()
		for key, value := range f.cfg.Stdio.Env {
			cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", key, value))
		}
		transport = &mcpSDK.CommandTransport{Command: cmd}
	default:
		return nil, fmt.Errorf("unsupported transport type: %s", f.cfg.Transport)
	}
	client := mcpSDK.NewClient(
		&mcpSDK.Implementation{
			Name:    f.cfg.Client.Name,
			Version: f.cfg.Client.Version,
		}, &mcpSDK.ClientOptions{})
	clientSession, err := client.Connect(ctx, transport, &mcpSDK.ClientSessionOptions{})
	if err != nil {
		return nil, fmt.Errorf("connect to MCP server: %w", err)
	}

	return &SDKSession{session: clientSession}, nil
}
