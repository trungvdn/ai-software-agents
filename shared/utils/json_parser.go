package utils

import (
	"strings"
)

func StripCodeFences(response string) string {
	// Strip markdown code blocks if present (e.g., ```json ... ```)
	jsonStr := strings.TrimSpace(response)
	if strings.HasPrefix(jsonStr, "```") {
		// Find the end of the opening code fence line
		lines := strings.Split(jsonStr, "\n")
		if len(lines) > 2 {
			// Skip the opening ``` line and rejoin
			jsonStr = strings.Join(lines[1:len(lines)-1], "\n")
		}
	}

	return strings.TrimSpace(jsonStr)
}
