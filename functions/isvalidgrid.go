package functions
// isValidGrid verifies that the current board does not
// already violate Sudoku rules.
func IsValidGrid(board *[9][9]rune) bool {

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {

			value := board[row][col]

			// Skip empty cells
			if value == '.' {
				continue
			}

			// Temporarily remove the value so it doesn't
			// conflict with itself during validation
			board[row][col] = '.'

			// If placing it back is illegal, the grid is invalid
			if !CanPlace(board, row, col, value) {
				board[row][col] = value
				return false
			}

			// Restore original value
			board[row][col] = value
		}
	}

	return true
}
