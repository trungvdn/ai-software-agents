package retrieval

import (
	"context"
	"log"
)

type MergeRetriever struct {
	retrievers []Retriever
}

func NewMergeRetriever(
	retrievers ...Retriever,
) *MergeRetriever {
	return &MergeRetriever{
		retrievers: retrievers,
	}
}

func (m *MergeRetriever) Retrieve(
	ctx context.Context,
	query string,
	topK int,
) ([]SearchResult, error) {

	var results []SearchResult
	for _, retriever := range m.retrievers {
		docs, err := retriever.Retrieve(
			ctx,
			query,
			topK,
		)

		if err != nil {
			log.Printf(
				"[MergeRetriever] Error retrieving from retriever: %T, error: %v",
				retriever,
				err,
			)
			continue
		}

		log.Printf(
			"[MergeRetriever] Retrieved %d docs from retriever: %T",
			len(docs),
			retriever,
		)

		// Deduplicate
		seen := map[string]struct{}{}
		for _, doc := range docs {
			if _, exists := seen[doc.Content]; !exists {
				seen[doc.Content] = struct{}{}
				results = append(results, doc)
			}
		}
	}
	return results, nil
}
