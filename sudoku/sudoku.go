package main

// We need these packages:
// fmt - for printing output
// os - for reading command-line arguments
import (
	"fmt"
	"os"
)

// MAIN FUNCTION - Entry point of the program
func main() {
	// os.Args contains all command-line arguments
	// os.Args[0] is always the program name (e.g., "./sudoku" or "main")
	// os.Args[1:] gives us the actual arguments (the 9 Sudoku rows)
	arguments := os.Args[1:]

	// STEP 1: Check if we have exactly 9 arguments
	// Why? A valid Sudoku must have 9 rows
	if len(arguments) != 9 {
		// Print "Error" and exit the program
		// The \n adds a newline, and cat -e will show it as "$"
		fmt.Println("Error")
		return
	}

	// STEP 2: Validate the input format and Sudoku rules
	// If validation fails, print error and exit
	if !isValidInput(arguments) {
		fmt.Println("Error")
		return
	}

	// STEP 3: Convert the strings to a 2D grid of integers
	// We use integers because they're easier to work with than characters
	// -1 represents an empty cell (was a dot in the input)
	grid := convertToGrid(arguments)

	// STEP 4: Try to solve the Sudoku
	// If solving fails, print error
	if !solveSudoku(grid) {
		fmt.Println("Error")
		return
	}

	// STEP 5: If we get here, solving was successful!
	// Print the solved grid
	printGrid(grid)
}

// isValidInput checks if the input follows all Sudoku rules
func isValidInput(rows []string) bool {
	// First, check each individual row's format
	for i := 0; i < 9; i++ {
		// Each row must be exactly 9 characters long
		if len(rows[i]) != 9 {
			return false
		}

		// Check each character in the row
		for j := 0; j < 9; j++ {
			char := rows[i][j] // Get the character at position j

			// Valid characters are either digits 1-9 or a dot
			// We check if it's NOT a valid character
			if char != '.' && (char < '1' || char > '9') {
				return false
			}
		}
	}

	// Now check if the initial puzzle follows Sudoku rules
	// We need to check rows, columns, and 3x3 boxes

	// Convert to grid for easier checking
	grid := make([][]int, 9)
	for i := 0; i < 9; i++ {
		grid[i] = make([]int, 9)
		for j := 0; j < 9; j++ {
			if rows[i][j] == '.' {
				grid[i][j] = 0 // Use 0 for empty cells in validation
			} else {
				// Convert char to int: '1' becomes 1, '2' becomes 2, etc.
				grid[i][j] = int(rows[i][j] - '0')
			}
		}
	}

	// Check each row for duplicates
	for i := 0; i < 9; i++ {
		// Create a map to track which numbers we've seen
		// Why a map? It's perfect for checking duplicates
		seen := make(map[int]bool)
		for j := 0; j < 9; j++ {
			num := grid[i][j]
			if num != 0 { // If it's not an empty cell
				if seen[num] { // If we've seen this number before
					return false // Duplicate found!
				}
				seen[num] = true // Mark this number as seen
			}
		}
	}

	// Check each column for duplicates
	for j := 0; j < 9; j++ {
		seen := make(map[int]bool)
		for i := 0; i < 9; i++ {
			num := grid[i][j]
			if num != 0 {
				if seen[num] {
					return false
				}
				seen[num] = true
			}
		}
	}

	// Check each 3x3 box for duplicates
	// We have 9 boxes total (3 rows of boxes × 3 columns of boxes)
	for boxRow := 0; boxRow < 3; boxRow++ {
		for boxCol := 0; boxCol < 3; boxCol++ {
			seen := make(map[int]bool)

			// Each box contains 3 rows and 3 columns
			// Start row = boxRow * 3, end at boxRow*3 + 3
			// Start col = boxCol * 3, end at boxCol*3 + 3
			for i := boxRow * 3; i < boxRow*3+3; i++ {
				for j := boxCol * 3; j < boxCol*3+3; j++ {
					num := grid[i][j]
					if num != 0 {
						if seen[num] {
							return false
						}
						seen[num] = true
					}
				}
			}
		}
	}

	// If we passed all checks, the input is valid
	return true
}

// convertToGrid transforms string rows into a 2D integer grid
// -1 represents empty cells (easier to work with than '.' or 0)
func convertToGrid(rows []string) [][]int {
	// Create a 2D slice: 9 rows, each row will have 9 columns
	grid := make([][]int, 9)

	for i := 0; i < 9; i++ {
		// Create each row with 9 columns
		grid[i] = make([]int, 9)
		for j := 0; j < 9; j++ {
			if rows[i][j] == '.' {
				grid[i][j] = -1 // -1 means empty cell
			} else {
				// Convert character to integer
				// '1' becomes 1, '2' becomes 2, etc.
				grid[i][j] = int(rows[i][j] - '0')
			}
		}
	}
	return grid
}

// solveSudoku uses backtracking to find the solution
// Returns true if solved, false if impossible
func solveSudoku(grid [][]int) bool {
	// Find the next empty cell (with value -1)
	row, col := findEmptyCell(grid)

	// If no empty cell found, the puzzle is solved!
	if row == -1 && col == -1 {
		return true
	}

	// Try placing digits 1-9 in the empty cell
	for num := 1; num <= 9; num++ {
		// Check if placing this number is valid
		if isValidMove(grid, row, col, num) {
			// Place the number
			grid[row][col] = num

			// Recursively try to solve the rest of the puzzle
			if solveSudoku(grid) {
				return true // If it works, we're done!
			}

			// If placing num didn't lead to a solution,
			// backtrack by removing it and try the next number
			grid[row][col] = -1
		}
	}

	// If no number works, puzzle is unsolvable from this state
	return false
}

// findEmptyCell locates the first empty cell in the grid
// Returns (-1, -1) if no empty cells exist
func findEmptyCell(grid [][]int) (int, int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == -1 {
				return i, j // Return row and column of empty cell
			}
		}
	}
	return -1, -1 // No empty cells found
}

// isValidMove checks if placing a number at (row, col) follows Sudoku rules
func isValidMove(grid [][]int, row, col, num int) bool {
	// Check the row for duplicates
	for j := 0; j < 9; j++ {
		if grid[row][j] == num {
			return false // Number already exists in this row
		}
	}

	// Check the column for duplicates
	for i := 0; i < 9; i++ {
		if grid[i][col] == num {
			return false // Number already exists in this column
		}
	}

	// Check the 3x3 box for duplicates
	// Find which box we're in
	boxRow := row / 3 // Integer division: rows 0-2 give 0, 3-5 give 1, 6-8 give 2
	boxCol := col / 3 // Same for columns

	// Check all cells in this 3x3 box
	for i := boxRow * 3; i < boxRow*3+3; i++ {
		for j := boxCol * 3; j < boxCol*3+3; j++ {
			if grid[i][j] == num {
				return false // Number already exists in this box
			}
		}
	}

	// If we passed all checks, the move is valid
	return true
}

// printGrid displays the solved Sudoku in the required format
func printGrid(grid [][]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// Print each number
			fmt.Print(grid[i][j])

			// Add a space after each number except the last in the row
			if j < 8 {
				fmt.Print(" ")
			}
		}
		// After each row, print a newline
		// cat -e will show this as "$" at the end of each line
		fmt.Println()
	}
}
