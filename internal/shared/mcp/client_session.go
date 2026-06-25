package mcp

import (
	"context"
	"encoding/json"
	"fmt"

	mcpSDK "github.com/modelcontextprotocol/go-sdk/mcp"
)

type SDKSession struct {
	session *mcpSDK.ClientSession
}

func (s *SDKSession) ListTools(ctx context.Context) ([]ToolInfo, error) {
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

func (s *SDKSession) CallTool(ctx context.Context, tool Tool, args any) (json.RawMessage, error) {
	params := &mcpSDK.CallToolParams{
		Name:      string(tool),
		Arguments: args,
	}

	result, err := s.session.CallTool(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("call tool %s: %w", tool, err)
	}

	payload, err := json.Marshal(result.Content)
	if err != nil {
		return nil, fmt.Errorf("marshal tool result: %w", err)
	}

	return payload, nil
}

func (s *SDKSession) Close() error {
	return s.session.Close()
}
