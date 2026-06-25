package domain

import "context"

type Client interface {
	CreatePage(
		ctx context.Context,
		page Page,
	) error

	UpdatePage(
		ctx context.Context,
		page Page,
	) error
}
