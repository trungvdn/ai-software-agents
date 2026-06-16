package embedding

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

type Embedder interface {
	Embed(
		ctx context.Context,
		text string,
	) ([]float32, error)
}

type OpenAIEmbedder struct {
	client *openai.Client
	model  openai.EmbeddingModel
}

// NewOpenAIEmbedder creates a new OpenAI embedder with text-embedding-3-small model
func NewOpenAIEmbedder(apiKey string) *OpenAIEmbedder {
	return &OpenAIEmbedder{
		client: openai.NewClient(apiKey),
		model:  openai.SmallEmbedding3,
	}
}

// Embed generates embeddings for the given text using OpenAI's API
func (e *OpenAIEmbedder) Embed(
	ctx context.Context,
	text string,
) ([]float32, error) {
	resp, err := e.client.CreateEmbeddings(ctx, openai.EmbeddingRequest{
		Input: []string{text},
		Model: e.model,
	})
	if err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, nil
	}

	return resp.Data[0].Embedding, nil
}
