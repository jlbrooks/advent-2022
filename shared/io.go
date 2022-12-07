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

func ReadContinuous(path string) (out []rune) {
	lines := ReadLines(path)
	for _, l := range lines {
		out = append(out, []rune(l)...)
		out = append(out, '\n')
	}

	return out
}
