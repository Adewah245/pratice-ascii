package main

import (
	"fmt"
	"strings"
)

func PrintBanner(text string, lines []string) {

	step1 := strings.Split(text, "\\n")

	for _, step := range step1 {
		if step == "" {
			fmt.Println()
			continue
		}
		for row := 0; row < 8; row++ {
			for _, ch := range step {
				id := (int(ch-32)*9 + 1)
				fmt.Print(lines[id+row])
			}
			fmt.Println()
		}
	}
}
