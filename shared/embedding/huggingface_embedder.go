package embedding

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type HuggingFaceEmbedder struct {
	client *http.Client
	apiKey string
	model  string
	apiURL string
}

type huggingFaceRequest struct {
	Inputs string `json:"inputs"`
}

// NewHuggingFaceEmbedder creates a new Hugging Face embedder (free tier available)
// API key can be obtained from https://huggingface.co/settings/tokens
// Popular free models: sentence-transformers/all-MiniLM-L6-v2, sentence-transformers/all-mpnet-base-v2
func NewHuggingFaceEmbedder(apiKey, model string) *HuggingFaceEmbedder {
	if model == "" {
		model = "sentence-transformers/all-MiniLM-L6-v2"
	}

	return &HuggingFaceEmbedder{
		client: &http.Client{},
		apiKey: apiKey,
		model:  model,
		apiURL: "https://api-inference.huggingface.co/pipeline/feature-extraction",
	}
}

// Embed generates embeddings using Hugging Face Inference API (free tier: 30K calls/month)
func (e *HuggingFaceEmbedder) Embed(
	ctx context.Context,
	text string,
) ([]float32, error) {
	reqBody := huggingFaceRequest{
		Inputs: text,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s/%s", e.apiURL, e.model), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", e.apiKey))
	req.Header.Set("Content-Type", "application/json")

	resp, err := e.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("hugging face error: status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var result []float32
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}
