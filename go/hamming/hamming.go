package hamming

import "errors"

var errDiffLength = errors.New("strings must be of equal length")

// Distance returns the hamming distance between two DNA strings
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errDiffLength
	}

	hamming := 0

	ar, br := []rune(a), []rune(b)

	for i := range ar {
		if ar[i] != br[i] {
			hamming++
		}
	}

	return hamming, nil
}
