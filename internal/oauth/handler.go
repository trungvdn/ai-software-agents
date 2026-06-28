package oauth

import (
	"errors"
	"log"
	"runtime"

	sdkauth "github.com/modelcontextprotocol/go-sdk/auth"
	"github.com/modelcontextprotocol/go-sdk/oauthex"
)

func getBrowserFactory() (Browser, error) {
	if runtime.GOOS == "windows" {
		return WindowsBrowser{}, nil
	}
	return nil, errors.New("not supported")
}

func NewAuthorizationCodeHandler(
	cfg OAuthConfig,
) (*sdkauth.AuthorizationCodeHandler, error) {
	callback := NewCallbackServer()
	log.Printf("getBrowserFactory")
	brs, err := getBrowserFactory()
	if err != nil {
		return nil, err
	}
	log.Printf("%T", brs)

	fetcher := NewAuthorizationCodeFetcher(
		brs,
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
