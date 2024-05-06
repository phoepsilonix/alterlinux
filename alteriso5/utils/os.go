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

func GetFileSizesInDir(dir string) (map[string]int64, error) {
	items, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	ret := map[string]int64{}

	for _, i := range items {
		info, err := i.Info()
		if err != nil {
			continue
		}
		ret[info.Name()] = info.Size()
	}
	return ret, nil
}

func GetFileSize(file string) (int64, error) {
	info, err := os.Stat(file)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}
