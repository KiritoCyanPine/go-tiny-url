package config

import "strings"

func isEmptyOrWhiteSpace(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}
