package dikstra

import (
	"fmt"
	"math"
)

func bellman(grid [][][2]int, root int) {
	q := make([]int, 0)
	dist := make([]int, len(grid))
	inq := make([]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		if i == root {
			dist[i] = 0
			continue
		}
		dist[i] = math.MaxInt
	}
	inq[root] = true
	q = append(q, root)
	for len(q) > 0 {
		l := len(q)
		for i := 0; i < l; i++ {
			parent := q[i]
			for _, val := range grid[parent] {
				if dist[val[0]] > dist[parent] + val[1] {
					dist[val[0]] = dist[parent] + val[1]
					if !inq[val[0]] {
						q = append(q, val[0])
					} 
				}
			}
		}
		q = q[l:]
	}
	fmt.Printf("dist: %v\n", dist)
}