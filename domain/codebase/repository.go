package codebase

import "context"

type CodeBaseRepository interface {
	Save(
		ctx context.Context,
		doc *CodeBase,
	) error

	SearchSimilar(
		ctx context.Context,
		embedding []float32,
		limit int,
	) ([]CodeBase, error)
}
