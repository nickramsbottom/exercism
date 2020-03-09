package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency finds the frequency of letters in multiple input texts
func ConcurrentFrequency(texts []string) FreqMap {
	results := FreqMap{}
	resultChannel := make(chan FreqMap, 10)

	for _, text := range texts {
		go func(s string) {
			resultChannel <- Frequency(s)
		}(text)
	}

	for range texts {
		result := <-resultChannel
		for r, count := range result {
			results[r] += count
		}
	}

	return results
}
