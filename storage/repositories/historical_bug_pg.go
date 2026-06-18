package repositories

import (
	"context"
	"database/sql"

	"github.com/trungvdn/ai-software-agents/domain/historicalbug"
	"github.com/trungvdn/ai-software-agents/shared/utils"
)

type HistoricalBugRepository struct {
	db *sql.DB
}

func NewHistoricalBugRepository(
	db *sql.DB,
) *HistoricalBugRepository {
	return &HistoricalBugRepository{
		db: db,
	}
}

func (r *HistoricalBugRepository) Save(
	ctx context.Context,
	bug []*historicalbug.HistoricalBug,
) error {
	query := `
		INSERT INTO historical_bugs (id, title, root_cause, fix_summary, severity, usage_count, embedding, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7::vector, $8)
		ON CONFLICT (id) DO UPDATE SET
			title = EXCLUDED.title,
			root_cause = EXCLUDED.root_cause,
			fix_summary = EXCLUDED.fix_summary,
			severity = EXCLUDED.severity,
			usage_count = EXCLUDED.usage_count,
			embedding = EXCLUDED.embedding
	`

	// Convert embedding to pgvector format

	// Use tx.BeginTx(...)

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, b := range bug {
		embeddingStr := utils.VectorToString(b.Embedding)

		_, err := tx.ExecContext(ctx, query, b.ID, b.Title, b.RootCause, b.FixSummary, b.Severity, b.UsageCount, embeddingStr, b.CreatedAt)
		if err != nil {
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (r *HistoricalBugRepository) SearchSimilar(
	ctx context.Context,
	embedding []float32,
	limit int,
) ([]historicalbug.HistoricalBug, error) {
	query := `
		SELECT id, title, root_cause, fix_summary, severity, usage_count
		FROM historical_bugs
		ORDER BY (embedding <=> $1::vector) ASC
		LIMIT $2
	`
	embeddingStr := utils.VectorToString(embedding)

	rows, err := r.db.QueryContext(ctx, query, embeddingStr, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bugs []historicalbug.HistoricalBug
	for rows.Next() {
		var bug historicalbug.HistoricalBug
		if err := rows.Scan(&bug.ID, &bug.Title, &bug.RootCause, &bug.FixSummary, &bug.Severity, &bug.UsageCount); err != nil {
			return nil, err
		}
		bugs = append(bugs, bug)
	}
	return bugs, nil
}

func (r *HistoricalBugRepository) IncrementUsageCount(
	ctx context.Context,
	id []string,
) error {
	query := `
		UPDATE historical_bugs
		SET usage_count = usage_count + 1, last_accessed = NOW()
		WHERE id = ANY($1)
	`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
