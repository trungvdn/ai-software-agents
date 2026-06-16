package memory

import "time"

type Episode struct {
	ID        string
	Bug       string
	RootCause string
	Fix       string
	CreatedAt time.Time
}
