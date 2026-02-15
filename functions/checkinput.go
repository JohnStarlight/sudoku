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

	used := [9]bool{}
	positional := 0

	for i := 1; i < len(args); i++ {
		arg := args[i]
		if strings.Contains(arg, "=") {
			parts := strings.SplitN(arg, "=", 2)
			if len(parts) != 2 || len(parts[0]) == 0 {
				return false
			}
			idx, err := strconv.Atoi(parts[0])
			if err != nil || idx < 1 || idx > 9 {
				return false
			}
			if used[idx-1] {
				// duplicate explicit row assignment
				return false
			}
			row := parts[1]
			if len(row) != 9 {
				return false
			}
			for _, r := range row {
				if (r < '1' || r > '9') && r != '.' {
					return false
				}
			}
			used[idx-1] = true
		} else {
			row := arg
			if len(row) != 9 {
				return false
			}
			for _, r := range row {
				if (r < '1' || r > '9') && r != '.' {
					return false
				}
			}
			positional++
		}
	}

	return true
}

/*
Previous implementation (commented out to preserve original user code):

func CheckInput(args []string) bool {

	//os.Args[0] is the name of the program. The length of arguments have to be exactly 10 in order for the rows to be 9
	if len(args) != 10 {

		return false

	}

	for i := 1; i <= 9; i++ { // Loop starts from 1 because os.Args[0] is the program's name.

		row := args[i]

		if len(row) != 9 { // We check if each row has exactly 9 characters

			return false

		}

		for _, r := range row {

			// only numbers 1-9 and . are allowed character inputs
			if (r < '1' || r > '9') && r != '.' {

				return false

			}

		}

	}
			return true
}
*/

/*
Previous implementation (commented out to preserve active implementation):

func CheckInput(args []string) bool {

	// Allow up to 9 row arguments (os.Args[0] + up to 9 rows).
	// Missing rows will be treated as empty (".........").
	if len(args) > 10 {
		return false
	}

	for i := 1; i <= 9; i++ { // Loop starts from 1 because os.Args[0] is the program's name.
		if i < len(args) {
			row := args[i]

			if len(row) != 9 { // We check if each provided row has exactly 9 characters
				return false
			}

			for _, r := range row {
				// only numbers 1-9 and . are allowed character inputs
				if (r < '1' || r > '9') && r != '.' {
					return false
				}
			}
		} else {
			// missing row — allowed, will be initialized as empty by main
			continue
		}
	}

	return true
}
*/
