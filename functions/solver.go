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
You can make it dramatically faster with a couple of classic Sudoku-solver upgrades, without changing your overall approach.
1) Biggest win: choose the next empty cell smarter (MRV)
Right now you scan top-left to bottom-right and pick the first '.' . That can explode the search.
Instead, always pick the empty cell with the fewest possible candidates (Minimum Remaining Values). This single change often makes “hard/unsolvable” cases go from “minutes” to “instant”.

Drop-in replacement for Solve

package functions

func Solve(board *[9][9]rune) bool {
	// Find the empty cell with the fewest candidates
	bestR, bestC := -1, -1
	bestCount := 10 // more than max possible (9)

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] != '.' {
				continue
			}
			count := 0
			for v := '1'; v <= '9'; v++ {
				if CanPlace(board, r, c, v) {
					count++
				}
			}

			// dead end early: this cell has no legal values
			if count == 0 {
				return false
			}

			if count < bestCount {
				bestCount = count
				bestR, bestC = r, c
				if bestCount == 1 {
					// can't get better than 1
					break
				}
			}
		}
	}

	// No empty cells -> solved
	if bestR == -1 {
		return true
	}

	// Try candidates for the best cell
	for v := '1'; v <= '9'; v++ {
		if CanPlace(board, bestR, bestC, v) {
			board[bestR][bestC] = v
			if Solve(board) {
				return true
			}
			board[bestR][bestC] = '.'
		}
	}
	return false
}
2) Make CanPlace faster (optional but big)
Your CanPlace scans row+col+box every time. That’s a lot of repeated work.
Speed-up idea: maintain 3 boolean tables:
rowUsed[9][10]
colUsed[9][10]
boxUsed[9][10]
Then “can I place?” becomes O(1). This is a bigger refactor, but it’s the next major leap.
3) Add a quick “constraint propagation” loop (optional)
Before guessing, repeatedly fill any cell that has only 1 possible value. This reduces branching a lot.
*/
