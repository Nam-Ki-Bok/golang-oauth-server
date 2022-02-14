package utils

import (
	"strings"
)

func SplitScope(scopes string) []string {
	return strings.Split(scopes, " ")
}
