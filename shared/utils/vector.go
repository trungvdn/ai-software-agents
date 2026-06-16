package utils

import (
	"fmt"
	"strings"
)

// VectorToString converts a float32 slice to pgvector format
// pgvector expects vectors in format: "[0.1,0.2,0.3]"
func VectorToString(embedding []float32) string {
	if len(embedding) == 0 {
		return "[]"
	}

	var sb strings.Builder
	sb.WriteString("[")

	for i, val := range embedding {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(fmt.Sprintf("%f", val))
	}

	sb.WriteString("]")
	return sb.String()
}
