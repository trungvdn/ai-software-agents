package codebase

import (
	"time"

	"github.com/google/uuid"
)

type CodeDocument struct {
	ID        uuid.UUID
	FilePath  string
	Content   string
	Embedding []float32
	Language  string
	UpdatedAt *time.Time
	CreatedAt time.Time
}
