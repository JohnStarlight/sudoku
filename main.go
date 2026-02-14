func main() {

	// First, the code checks for correct input through the CheckInput function
	if CheckInput(os.Args) == false {
		fmt.Println("Error: Wrong Input") 
		return
	}

	//if input passes the check, we continue to make the board
	// we create a 9x9 board
	var board [9][9]rune

	for i := 1; i <= 9; i++ { // Loop starts from 1 because os.Args[0] is the program's name.

		row := os.Args[i]

		for index, r := range row { 
			// "range" splits the string:
			// 'index' is the position of the character (0-8)
			// 'r' is the actual character (rune) at that position

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
