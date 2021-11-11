package dfs

import "fmt"

var parent map[int]int
var result []int

func dfst(grid [][]int){
	
	parent = make(map[int]int, 0)

	for i := 0; i < len(grid); i++ {
		if _, ok := parent[i]; !ok {
			dfsvisitt(grid, i)
		}
	}
	fmt.Printf("result: %v\n", result)
}

func dfsvisitt(grid [][]int, i int) {

	for _, val := range grid[i] {
		if _, ok := parent[val]; !ok {
			parent[val] = i
			dfsvisitt(grid, val)
		}
	}
	result = append(result, i)

}