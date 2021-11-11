package backtrack

import "fmt"

func wrapper() {

	grid := [8][8]bool{}
	Solve(grid, 0)
}

func Solve(grid [8][8]bool, row int) bool {
	if row == 8 {
		for _, v := range grid {
			fmt.Println(v)
		}
		return true
	}
	for col := 0; col < 8 && checkIfGoodChoice(grid, col, row); col++ {
		grid[row][col] = true
		if Solve(grid, row+1) {
			return true
		}
		grid[row][col] = false
	}
	return false
}

func checkIfGoodChoice(grid [8][8]bool, col, row int) bool {

	for temprow := 0; temprow < 8; temprow++ {
		if grid[temprow][col] {
			return false
		}
	}
	var diff int
	var tempcount int = 0
	var temprow int
	var tempcol int
	if row < col {
		diff = col - row
		temprow = tempcount
		tempcol = tempcount + diff
	} else {
		diff = row - col
		temprow = tempcount + diff
		tempcol = tempcount
	}
	for temprow < 8 && tempcol < 8 {
		if grid[temprow][tempcol] && tempcol != col && temprow != row {
			return false
		}
		temprow++
		tempcol++
	}
	return true
}
