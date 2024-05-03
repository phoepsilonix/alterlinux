package utils

import "strings"

func LoadEnvFile(path string) (map[string]string, error) {
	lines, err := ReadFileLine(path)
	if err != nil {
		return nil, err
	}

	env := make(map[string]string)

	for _, line := range lines {
		if strings.HasPrefix(line, "#") || !strings.Contains(line, "=") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		value := strings.TrimPrefix(parts[1], "\"")
		value = strings.TrimSuffix(value, "\"")
		env[parts[0]] = value
	}

	return env, nil

}
