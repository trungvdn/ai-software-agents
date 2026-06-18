package retrieval

import (
	"context"
	"log"
	"strings"
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
		for _, doc := range docs {
			exists := false
			for _, existing := range results {
				if strings.Contains(existing.Content, doc.Content) || strings.Contains(doc.Content, existing.Content) {
					exists = true
					break
				}
			}
			if !exists {
				results = append(results, doc)
			}
		}
	}
	return results, nil
}
