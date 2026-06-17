package retrieval

import "context"

type Retriever interface {
	Retrieve(
		ctx context.Context,
		query string,
		topK int,
	) ([]SearchResult, error)
}
