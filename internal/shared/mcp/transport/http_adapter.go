package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/trungvdn/ai-software-agents/internal/shared/mcp"
)

type HTTPAdapter struct {
	httpClient *http.Client
	endpoint   string
}

func NewHTTPAdapter(httpClient *http.Client,
	endpoint string) *HTTPAdapter {
	return &HTTPAdapter{
		httpClient: httpClient,
		endpoint:   endpoint,
	}
}

func (h *HTTPAdapter) Call(
	ctx context.Context,
	request mcp.Request,
	options mcp.CallOptions,
) (*mcp.Response, error) {
	payload, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	if options.Timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, options.Timeout)
		defer cancel()
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, h.endpoint, bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := h.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result struct {
		Result any `json:"result"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return &mcp.Response{Result: result.Result}, nil
}
