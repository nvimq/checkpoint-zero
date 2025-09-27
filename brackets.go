package main

import (
	"fmt"
	"os"
)

func isCorrectlyBracketed(s string) bool {
	stack := []rune{}

	// Map of closing brackets to opening brackets
	brackets := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		switch char {
		case '(', '[', '{':
			// Opening bracket - push to stack
			stack = append(stack, char)
		case ')', ']', '}':
			// Closing bracket - check if it matches
			if len(stack) == 0 {
				// No opening bracket to match
				return false
			}
			// Pop from stack and check if it matches
			lastOpening := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if lastOpening != brackets[char] {
				// Mismatched brackets
				return false
			}
		}
		// Ignore all other characters
	}

	// All brackets should be closed
	return len(stack) == 0
}

func main() {
	args := os.Args[1:]

	// If no arguments, print a newline
	if len(args) == 0 {
		fmt.Println()
		return
	}

	// Check each argument
	for _, arg := range args {
		if isCorrectlyBracketed(arg) {
			fmt.Println("OK")
		} else {
			fmt.Println("Error")
		}
	}
}
