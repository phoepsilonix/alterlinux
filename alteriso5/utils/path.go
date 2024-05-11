package utils

import (
	"path"
	"strings"
)

func Slash(elem ...string) string {
	arg := []string{}
	for _, p := range elem {
		arg = append(arg, strings.Split(p, "/")...)
	}
	return path.Join(arg...)
}
