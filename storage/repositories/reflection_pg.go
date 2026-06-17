package repositories

import (
	"context"
	"database/sql"

	"github.com/trungvdn/ai-software-agents/domain/reflection"
	"github.com/trungvdn/ai-software-agents/shared/utils"
)

type ReflectionRepository struct {
	db *sql.DB
}

func NewReflectionRepository(
	db *sql.DB,
) *ReflectionRepository {
	return &ReflectionRepository{
		db: db,
	}
}

func (r *ReflectionRepository) Save(
	ctx context.Context,
	reflectionData reflection.Reflection,
) error {
	query := `
		INSERT INTO reflections (id, content, importance_score, usage_count, last_accessed, embedding, created_at)
		VALUES ($1, $2, $3, $4, $5, $6::vector, $7)
		ON CONFLICT (id) DO UPDATE SET
			content = EXCLUDED.content,
			importance_score = EXCLUDED.importance_score,
			usage_count = EXCLUDED.usage_count,
			last_accessed = EXCLUDED.last_accessed,
			embedding = EXCLUDED.embedding
	`

	// Convert embedding to pgvector format
	embeddingStr := utils.VectorToString(reflectionData.Embedding)

	_, err := r.db.ExecContext(ctx, query,
		reflectionData.ID,
		reflectionData.Content,
		reflectionData.ImportanceScore,
		reflectionData.UsageCount,
		reflectionData.LastAccessed,
		embeddingStr,
		reflectionData.CreatedAt,
	)

	return err
}

func (r *ReflectionRepository) SearchSimilar(
	ctx context.Context,
	embedding []float32,
	limit int,
) ([]reflection.SimilarReflection, error) {
	query := `
		SELECT id, content, 1 - (embedding <=> $1) as similarity, usage_count, importance_score
		FROM reflections
		ORDER BY (embedding <-> $1::vector) ASC
		LIMIT $2
	`

	// Convert embedding to pgvector format
	embeddingStr := utils.VectorToString(embedding)

	rows, err := r.db.QueryContext(ctx, query, embeddingStr, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []reflection.SimilarReflection
	for rows.Next() {
		var refl reflection.SimilarReflection
		if err := rows.Scan(&refl.Reflection.ID, &refl.Reflection.Content, &refl.Similarity, &refl.Reflection.UsageCount, &refl.Reflection.ImportanceScore); err != nil {
			return nil, err
		}
		results = append(results, refl)
	}
	return results, nil
}

func (r *ReflectionRepository) IncrementUsageCount(
	ctx context.Context,
	id string,
) error {
	query := `
		UPDATE reflections
		SET usage_count = usage_count + 1, last_accessed = NOW()
		WHERE id = $1
	`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
