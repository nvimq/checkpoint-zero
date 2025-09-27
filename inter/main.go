package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	s1 := []rune(os.Args[1])
	s2 := []rune(os.Args[2])
	seen := make(map[rune]bool)

	// Check if rune exists in s2
	inSecond := func(r rune) bool {
		for _, c := range s2 {
			if c == r {
				return true
			}
		}
		return false
	}

	for _, r := range s1 {
		if !seen[r] && inSecond(r) {
			z01.PrintRune(r)
			seen[r] = true
		}
	}
	z01.PrintRune('\n')
}
