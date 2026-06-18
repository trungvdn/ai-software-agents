package historicalbug

import (
	"time"

	"github.com/google/uuid"
)

type HistoricalBug struct {
	ID         uuid.UUID
	Title      string
	RootCause  string
	FixSummary string
	Severity   string // LOW,MEDIUM,HIGH,CRITICAL
	UsageCount int
	Embedding  []float32
	CreatedAt  time.Time
}
