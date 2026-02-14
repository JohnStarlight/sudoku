package functions

func CheckInput(args []string) bool {

//os.Args[0] is the name of the program. The length of arguments have to be exactly 10 in order for the rows to be 9
	if len(agrs) != 10 {

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

