package codebase

import "context"

type CodeDocumentRepository interface {
	Save(
		ctx context.Context,
		doc *CodeDocument,
	) error

	SearchSimilar(
		ctx context.Context,
		embedding []float32,
		limit int,
	) ([]CodeDocument, error)
}
