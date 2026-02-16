package main

import (
	"fmt"
	"os"

	// strconv is used to convert string row indices to integers.
	"strconv"
	// strings is used to parse command-line arguments for explicit row assignments.
	"strings"

	"sudoku/functions"
)

func main() {
	if !functions.CheckInput(os.Args) {
		printInputError()
		return
	}

	board, given := buildBoard(os.Args)

	if !functions.IsValidGrid(&board) {
		printGridError()
		return
	}

	if !functions.Solve(&board) {
		printNoSolution()
		return
	}

	printBoard(board, given)
}

//
// -------------------- BUILD BOARD --------------------
//

// initBoard initializes the board with '.' in all cells.
func initBoard(board *[9][9]rune) {
	for r := range board {
		for c := range board[r] {
			board[r][c] = '.'
		}
	}
}

// fillRow fills a specific row of the board with the provided string and marks given cells.
func fillRow(board *[9][9]rune, given *[9][9]bool, r int, rowStr string) {
	for c, ch := range rowStr {
		board[r][c] = ch
		if ch != '.' {
			given[r][c] = true
		}
	}
}

// buildBoard constructs the Sudoku board and given cells from command-line arguments.
func buildBoard(args []string) ([9][9]rune, [9][9]bool) {
	var board [9][9]rune
	var given [9][9]bool
	var rowsFilled [9]bool
	// Initialize the board with '.' for empty cells
	initBoard(&board)

	positional := []string{}

	for i := 1; i < len(args); i++ {
		arg := args[i]
		// Check if the argument is in the form "row=..." to fill a specific row
		if strings.Contains(arg, "=") {
			parts := strings.SplitN(arg, "=", 2)
			// Validate the row index and fill the specified row
			idx, _ := strconv.Atoi(parts[0])
			fillRow(&board, &given, idx-1, parts[1])
			// Mark this row as filled from explicit input
			rowsFilled[idx-1] = true
		} else {
			// Collect positional rows to fill later
			positional = append(positional, arg)
		}
	}

	// assign positional rows
	pos := 0
	for r := 0; r < 9 && pos < len(positional); r++ {
		if !rowsFilled[r] {
			fillRow(&board, &given, r, positional[pos])
			pos++
		}
	}

	return board, given
}

//
// -------------------- PRINT --------------------
//

// printBoard displays the Sudoku board in a formatted way, highlighting given cells in green.
func printBoard(board [9][9]rune, given [9][9]bool) {
	const green = "\x1b[92m"
	const reset = "\x1b[0m"
	// cell returns the string representation of a cell, coloring given cells green.
	cell := func(r, c int) string {
		ch := board[r][c]
		// Empty cells are shown as spaces
		if ch == '.' {
			return " "
		}
		// Given cells are colored green
		if given[r][c] {
			return fmt.Sprintf("%s%c%s", green, ch, reset)
		}
		return string(ch)
	}
	// Print the board with box separators
	for r := 0; r < 9; r++ {
		if r%3 == 0 {
			fmt.Println("o-------o-------o-------o")
		}

		fmt.Printf("| %s %s %s | %s %s %s | %s %s %s |\n",
			cell(r, 0), cell(r, 1), cell(r, 2),
			cell(r, 3), cell(r, 4), cell(r, 5),
			cell(r, 6), cell(r, 7), cell(r, 8),
		)
	}

	fmt.Println("o-------o-------o-------o")
}

//
// -------------------- ERRORS --------------------
//

func printInputError() {
	fmt.Println(`Error: Wrong input.

Please enter up to 9 Sudoku rows.
Each row must have exactly 9 characters.
Use numbers 1-9 for known cells and . for empty cells.
		
You can write rows normally:
534.7....
		
Or set a specific row:
1=534.7....
		
You may also combine both methods.`)
}

func printGridError() {
	fmt.Println("Error: This Sudoku board breaks the game rules.")
}

func printNoSolution() {
	fmt.Println("Error: This Sudoku puzzle cannot be solved.")
}
