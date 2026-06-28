package auth

type Browser interface {
	Open(url string) error
}
