package functions

import (
	"strconv"
	"strings"
)

func CheckInput(args []string) bool {
	// Allow program name + up to 9 row specifications. Each row-spec may be either:
	// - a 9-char string with digits/dots (positional), or
	// - "N=XXXXXXXXX" where N is 1..9 and the RHS is a 9-char row for that exact row index.
	if len(args) > 10 {
		return false
	}
	// Track which rows have been assigned (either positionally or explicitly).
	used := [9]bool{}
	positional := 0
	// Process each argument after the program name.
	for i := 1; i < len(args); i++ {
		arg := args[i]
		// Check if it's an explicit row assignment (contains '=').
		if strings.Contains(arg, "=") {
			parts := strings.SplitN(arg, "=", 2)
			if len(parts) != 2 || len(parts[0]) == 0 {
				return false
			}
			// Validate the row index (1-9).
			idx, err := strconv.Atoi(parts[0])
			if err != nil || idx < 1 || idx > 9 {
				return false
			}
			// Check for duplicate row assignment.
			if used[idx-1] {
				// duplicate explicit row assignment
				return false
			}
			// Validate the row content.
			row := parts[1]
			if len(row) != 9 {
				return false
			}
			// Validate characters in the row.
			for _, r := range row {
				if (r < '1' || r > '9') && r != '.' {
					return false
				}
			}
			// Mark this row index as used.
			used[idx-1] = true
		} else {
			row := arg
			if len(row) != 9 {
				return false
			}
			// Check if we've already assigned this row index via an explicit assignment.
			for _, r := range row {
				if (r < '1' || r > '9') && r != '.' {
					return false
				}
			}
			//
			positional++
		}
	}

	return true
}
