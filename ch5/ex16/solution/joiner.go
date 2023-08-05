package joiner

import "strings"

func Join(sep string, list ...string) string {
	return strings.Join(list, sep)
}