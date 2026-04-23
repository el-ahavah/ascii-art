package main

import (
	"fmt"
	"strings"
)

func printAscii(asciiMap map[rune][]string, input string) {
	if input == "\n" {
		fmt.Println()
		return
	}
	if input == "" {
		return
	}
	inputLines := strings.Split(input, "\n")

	for _, line := range inputLines {
		if line == "" {
			fmt.Println()
			continue
		}
		for i := 0; i < 8; i++ {
			for _, ch := range line {
				if ch < 32 || ch > 126 {
					ch = ' '
				}
				fmt.Print(asciiMap[ch][i])
			}
			fmt.Println()
		}
	}
}
