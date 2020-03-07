package luhn

import (
	"strings"
)

// Valid determines whether a string confirms to the Luhn algorithm
func Valid(input string) bool {
	strippedInput := strings.ReplaceAll(input, " ", "")
	inputLength := len(strippedInput)

	if inputLength <= 1 {
		return false
	}

	isDouble := len(strippedInput)%2 == 0
	sum := 0
	for _, char := range strippedInput {
		integer := int(char - '0')
		if integer < 0 || integer > 9 {
			return false
		}
		if isDouble {
			integer *= 2
			if integer > 9 {
				integer -= 9
			}
		}
		sum += integer
		isDouble = !isDouble
	}
	return sum%10 == 0
}
