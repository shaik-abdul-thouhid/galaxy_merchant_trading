package utils

import "strings"

func TrimQuestion(str, pattern string) string {
	str = strings.Trim(str, pattern)
	str = strings.ReplaceAll(str, "?", "")
	str = strings.TrimSpace(str)

	return str
}

func PrefixTrimQuestion(str, pattern string) string {
	str = strings.TrimPrefix(str, pattern)
	str = strings.ReplaceAll(str, "?", "")
	str = strings.TrimSpace(str)

	return str
}

func SuffixTrimQuestion(str, pattern string) string {
	str = strings.TrimSuffix(str, pattern)
	str = strings.ReplaceAll(str, "?", "")
	str = strings.TrimSpace(str)

	return str
}
