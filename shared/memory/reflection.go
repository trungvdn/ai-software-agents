package memory

import "time"

type Reflection struct {
	ID           string
	Content      string
	Importance   float64
	UsageCount   int
	LastAccessed time.Time
}
