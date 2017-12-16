package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram determines whether or not a string and duplicated character
func IsIsogram(input string) bool {
	distinctMap := make(map[rune](bool))
	lowerInput := strings.ToLower(input)
	for _, c := range lowerInput {
		if unicode.IsLetter(c) {
			if _, ok := distinctMap[c]; !ok {
				distinctMap[c] = true
			} else {
				return false
			}
		}
	}
	return true
}
