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
	rows := strings.Split(stringMatrix, "\n")

	var m Matrix

	for _, row := range rows {
		trimmedRow := strings.TrimSpace(row)
		splitRow := strings.Split(trimmedRow, " ")

		var splitRowInt []int

		for _, char := range splitRow {
			integer, ok := strconv.Atoi(char)
			if ok != nil {
				return nil, ok
			}
			splitRowInt = append(splitRowInt, integer)
		}

		m = append(m, splitRowInt)
	}

	previousLength := len(m[0])

	for i := 1; i < len(m); i++ {
		if len(m[i]) != previousLength {
			return nil, errors.New("uneven rows")
		}
	}

	return Matrix(m), nil
}

// Rows get the rows of the matrix
func (m Matrix) Rows() [][]int {
	var matrixCopy [][]int

	for i := range m {
		var matrixRow []int
		for j := range m[0] {
			matrixRow = append(matrixRow, m[i][j])
		}
		matrixCopy = append(matrixCopy, matrixRow)
	}

	return matrixCopy
}

// Cols get the columns of the matrix
func (m Matrix) Cols() [][]int {
	var matrixCopy [][]int

	for i := range m[0] {
		var matrixRow []int
		for j := range m {
			matrixRow = append(matrixRow, m[j][i])
		}
		matrixCopy = append(matrixCopy, matrixRow)
	}

	return matrixCopy
}

// Set a value in Matrix
func (m Matrix) Set(r, c, val int) bool {
	if r < 0 || c < 0 || r >= len(m) || c >= len(m[0]) {
		return false
	}

	m[r][c] = val
	return true
}
