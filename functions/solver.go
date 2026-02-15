package functions

// Solve attempts to solve the Sudoku board in-place using backtracking.
// It returns true when a complete valid solution is found, otherwise false.
func Solve(board *[9][9]rune) bool {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				for v := '1'; v <= '9'; v++ {
					if CanPlace(board, r, c, v) {
						board[r][c] = v
						if Solve(board) {
							return true
						}
						board[r][c] = '.'
					}
				}
				// no valid value found for this cell -> backtrack
				return false
			}
		}
	}
	// no empty cells -> solved
	return true
}

/*
Previous implementation (commented out to preserve original user code):

package functions

func solver() {
}
*/
