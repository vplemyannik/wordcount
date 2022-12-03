package app

import "strings"

func Calculate(text string) int {
	if len(text) == 0 {
		return 0
	}
	words := strings.Split(text, " ")
	return len(words)
}
