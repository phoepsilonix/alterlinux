package utils

import (
	"os"
	"strings"
)

func ReadFileLine(path string) ([]string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(bytes), "\n")
	return lines, nil
}
