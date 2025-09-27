package piscine

import "strings"

// FifthAndSkip processes a string by inserting a space after every fifth
// non-space character and then skipping the sixth non-space character.
// It returns "Invalid Input\n" if the string has fewer than 5 non-space
// characters, and "\n" if the input string is empty.
func FifthAndSkip(str string) string {
	// First, handle the edge case for an empty string.
	if len(str) == 0 {
		return "\n"
	}

	// Create a sanitized string by removing all spaces. This is the string
	// we will iterate over to apply the core logic.
	sanitizedStr := strings.ReplaceAll(str, " ", "")

	// Handle the edge case for strings with fewer than 5 non-space characters.
	if len(sanitizedStr) < 5 {
		return "Invalid Input\n"
	}

	// Convert the sanitized string to a slice of runes to get the correct
	// number of characters, which is crucial for handling multi-byte runes
	// and accurately checking the length.
	sanitizedRunes := []rune(sanitizedStr)
	sanitizedLen := len(sanitizedRunes)

	// Use a strings.Builder for efficient string concatenation.
	var result strings.Builder

	// `processedChars` counts all non-space characters from the original string,
	// including those that will be skipped. This is crucial for correctly
	// identifying the 6th character to be skipped.
	var processedChars int
	// `outputtedChars` counts the characters that have been successfully
	// appended to the `result` string. This is used to place the spaces correctly.
	var outputtedChars int

	// The `for...range` loop correctly iterates over runes.
	for _, r := range sanitizedRunes {
		processedChars++

		// Check if the current character is the sixth one in a sequence.
		// If so, we skip it by continuing to the next iteration.
		if processedChars%6 == 0 {
			continue
		}

		// Append the character to the result.
		result.WriteRune(r)
		outputtedChars++

		// After appending, check if we have completed a block of 5 characters.
		// We add a space if the count is a multiple of 5, but we must also ensure
		// that there are more characters left to be outputted to avoid a
		// trailing space. We check against the rune length, not the byte length.
		if outputtedChars%5 == 0 && processedChars < sanitizedLen {
			result.WriteString(" ")
		}
	}

	// The logic above ensures no trailing space is added.
	// We simply append the final newline as per the requirements.
	return result.String() + "\n"
}
