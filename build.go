package main

import (
	"fmt"
	"os"
	"strings"
)

func buildAsciiMap(bannerFile string) (map[rune][]string, error) {
	file, err := os.ReadFile(bannerFile)
	if err != nil {
		fmt.Println("Error reading banner:", err)
		return nil, fmt.Errorf("error reading banner: %w", err)
	}

	lines := strings.Split(string(file), "\n")
	asciiMap := make(map[rune][]string)

	for i := 32; i <= 126; i++ {
		start := (i-32)*9 + 1
		end := start + 8

		if end > len(lines) {
			return nil, fmt.Errorf("banner file is malformed or incomplete")
		}

		asciiMap[rune(i)] = lines[start:end]
	}

	return asciiMap
}
