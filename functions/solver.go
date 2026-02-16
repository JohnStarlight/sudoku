package functions

// Solve using:
// - backtracking
// - MRV (fewest candidates first)
// - simple readable logic
func Solve(board *[9][9]rune) bool {
	return dfs(board)
}

func dfs(board *[9][9]rune) bool {
	// Find the empty cell with the fewest candidates
	r, c, candidates, found := findBestCell(board)

	// no empty cells -> solved
	if !found {
		return true
	}

	// no candidates -> dead end
	if len(candidates) == 0 {
		return false
	}
	// try candidates for this cell
	for _, v := range candidates {
		board[r][c] = v
		// Recurse to solve the rest of the board with this value placed.
		if dfs(board) {
			return true
		}
		// Backtrack
		board[r][c] = '.'
	}

	return false
}

// ----------------------------------------------------
// Find empty cell with fewest legal values (MRV)
// returns row, col, candidates, found
// ----------------------------------------------------
func findBestCell(board *[9][9]rune) (int, int, []rune, bool) {
	bestCount := 10
	bestR, bestC := -1, -1
	var bestCandidates []rune
	// Loop through all cells to find the one with the fewest legal values
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] != '.' {
				continue
			}
			// Get candidates for this cell
			cands := getCandidates(board, r, c)
			// If no legal values, this path is a dead end
			if len(cands) < bestCount {
				bestCount = len(cands)
				bestCandidates = cands
				bestR, bestC = r, c

				if bestCount == 1 {
					return bestR, bestC, bestCandidates, true
				}
			}
		}
	}
	// No empty cells -> solved
	if bestR == -1 {
		return 0, 0, nil, false
	}

	return bestR, bestC, bestCandidates, true
}

// ----------------------------------------------------
// Collect allowed digits once (no repeated checks)
// ----------------------------------------------------
func getCandidates(board *[9][9]rune, r, c int) []rune {
	used := [9]bool{} // index 0 = digit 1

	// row
	for i := 0; i < 9; i++ {
		if board[r][i] != '.' {
			used[board[r][i]-'1'] = true
		}
	}

	// col
	for i := 0; i < 9; i++ {
		if board[i][c] != '.' {
			used[board[i][c]-'1'] = true
		}
	}

	// box
	br, bc := (r/3)*3, (c/3)*3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[br+i][bc+j] != '.' {
				used[board[br+i][bc+j]-'1'] = true
			}
		}
	}
	// Collect unused digits as candidates
	var res []rune
	for i := 0; i < 9; i++ {
		if !used[i] {
			res = append(res, rune('1'+i))
		}
	}

	return res
}
