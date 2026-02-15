# Sudoku Solver (Go)

A Sudoku solver written in Go, using a **recursive backtracking** to find solutions for 9x9 grids. The project features a modular architecture, input validation, and a stylized terminal output.


## 🚀 Key Features
* **Backtracking Logic:** Efficiently solves even the most complex Sudoku puzzles by exploring possible paths and retracting upon dead ends.
* **Flexible Input Handling:** Supports both standard positional arguments (9 rows in order) and explicit row assignments (e.g., 1=53..7....).
* **Pre-Processing Validation:** Before solving, the program verifies:
    * Command-line argument count and formatting.
    * Character validity (only `1-9` and `.` are allowed).
    * Initial grid integrity (checking for rule violations in the starting numbers).
* **Enhanced UI:** Prints the solved board using ASCII borders and highlights the user-provided numbers in **Green** for better readability.


## 📁 Project Structure
The project follows a modular "separation of concerns" approach:
```
.
├── main.go                # Entry point: Coordinates parsing & execution flow
├── go.mod                 # Go module definition
├── functions/             # Logic & Helper Functions
│   ├── checkinput.go      # Validates command-line argument syntax
│   ├── isvalidgrid.go     # Checks initial board for Sudoku rule violations
│   ├── canplace.go        # Utility: Validates if a digit fits in a specific cell
│   └── solver.go          # Core Algorithm: Recursive Backtracking Solver
└── README.md              # Documentation
```

## 🛠️ Usage
To run the program, use the go run command followed by the 9 rows of the Sudoku grid. Use a dot (.) for empty cells.

**Standard Positional Input:**
go run . "53..7...." "6..195..." ".98....6." "8...6...3" "4..8.3..1" "7...2...6" ".6....28." "...419..5" "....8..79"

**Row Assignment Input:**
go run . 1=53..7.... 5=4..8.3..1 (etc.)


## ⚠️ Error Handling
The program will display `Error` if:
* The number of row arguments provided is incorrect.
* Any row does not contain exactly 9 characters.
* The input contains invalid characters.
* The starting board is inherently unsolvable or violates Sudoku rules.

## 📝 Authors
* Dimitrios Mitsios Antonakos #dmitsios
* Alexandros Mylonas #almylonas
* Ioannis Vogiakelis #ivogiake