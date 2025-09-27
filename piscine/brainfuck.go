package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	code := os.Args[1]
	memory := make([]byte, 2048)
	pointer := 0
	codePointer := 0

	for codePointer < len(code) {
		switch code[codePointer] {
		case '>':
			pointer++
			if pointer >= 2048 {
				pointer = 0 // Wrap around
			}
		case '<':
			pointer--
			if pointer < 0 {
				pointer = 2047 // Wrap around
			}
		case '+':
			memory[pointer]++
		case '-':
			memory[pointer]--
		case '.':
			z01.PrintRune(rune(memory[pointer]))
		case '[':
			if memory[pointer] == 0 {
				// Skip forward to matching ]
				bracketCount := 1
				for bracketCount > 0 && codePointer < len(code)-1 {
					codePointer++
					if code[codePointer] == '[' {
						bracketCount++
					} else if code[codePointer] == ']' {
						bracketCount--
					}
				}
			}
		case ']':
			if memory[pointer] != 0 {
				// Jump back to matching [
				bracketCount := 1
				for bracketCount > 0 && codePointer > 0 {
					codePointer--
					if code[codePointer] == ']' {
						bracketCount++
					} else if code[codePointer] == '[' {
						bracketCount--
					}
				}
			}
			// All other characters are ignored (comments)
		}
		codePointer++
	}
}
