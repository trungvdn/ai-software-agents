package auth

import (
	"context"

	"golang.org/x/oauth2"
)

type TokenStore interface {
	Load(ctx context.Context) (*oauth2.Token, error)
	Save(ctx context.Context, token *oauth2.Token) error
}
