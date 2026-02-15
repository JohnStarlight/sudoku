# Sudoku Solver (Go)

A Sudoku solver written in Go, using a **recursive backtracking algorithm** to solve 9×9 Sudoku grids. The project features a modular architecture, input validation, and a stylized terminal output.


## 🚀 Key Features
* **Backtracking Logic:** Efficiently solves even the most complex Sudoku puzzles by exploring possible paths and backtracking when dead ends are reached.
* **Flexible Input Handling:** Supports both standard positional rows (up to 9 rows) and explicit row assignments (e.g., 1=53..7....). Missing rows are treated as empty.
* **Pre-Processing Validation:** Before solving, the program verifies:
    * Command-line argument count and formatting.
    * Character validity (only `1-9` and `.` are allowed).
    * Initial grid integrity (checking for rule violations in the starting numbers).
* **Enhanced UI:** Prints the solved board using ASCII borders and highlights the user-provided numbers in **Green** for better readability.

## 🧠 Solver Strategy
The solver uses a recursive backtracking algorithm to explore possible values.
If a conflict occurs, the algorithm backtracks and tries alternative values until
a valid solution is found or all possibilities are exhausted.


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
Run the program with up to 9 Sudoku rows.
Use a dot (.) for empty cells.

**Standard positional input:**
go run . "53..7...." "6..195..." ".98....6."

**Row assignment input:**
go run . 1=53..7.... 5=4..8.3..1

**Mixed input (allowed):**
go run . "53..7...." 5=4..8.3..1 ".98....6."


## ⚠️ Error Handling
The program displays clear error messages in the following cases:

* More than 9 rows are provided.
* Any row does not contain exactly 9 characters.
* Invalid characters are used (only numbers 1–9 and . are allowed).
* A row number assignment is invalid or duplicated.
* The starting grid violates Sudoku rules (duplicate numbers in a row, column, or 3×3 box).
* The puzzle has no possible solution.

## 📜 License
This project is licensed under the MIT License.

## 📝 Authors
* Dimitrios Mitsios Antonakos #dmitsios
* Alexandros Mylonas #almylonas
* Ioannis Vogiakelis #ivogiake