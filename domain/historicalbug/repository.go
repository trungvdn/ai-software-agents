package historicalbug

import "context"

type HistoricalBugRepository interface {
	Save(
		ctx context.Context,
		bug []*HistoricalBug,
	) error

	SearchSimilar(
		ctx context.Context,
		embedding []float32,
		limit int,
	) ([]HistoricalBug, error)
}
