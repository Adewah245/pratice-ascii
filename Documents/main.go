package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(`usage: go run . "your text"`)
		return
	}

	data, err := os.ReadFile("banner/thinkertoy.txt")
	if err != nil {
		fmt.Println("error:", err)
	}
	lines := strings.Split(string(data), "\n")
	result := rstringBanner(os.Args[1], lines)
	fmt.Print(result)

}
