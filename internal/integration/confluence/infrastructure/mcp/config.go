package mcp

import "time"

type Config struct {
	Transport string

	Command string
	Args    []string
	Env     map[string]string

	Endpoint string

	Timeout time.Duration
}
