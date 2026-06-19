package repositories

import (
	"context"
	"database/sql"

	"github.com/trungvdn/ai-software-agents/domain/codebase"
	"github.com/trungvdn/ai-software-agents/shared/utils"
)

type CodeBaseRepository struct {
	db *sql.DB
}

func NewCodeBaseRepository(
	db *sql.DB,
) *CodeBaseRepository {
	return &CodeBaseRepository{
		db: db,
	}
}

func (r *CodeBaseRepository) Save(
	ctx context.Context,
	codebase *codebase.CodeBase,
) error {
	query := `
		INSERT INTO codebase (id, file_path, content, embedding, language, created_at, updated_at)
		VALUES ($1, $2, $3, $4::vector, $5, $6, $7)
		ON CONFLICT (id) DO UPDATE SET
			file_path = EXCLUDED.file_path,
			content = EXCLUDED.content,
			embedding = EXCLUDED.embedding,
			language = EXCLUDED.language,
			created_at = EXCLUDED.created_at,
			updated_at = EXCLUDED.updated_at,
			embedding = EXCLUDED.embedding
	`

	// Convert embedding to pgvector format
	embeddingStr := utils.VectorToString(codebase.Embedding)

	_, err := r.db.ExecContext(ctx, query,
		codebase.ID,
		codebase.FilePath,
		codebase.Content,
		embeddingStr,
		codebase.Language,
		codebase.CreatedAt,
		codebase.UpdatedAt,
	)

	return err
}

func (r *CodeBaseRepository) SearchSimilar(
	ctx context.Context,
	embedding []float32,
	limit int,
) ([]codebase.CodeBase, error) {
	query := `
		SELECT id, file_path, content, embedding, language
		FROM codebase
		ORDER BY (embedding <=> $1::vector) ASC
		LIMIT $2
	`

	// Convert embedding to pgvector format
	embeddingStr := utils.VectorToString(embedding)

	rows, err := r.db.QueryContext(ctx, query, embeddingStr, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []codebase.CodeBase
	for rows.Next() {
		var cb codebase.CodeBase
		if err := rows.Scan(&cb.ID, &cb.FilePath, &cb.Content, &cb.Embedding, &cb.Language); err != nil {
			return nil, err
		}
		results = append(results, cb)
	}
	return results, nil
}
