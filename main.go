package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"sudoku/functions"
)

func main() {
	if functions.CheckInput(os.Args) == false {
		fmt.Println(`Error: Wrong input.
		
Please enter up to 9 Sudoku rows.
Each row must have exactly 9 characters.
Use numbers 1-9 for known cells and . for empty cells.
		
You can write rows normally:
534.7....
		
Or set a specific row:
1=534.7....
		
You may also combine both methods`)
		return
	}

	var board [9][9]rune
	var given [9][9]bool
	// initialize all cells to '.'
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			board[r][c] = '.'
			given[r][c] = false
		}
	}

	rowsFilled := [9]bool{}
	positional := []string{}

	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		if strings.Contains(arg, "=") {
			parts := strings.SplitN(arg, "=", 2)

			idx, _ := strconv.Atoi(parts[0])
			rowStr := parts[1]

			for j, ch := range rowStr {
				board[idx-1][j] = ch
				if ch != '.' {
					given[idx-1][j] = true
				}
			}
			rowsFilled[idx-1] = true
		} else {
			positional = append(positional, arg)
		}
	}

	// assign positional rows to next free rows
	posIdx := 0
	for r := 0; r < 9 && posIdx < len(positional); r++ {
		if !rowsFilled[r] {
			rowStr := positional[posIdx]
			for j, ch := range rowStr {
				board[r][j] = ch
				if ch != '.' {
					given[r][j] = true
				}
			}
			rowsFilled[r] = true
			posIdx++
		}
	}

	if !functions.IsValidGrid(&board) {
		fmt.Println(`Error: Wrong input.

This Sudoku board breaks the game rules.
Check for repeated numbers in a row, column or 3x3 box.`)
		return
	}

	if !functions.Solve(&board) {
		fmt.Println(`Error: No solution found.
		
This Sudoku puzzle cannot be solved.`)
		return
	}

	// Print solved board in ASCII box format (user-provided numbers in green)
	const green = "\x1b[92m"
	const reset = "\x1b[0m"

	cell := func(rr, cc int) string {
		ch := board[rr][cc]
		if ch == '.' {
			return "."
		}
		if given[rr][cc] {
			return fmt.Sprintf("%s%c%s", green, ch, reset)
		}
		return fmt.Sprintf("%c", ch)
	}

	for r := 0; r < 9; r++ {
		if r%3 == 0 {
			fmt.Println("o-------o-------o-------o")
		}
		// Print row with spacing and box separators
		fmt.Printf("| %s %s %s | %s %s %s | %s %s %s |\n",
			cell(r, 0), cell(r, 1), cell(r, 2),
			cell(r, 3), cell(r, 4), cell(r, 5),
			cell(r, 6), cell(r, 7), cell(r, 8),
		)
	}
	fmt.Println("o-------o-------o-------o")
}

// Previous commented code (moved here to preserve originals):
//
// package main
//
// // Import fmt and os by John
// import (
// 	"fmt"
// 	"os"
// 	"sudoku/functions"
// )
//
// func main() {
//
// 	// First, the code checks for correct input through the CheckInput function
// 	if functions.CheckInput(os.Args) == false {
// 		fmt.Println("Error: Wrong Input")
// 		return
// 	}
//
// 	//if input passes the check, we continue to make the board
// 	// we create a 9x9 board
// 	var board [9][9]rune
//
// 	for i := 1; i <= 9; i++ { // Loop starts from 1 because os.Args[0] is the program's name.
//
// 		row := os.Args[i]
//
// 		for index, r := range row {
// 			// "range" splits the string:
// 			// 'index' is the position of the character (0-8)
// 			// 'r' is the actual character (rune) at that position
//
// 			// Now we assign the character to the 9x9 board.
// 			// We use [i-1] as the Row position because the code's position starts from 0. so in a 9 row line, the positions are 0-8.
// 			// For Column position, we use index
// 			board[i-1][index] = r
//
// 		}
// 	}
//
// 	if !functions.IsValidGrid(&board) {
// 		fmt.Println("Error: Wrong Input")
// 		return
// 	}
//
// 	// Solve the board
// 	if !functions.Solve(&board) {
// 		fmt.Println("Error: N-o -solution")
// 		return
// 	}
//
// 	// Print solved board (rows of numbers separated by spaces)
// 	for r := 0; r < 9; r++ {
// 		for c := 0; c < 9; c++ {
// 			if c > 0 {
// 				fmt.Print(" ")
// 			}
// 			fmt.Printf("%c", board[r][c])
// 		}
// 		fmt.Println()
// 	}
//
// }
//
// /* Print example.
//
// o-------o-------o-------o
// | 1 2 3 | 4 5 6 | 7 8 9 |
// | 4 5 6 | 7 8 9 | 1 2 3 |
// | 7 8 9 | 1 2 3 | 4 5 6 |
// o-------o-------o-------o
// | 9 1 2 | 3 4 5 | 6 7 8 |
// | 3 4 5 | 6 7 8 | 9 1 2 |
// | 6 7 8 | 9 1 2 | 3 4 5 |
// o-------o-------o-------o
// | 8 9 1 | 2 3 4 | 5 6 7 |
// | 2 3 4 | 5 6 7 | 8 9 1 |
// | 5 6 7 | 8 9 1 | 2 3 4 |
// o-------o-------o-------o
// */
