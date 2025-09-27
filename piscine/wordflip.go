package piscine

import (
	"strings"
)

func WordFlip(str string) string {
	if str == "" {
		return "Invalid Output\n"
	}

	str = strings.TrimSpace(str)
	if str == "" {
		return "\n"
	}

	parts := strings.Split(str, " ")
	words := []string{}
	for _, w := range parts {
		if w != "" {
			words = append(words, w)
		}
	}

	// reverse words
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ") + "\n"
}
