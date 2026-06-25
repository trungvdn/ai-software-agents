package domain

import "context"

type ConfluenceClient interface {
	CreatePage(
		ctx context.Context,
		page Page,
	) (*Page, error)

	UpdatePage(
		ctx context.Context,
		page Page,
	) error
}
