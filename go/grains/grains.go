package grains

import (
	"errors"
)

// Square returns how many grains are on any square
func Square(n int) (uint64, error) {
	if 1 > n || n > 64 {
		return 0, errors.New("Invalid")
	}

	return 1 << (n - 1), nil
}

// Total returns the total number of grains on a 64 square board
func Total() uint64 {
	return 1<<64 - 1
}
