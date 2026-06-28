package oauth

type OAuthConfig struct {
	ClientID    string
	RedirectURI string

	Scopes    []string
	ServerURL string
}
