package shared

import (
	"os"
	"strings"
)

func ReadLines(path string) (out []string) {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	fileText := string(file)
	lines := strings.Split(fileText, "\n")

	for _, line := range lines {
		if len(strings.TrimSpace(line)) >= 0 {
			out = append(out, line)
		}
	}

	return out
}
