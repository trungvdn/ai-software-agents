package retrieval

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
		log.Printf(
			"[MergeRetriever] Retrieved %d docs from retriever: %T",
			len(docs),
			retriever,
		)
		results = append(results, docs...)
	}
	return results, nil
}
