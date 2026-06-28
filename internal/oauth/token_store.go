package oauth

import (
	"context"

	"golang.org/x/oauth2"
)

type TokenStore interface {
	Save(
		context.Context,
		*oauth2.Token,
	) error

	Load(
		context.Context,
	) (*oauth2.Token, error)
}
