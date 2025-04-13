package utils

import "strings"

func ExtractBearerToken(authHeader string) string {
	const prefix = "Bearer "
	tokenIndex := strings.Index(authHeader, prefix)
	if tokenIndex == -1 || tokenIndex != 0 {
		return ""
	}
	return authHeader[tokenIndex+len(prefix):]

}
