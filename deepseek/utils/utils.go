package utils

import "strings"

func JoinEndpointParts(parts ...string) string {
	var cleanedParts []string
	for _, part := range parts {
		cleanedParts = append(cleanedParts, part)
	}

	return strings.Join(parts, "/")
}
