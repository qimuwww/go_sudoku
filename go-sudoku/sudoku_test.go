// Author: wkj
package sudoku

import "fmt"

func ExampleSudoku() {
	var sudokuSlice = [][6][6]int{
		{
			{1, 0, 0, 2, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 4, 0, 5, 0},
			{0, 0, 5, 0, 4, 0},
			{3, 0, 0, 1, 0, 0},
		}, {
			{1, 4, 0, 0, 0, 0},
			{0, 0, 3, 0, 0, 0},
			{0, 0, 0, 3, 0, 0},
			{0, 5, 0, 0, 0, 1},
			{5, 0, 0, 1, 0, 6},
			{0, 2, 0, 0, 0, 3},
		},
	}

	for i := range sudokuSlice {
		fmt.Println("sudoku question:")
		PrintSudoku(&sudokuSlice[i])
		dst := Sudoku(sudokuSlice[i])
		fmt.Println("solved, one answer is:")

		PrintSudoku(&dst)
	}
	// Output:
	// sudoku question:
	// 1 0 0 | 2 0 0
	// 0 0 0 | 0 0 0
	// -------------
	// 0 0 0 | 0 0 0
	// 0 0 4 | 0 5 0
	// -------------
	// 0 0 5 | 0 4 0
	// 3 0 0 | 1 0 0
	// solved, one answer is:
	// 1 5 6 | 2 3 4
	// 4 2 3 | 5 1 6
	// -------------
	// 5 6 1 | 4 2 3
	// 2 3 4 | 6 5 1
	// -------------
	// 6 1 5 | 3 4 2
	// 3 4 2 | 1 6 5
	// sudoku question:
	// 1 4 0 | 0 0 0
	// 0 0 3 | 0 0 0
	// -------------
	// 0 0 0 | 3 0 0
	// 0 5 0 | 0 0 1
	// -------------
	// 5 0 0 | 1 0 6
	// 0 2 0 | 0 0 3
	// solved, one answer is:
	// 1 4 5 | 6 3 2
	// 2 6 3 | 5 1 4
	// -------------
	// 4 1 2 | 3 6 5
	// 3 5 6 | 2 4 1
	// -------------
	// 5 3 4 | 1 2 6
	// 6 2 1 | 4 5 3
}
