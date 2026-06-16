package embedding

import (
	"context"
	"hash/fnv"
)

// MockEmbedder generates deterministic embeddings from text (for testing)
// No API calls, completely free
type MockEmbedder struct {
	dimension int
}

// NewMockEmbedder creates a new mock embedder for testing
// Generates consistent embeddings based on text hash
func NewMockEmbedder(dimension int) *MockEmbedder {
	if dimension == 0 {
		dimension = 384 // Default dimension for testing
	}
	return &MockEmbedder{
		dimension: dimension,
	}
}

// Embed generates a deterministic embedding based on text hash
func (m *MockEmbedder) Embed(
	ctx context.Context,
	text string,
) ([]float32, error) {
	hash := fnv.New32a()
	hash.Write([]byte(text))
	seed := hash.Sum32()

	embedding := make([]float32, m.dimension)

	// Generate pseudo-random floats using seed
	rng := uint64(seed)
	for i := 0; i < m.dimension; i++ {
		rng = rng*1103515245 + 12345
		val := float32((rng>>16)%32768) / 32768.0
		embedding[i] = val - 0.5
	}

	return embedding, nil
}
