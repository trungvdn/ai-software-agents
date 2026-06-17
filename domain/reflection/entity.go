package reflection

import (
	"time"

	"github.com/google/uuid"
)

type Reflection struct {
	ID              uuid.UUID
	Content         string
	ImportanceScore float64
	UsageCount      int
	LastAccessed    *time.Time
	CreatedAt       time.Time
	Embedding       []float32
}

type SimilarReflection struct {
	Reflection Reflection
	Similarity float64
}
