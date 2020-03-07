package raindrops

import "strconv"

// Convert turns a number into raindrop sounds
func Convert(number int) string {
	droplets := ""

	if number%3 == 0 {
		droplets += "Pling"
	}

	if number%5 == 0 {
		droplets += "Plang"
	}

	if number%7 == 0 {
		droplets += "Plong"
	}

	if droplets == "" {
		return strconv.Itoa(number)
	}

	return droplets
}
