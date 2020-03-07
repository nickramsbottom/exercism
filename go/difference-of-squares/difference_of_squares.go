package diffsquares

// SquareOfSum calculates square of the sum of numbers to n
func SquareOfSum(n int) int {
	sum := (n * (n + 1) / 2)
	return sum * sum
}

// SumOfSquares calculates sum of squares to n
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference calculate square of sum minus sum of squares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
