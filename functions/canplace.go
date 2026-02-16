package functions

// canPlace checks if a value can be placed at (row, col)
// without breaking Sudoku rules.
func CanPlace(board *[9][9]rune, row, col int, value rune) bool {

	// Check the entire row
	for c := 0; c < 9; c++ {
		if board[row][c] == value {
			return false
		}
	}

	// Check the entire column
	for r := 0; r < 9; r++ {
		if board[r][col] == value {
			return false
		}
	}

	// Find the top-left corner of the 3x3 box
	startRow, startCol := (row/3)*3, (col/3)*3

	// Check the 3x3 box
	for r := startRow; r < startRow+3; r++ {
		for c := startCol; c < startCol+3; c++ {
			if board[r][c] == value {
				return false
			}
		}
	}

	// No conflicts found
	return true
}
