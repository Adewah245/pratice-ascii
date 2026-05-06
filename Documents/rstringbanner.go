package main

import "strings"

func rstringBanner(text string, lines []string) string{
	// build magic house
	var result strings.Builder
	// use magic split incoming text
	step1 := strings.Split(text, "\\n")
	// sending our junior brother
	for _, step := range step1 {
		// check if you found empty box bring it like that
		if step == ""{
			// give it to our magic builder
			result.WriteString("\n")
			// go back again
			continue
		}
		// enter inside if you see any thing from the do post count
		for row := 0; row < 8; row++ {
			// get everything inside what you count
			for _, ch := range step {
				// what and what you found when counting
				id := (int(ch-32)*9+1)
				// send everything to our maggic builder
				result.WriteString(lines[id+row])
			}
			// write it in newline
			result.WriteString("\n")
		}
	}
	return result.String()

}