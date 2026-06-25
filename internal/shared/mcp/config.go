package mcp

import "time"

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
