package isogram

import "unicode"

// IsIsogram determines whether a word is an isogram
func IsIsogram(word string) bool {
	charMap := make(map[rune]bool)

	for _, char := range word {
		if !unicode.IsLetter(char) {
			continue
		}

		lowerCaseChar := unicode.ToLower(char)

		if charMap[lowerCaseChar] {
			return false
		}

		charMap[lowerCaseChar] = true
	}

	return true
}
