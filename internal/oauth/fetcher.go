package oauth

import (
	"context"
	"fmt"

	sdkauth "github.com/modelcontextprotocol/go-sdk/auth"
)

type AuthorizationCodeFetcher struct {
	browser  Browser
	callback *CallbackServer
}

func NewAuthorizationCodeFetcher(
	browser Browser,
	callback *CallbackServer,
) *AuthorizationCodeFetcher {
	return &AuthorizationCodeFetcher{
		browser:  browser,
		callback: callback,
	}
}

func (a *AuthorizationCodeFetcher) Fetch(ctx context.Context, args *sdkauth.AuthorizationArgs) (*sdkauth.AuthorizationResult, error) {
	if err := a.callback.Start("127.0.0.1:8080"); err != nil {
		return nil, err
	}

	if err := a.browser.Open(args.URL); err != nil {
		fmt.Printf("Open the following URL in a browser to continue authorization:\n%s\n", args.URL)
	}
	results, err := a.callback.Wait(ctx)
	if err != nil {
		return nil, err
	}
	return results, nil
}
