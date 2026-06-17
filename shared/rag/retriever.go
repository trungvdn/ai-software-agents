package rag

import "context"

type SearchResult struct {
	ID      string
	Content string
	Score   float64
	Source  string
}

type Retriever interface {
	Retrieve(
		ctx context.Context,
		query string,
		topK int,
	) ([]SearchResult, error)
}
