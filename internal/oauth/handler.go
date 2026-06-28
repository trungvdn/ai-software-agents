package oauth

import (
	sdkauth "github.com/modelcontextprotocol/go-sdk/auth"
	"github.com/modelcontextprotocol/go-sdk/oauthex"
)

type AuthorizationCodeHandler struct {
	cfg OAuthConfig
}

func NewAuthorizationCodeHandler(
	cfg OAuthConfig,
) (*sdkauth.AuthorizationCodeHandler, error) {
	callback := NewCallbackServer()
	fetcher := NewAuthorizationCodeFetcher(
		WindowsBrowser{},
		callback,
	)
	handler, err := sdkauth.NewAuthorizationCodeHandler(
		&sdkauth.AuthorizationCodeHandlerConfig{
			RedirectURL: cfg.RedirectURI,
			PreregisteredClient: &oauthex.ClientCredentials{
				ClientID: cfg.ClientID,
			},
			AuthorizationCodeFetcher: fetcher.Fetch,
		},
	)
	if err != nil {
		return nil, err
	}
	return handler, nil
}
