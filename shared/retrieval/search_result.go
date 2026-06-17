package retrieval

type SearchResult struct {
	ID      string
	Content string
	Score   float64
	Source  string

	Metadata SearchMetadata
}

type SearchMetadata struct {
	ImportanceScore float64
	UsageCount      int
}
