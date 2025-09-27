package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	oldRunes := []rune(os.Args[2])
	newRunes := []rune(os.Args[3])
	if len(oldRunes) != 1 || len(newRunes) != 1 {
		return
	}
	old, new := oldRunes[0], newRunes[0]
	for _, r := range []rune(os.Args[1]) {
		if r == old {
			z01.PrintRune(new)
		} else {
			z01.PrintRune(r)
		}
	}
	z01.PrintRune('\n')
}
