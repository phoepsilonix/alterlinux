package utils

import (
	"path"
	"strings"
)

func Slash(base string, elem ...string) string {
	arg := []string{base}
	for _, p := range elem {
		arg = append(arg, strings.Split(p, "/")...)
	}
	return path.Join(arg...)
}
