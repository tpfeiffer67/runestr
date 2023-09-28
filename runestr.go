// The runestr package offers several functions that operate
// on strings at the rune level.
// These functions allow for tasks such as calculating the length of a string (count the number of runes),
// padding a string on the right side, etc.
package runestr

import (
	"strings"
	"unicode/utf8"
)

// Length returns the number of Unicode code points (runes) in a string.
func Length(s string) int {
	return utf8.RuneCountInString(s)
}

// TODO Use a stringbuilder

// PadRight pads a string on the right side with a fill pattern until it reaches the given width.
// If the fill pattern doesn't fit exactly, the last character(s) of the fill pattern are repeated.
func PadRight(s string, fillPattern string, width int) string {
	// If the fill pattern is empty, return the original string.
	if len(fillPattern) < 1 {
		return s
	}
	// Calculate the current length of the string.
	currentLength := Length(s)
	// If the current length is less than the desired width, pad the string with the fill pattern.
	if currentLength < width {
		// Repeat the fill pattern as many times as necessary to fill the remaining space.
		padCount := (width - currentLength) / Length(fillPattern)
		paddedString := s + strings.Repeat(fillPattern, padCount)
		// If the padded string still doesn't fit exactly, add the last character(s) of the fill pattern.
		for Length(paddedString) < width {
			r, size := utf8.DecodeRuneInString(fillPattern)
			paddedString += string(r)
			fillPattern = fillPattern[size:]
		}
		return paddedString
	}
	// If the current length is equal to or greater than the desired width, return the original string.
	return s
}

// Left returns the leftmost n runes of a string.
func Left(s string, n int) string {
	runes := []rune(s)
	if len(runes) > n {
		return string(runes[:n])
	}
	return s
}

// Right returns the rightmost n runes of a string.
func Right(s string, n int) string {
	runes := []rune(s)
	length := len(runes)
	if length > n {
		return string(runes[length-n:])
	}
	return s
}

// SplitOnNearestSpace splits a string into two parts at the nearest space character to the n-th position.
// The first part includes all characters up to (but not including) the space character. If there is no space character, the word can be truncated.
// The second part includes the rest of the stringwith leading and trailing spaces removed.
func SplitOnNearestSpace(s string, max int) (first string, second string) {
	// Trim leading and trailing spaces from the string.
	trimmed := strings.TrimSpace(s)
	// If the string is already shorter than n, return the whole string as the first part.
	if Length(trimmed) <= max {
		return trimmed, ""
	}
	runes := []rune(trimmed)
	// Look for the nearest space character to the n-th position.
	for i := max; i >= 1; i-- {
		if runes[i] == ' ' && runes[i-1] != ' ' {
			first = string(runes[:i])
			second = strings.TrimSpace(string(runes[i+1:]))
			return
		}
	}
	// If no space character was found, return the first n characters as the first part.
	first = string(runes[:max])
	second = strings.TrimSpace(string(runes[max:]))
	return
}

// RuneAtPosition returns the Unicode code point (rune) at the specified position in a string,
// or the specified default value if the position is out of bounds.
func RuneAtPosition(s string, pos int, def rune) rune {
	rstr := []rune(s)
	if len(rstr) > pos {
		return rstr[pos]
	}
	return def
}
