package mcp

import (
	"context"
	"encoding/json"
	"fmt"

	mcpSDK "github.com/modelcontextprotocol/go-sdk/mcp"
)

type sdkSession struct {
	session *mcpSDK.ClientSession
}

func NewSdkSession() *sdkSession {
	return &sdkSession{}
}

func (s *sdkSession) ListTools(ctx context.Context) ([]ToolInfo, error) {
	result, err := s.session.ListTools(ctx, &mcpSDK.ListToolsParams{})
	if err != nil {
		return nil, fmt.Errorf("list tools: %w", err)
	}

	tools := make([]ToolInfo, 0, len(result.Tools))
	for _, tool := range result.Tools {
		if tool == nil {
			continue
		}
		tools = append(tools, ToolInfo{Name: Tool(tool.Name), Description: tool.Description})
	}
	return tools, nil
}

func (s *sdkSession) CallTool(ctx context.Context, tool Tool, args any) (json.RawMessage, error) {
	params := &mcpSDK.CallToolParams{
		Name:      string(tool),
		Arguments: args,
	}

	result, err := s.session.CallTool(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("call tool %s: %w", tool, err)
	}

	payload, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("marshal tool result: %w", err)
	}

	return payload, nil
}

func (s *sdkSession) Close() error {
	return s.session.Close()
}
