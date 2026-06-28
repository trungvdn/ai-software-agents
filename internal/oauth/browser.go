package oauth

type Browser interface {
	Open(url string) error
}
