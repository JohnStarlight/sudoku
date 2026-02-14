package main

import (
	"fmt"
	"os"
)

func main() {

	//os.Args[0] is the name of the program. The arguments have to be exactly 10 in order for the rows to be 9
	if len(os.Args) != 10 {

		fmt.Println("Error: wrong imput") // error message in case the user asks for either more or less rows than 9
		return                            // stopping the program in any error case, no need to scan the rest of the code if the imput is wrong

	}

	// we create a 9x9 board, which uses runes because the imputs are gonna be either numbers ('1'-'9') or dots ('.')
	var board [9][9]rune

	for i := 1; i <= 9; i++ { // Loop starts from 1 because os.Args[0] is the program's name.

		row := os.Args[i]

		if len(row) != 9 { // We check if each row has exactly 9 characters

			fmt.Println("Error: wrong imput")
			return

		}

		for index, r := range row {

			// "range" splits the string:
			// 'index' is the position of the character (0-8)
			// 'r' is the actual character (rune) at that position

			if (r < '1' || r > '9') && r != '.' {

				fmt.Println("Error: wrong imput") // only numbers 1-9 and . are allowed character imputs
				return

			}

			// Now we assign the character to the 9x9 board.
			// We use [i-1] as the Row position because the code's position starts from 0. so in a 9 row line, the positions are 0-8.
			// For Column position, we use index

			board[i-1][index] = r

		}

	}

}

/* Print example.

o- - -o- - -o- - -o
|1 2 3|4 5 6|7 8 9|
|4 5 6|7 8 9|1 2 3|
|7 8 9|1 2 3|4 5 6|
o- - -o- - -o- - -o
|9 1 2|3 4 5|6 7 8|
|3 4 5|6 7 8|9 1 2|
|6 7 8|9 1 2|3 4 5|
o- - -o- - -o- - -o
|8 9 1|2 3 4|5 6 7|
|2 3 4|5 6 7|8 9 1|
|5 6 7|8 9 1|2 3 4|
o- - -o- - -o- - -o
*/
