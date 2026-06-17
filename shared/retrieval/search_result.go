package retrieval

type SearchResult struct {
	ID      string
	Content string
	Score   float64
	Source  string

	Metadata map[string]string
}
