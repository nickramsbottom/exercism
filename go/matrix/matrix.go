package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix represents a two dimensional matrix
type Matrix [][]int

// New make a matrix
func New(stringMatrix string) (Matrix, error) {
	var m Matrix

	for _, row := range strings.Split(stringMatrix, "\n") {
		trimRow := strings.TrimSpace(row)

		var rowInts []int

		for _, char := range strings.Split(trimRow, " ") {
			integer, ok := strconv.Atoi(char)
			if ok != nil {
				return nil, ok
			}
			rowInts = append(rowInts, integer)
		}

		if len(m) > 0 && len(rowInts) != len(m[0]) {
			return nil, errors.New("uneven rows")
		}

		m = append(m, rowInts)
	}

	return Matrix(m), nil
}

// Rows get the rows of the matrix
func (m Matrix) Rows() [][]int {
	var mCopy [][]int

	for i := range m {
		var row []int
		for j := range m[0] {
			row = append(row, m[i][j])
		}
		mCopy = append(mCopy, row)
	}

	return mCopy
}

// Cols get the columns of the matrix
func (m Matrix) Cols() [][]int {
	var mCopy [][]int

	for i := range m[0] {
		var row []int
		for j := range m {
			row = append(row, m[j][i])
		}
		mCopy = append(mCopy, row)
	}

	return mCopy
}

// Set a value in Matrix
func (m Matrix) Set(r, c, val int) bool {
	if r < 0 || c < 0 || r >= len(m) || c >= len(m[0]) {
		return false
	}

	m[r][c] = val
	return true
}
