package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/trungvdn/ai-software-agents/internal/shared/mcp"
)

func TestHTTPAdapterCall(t *testing.T) {
	var receivedRequest mcp.Request

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		if r.Method != http.MethodPost {
			t.Fatalf("expected POST request, got %s", r.Method)
		}
		if err := json.NewDecoder(r.Body).Decode(&receivedRequest); err != nil {
			t.Fatalf("decode request: %v", err)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{"result": "adapter-ok"})
	}))
	defer server.Close()

	adapter := NewHTTPAdapter(server.Client(), server.URL)
	resp, err := adapter.Call(context.Background(), mcp.Request{Tool: "echo", Arguments: map[string]string{"message": "hi"}}, mcp.CallOptions{Timeout: 5 * time.Second})
	if err != nil {
		t.Fatalf("Call returned error: %v", err)
	}
	if resp == nil {
		t.Fatal("expected response, got nil")
	}
	if receivedRequest.Tool != "echo" {
		t.Fatalf("expected tool echo, got %s", receivedRequest.Tool)
	}
	if resp.Result != "adapter-ok" {
		t.Fatalf("expected result adapter-ok, got %#v", resp.Result)
	}
}
