package main

import (
	"fmt"
	"os"
	"strings"
)

func standard() {
	if len(os.Args) != 2 {
		fmt.Println(`Usage: go run main.go "your txt"`)
		return
	}

	data, err := os.ReadFile("banner/standard.txt")
	if err != nil {
		fmt.Println("file not found!", err)
		return
	}

	lines := strings.Split(string(data), "\n")
	text := os.Args[1]
	val := strings.Split(text, "\\n")

	for _, va := range val {
		if va == "" {
			fmt.Println()
			continue
		}
		for row := 0; row < 8; row++ {
			for _, ch := range va {
				id := (int(ch-32)*9 + 1)
				fmt.Print(lines[id+row])
			}
			fmt.Println()
		}

	}
}
