package retrieval

import "context"

type Retriever interface {
	Retrieve(
		ctx context.Context,
		query string,
		topK int,
	) ([]SearchResult, error)
}

type ReflectionRetriever interface {
	Retrieve(
		ctx context.Context,
		query string,
		topK int,
	) ([]SearchResult, error)
}

type HistoricalBugRetriever interface {
	Retrieve(
		ctx context.Context,
		query string,
		topK int,
	) ([]SearchResult, error)
}
