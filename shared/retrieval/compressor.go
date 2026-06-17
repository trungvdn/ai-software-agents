package retrieval

type Compressor interface {
	Compress(text string) (string, error)
}
