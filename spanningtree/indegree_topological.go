package spanningtree

import "fmt"

func tropicalsort(grid [][]int) int {
	indegree := make([]int, len(grid))
	for _, val := range grid {
		for _, elem := range val {
			indegree[elem]++
		}
	}
	q := make([]int, 0)
	for i := range indegree {
		if indegree[i] == 0 {
			q = append(q, i)
		}
	}
	final := make([]int, 0)
	for len(q) > 0 {
		first := q[0]
		final = append(final, first)
		q = q[1:]
		for _, neighbours := range grid[first] {
			indegree[neighbours]--
			if indegree[neighbours] == 0 {
				q = append(q, neighbours)
			}
		} 
	}
	fmt.Printf("final: %v\n", final)
	return 1
}