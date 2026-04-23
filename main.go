package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run . \"text\" [standard.txt|shadow.txt|thinkertoy.txt]")
		os.Exit(1)
	}

	valid := map[string]bool{
		"standard.txt":   true,
		"shadow.txt":     true,
		"thinkertoy.txt": true,
	}

	if len(os.Args) > 2 && !valid[os.Args[2]] {
		fmt.Println("Invalid banner...")
		os.Exit(1)
	}
	bannerFile := "banners/standard.txt"

	if len(os.Args) == 3 {
		bannerFile = "banners/" + os.Args[2]
	}

	input := os.Args[1]
	input = strings.ReplaceAll(input, "\\n", "\n")

	asciiMap := buildAsciiMap(bannerFile)
	if asciiMap == nil {
		os.Exit(1)
	}

	printAscii(asciiMap, input)

}
