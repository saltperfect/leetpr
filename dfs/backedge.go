package dfs

import "fmt"

var visiting map[int]bool
var backedge map[int]int

func dfs(grid [][]int) {
	
	parent := make(map[int]int, 0)
	visiting = make(map[int]bool, 0)
	backedge =  make( map[int]int, 0 )

	for i := 0; i < len(grid); i++ {
		if _, ok := parent[i]; !ok {
			dfsvisit(grid, i, parent)
		}
	}
	fmt.Printf("parent map: %v\n", parent )
	fmt.Printf("visiting map: %v\n", visiting )
	fmt.Printf("backedge map: %v\n", backedge )

}

func dfsvisit(grid [][]int, i int, parent map[int]int) {
	visiting[i] = true
	for _, val := range grid[i] {
		if _, ok := parent[val]; !ok {
			parent[val] = i
			dfsvisit(grid, val, parent)
		}
		if visiting[val]{
			backedge[i] = val
		}
	}
	visiting[i] = false
}