package llm

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/ollama/ollama/api"
)

type OllamaConfig struct {
	Endpoint string
	Model    string
}

type OllamaClient struct {
	Config OllamaConfig
	Client *api.Client
}

type ChatResponse struct {
	Content   string
	Reasoning string
}

type AnalysisResponse struct {
	Analysis     string
	RootCause    string
	SuggestedFix string
}

func NewOllamaClient(config OllamaConfig) (*OllamaClient, error) {
	// Parse endpoint URL
	u, err := url.Parse(config.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("invalid endpoint URL: %w", err)
	}

	// Create HTTP client
	httpClient := &http.Client{}

	// Create Ollama API client
	client := api.NewClient(u, httpClient)
	if client == nil {
		return nil, fmt.Errorf("failed to create ollama client")
	}

	return &OllamaClient{Config: config, Client: client}, nil
}

// Chat sends a prompt to Ollama and returns the response with reasoning
func (c *OllamaClient) Chat(ctx context.Context, prompt string) (string, error) {
	if c.Client == nil {
		return "", fmt.Errorf("ollama client not initialized")
	}

	request := &api.GenerateRequest{
		Model:  c.Config.Model,
		Prompt: prompt,
	}

	var fullResponse string

	err := c.Client.Generate(ctx, request, func(resp api.GenerateResponse) error {
		fullResponse += resp.Response
		return nil
	})

	if err != nil {
		return "", fmt.Errorf("failed to generate response: %w", err)
	}

	return fullResponse, nil
}

// ChatWithHistory sends messages with conversation history to Ollama
func (c *OllamaClient) ChatWithHistory(ctx context.Context, messages []map[string]string) (string, error) {
	if c.Client == nil {
		return "", fmt.Errorf("ollama client not initialized")
	}

	// Build prompt from messages
	prompt := ""
	for _, msg := range messages {
		if role, ok := msg["role"]; ok {
			if content, ok := msg["content"]; ok {
				prompt += fmt.Sprintf("%s: %s\n", role, content)
			}
		}
	}

	return c.Chat(ctx, prompt)
}

// Analyze sends an analysis prompt and returns structured analysis response
func (c *OllamaClient) Analyze(ctx context.Context, prompt string) (*AnalysisResponse, error) {
	if c.Client == nil {
		return nil, fmt.Errorf("ollama client not initialized")
	}

	// Create analysis prompt
	analysisPrompt := fmt.Sprintf(`%s

Please provide your analysis in the following format:

=== ANALYSIS ===
<detailed analysis>

Potential root cause:
<root cause description>

Suggested fix:
<suggested fix or solution>`, prompt)

	response, err := c.Chat(ctx, analysisPrompt)
	if err != nil {
		return nil, err
	}

	// Parse response into structured format
	analysis := parseAnalysisResponse(response)
	return analysis, nil
}

// parseAnalysisResponse extracts sections from the response
func parseAnalysisResponse(response string) *AnalysisResponse {
	analysisResp := &AnalysisResponse{
		Analysis:     response,
		RootCause:    "",
		SuggestedFix: "",
	}

	// Simple parsing - look for keywords
	lines := fmt.Sprintf("%s\n", response)

	// Extract root cause section
	if idx := findSection(response, "root cause"); idx >= 0 {
		analysisResp.RootCause = extractSection(response, idx, "Suggested fix")
	}

	// Extract suggested fix section
	if idx := findSection(response, "suggested fix"); idx >= 0 {
		analysisResp.SuggestedFix = extractSection(response, idx, "")
	}

	return analysisResp
}

// findSection finds the index of a section keyword
func findSection(text, keyword string) int {
	for i := 0; i < len(text)-len(keyword); i++ {
		if text[i:i+len(keyword)] == keyword {
			return i
		}
	}
	return -1
}

// extractSection extracts text between two keywords
func extractSection(text string, startIdx int, endKeyword string) string {
	if startIdx < 0 {
		return ""
	}

	startIdx += len("suggested fix") + 1

	endIdx := len(text)
	if endKeyword != "" {
		if idx := findSection(text[startIdx:], endKeyword); idx >= 0 {
			endIdx = startIdx + idx
		}
	}

	return text[startIdx:endIdx]
}
