package greedy
/*
[
	[3,0,8,4],
	[2,4,5,7],
	[9,2,6,3],
	[0,3,1,0]



	[8, 4, 8, 7],
	[7, 4, 7, 7],
	[9, 4, 8, 7],
	[3, 3, 3, 3],

	if i want to find columns
	grid[0][0]
	grid[1][0]
]

*/

func initial(grid [][]int) ([]int, []int) {
	var localMaxRow []int
	for i := range grid {
		maxr := 0
		for _, val := range grid[i]{
			maxr = max(maxr, val)
		}
		localMaxRow = append(localMaxRow, maxr)
	}
	var localMaxCol []int
	for i := 0; i < len(grid); i++ {
		maxc := 0
		for j := 0; j < len(grid); j++ {
			maxc = max(maxc, grid[j][i])
		}
		localMaxCol = append(localMaxCol, maxc)
	}
	return localMaxRow, localMaxCol
}

func max(a, b int) int {
	if a > b{
		return a
	}
	return b
}

func min (a, b int ) int {
	if a < b {
		return a
	}
	return b
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	maxRow, maxCol := initial(grid)
	sum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0;  j < len(grid[i]); j++ {
			sum +=  (min(maxRow[i], maxCol[j]) - grid[i][j])
		}
	}
	return sum
}
