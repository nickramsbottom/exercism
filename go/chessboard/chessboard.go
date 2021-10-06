package chessboard

// Rank stores if a square is occupied by a piece
type Rank []bool

// Chessboard contains a map of eight Ranks, accessed with values from 1 to 8
type Chessboard map[int]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank int) (ret int) {
	for _, value := range cb[rank] {
		if value {
			ret++
		}
	}

	return ret
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) (ret int) {
	if file > len(cb[1]) {
		return 0
	}

	for i := 1; i <= len(cb); i++ {
		if cb[i][file-1] {
			ret++
		}
	}
	return ret
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	return len(cb) * len(cb)
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) (ret int) {
	for _, row := range cb {
		for _, sq := range row {
			if sq {
				ret++
			}
		}
	}
	return ret
}
