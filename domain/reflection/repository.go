package reflection

import "context"

type ReflectionRepository interface {
	Save(
		ctx context.Context,
		reflection Reflection,
	) error

	SearchSimilar(
		ctx context.Context,
		embedding []float32,
		limit int,
	) ([]SimilarReflection, error)
}
