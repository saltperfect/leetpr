package spanningtree

import (
	"fmt"
	"math"
	"container/heap"
)

type minheap [][2]int

func (mh minheap) Len() int {
	return len(mh)
}
func (mh minheap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}
func (mh minheap) Less (i, j int)bool {
	return mh[i][1] < mh[j][1]
}
func (mh *minheap) Push(a interface{}) {
	val := a.([2]int)
	*mh = append(*mh, val)
}
func (mh *minheap) Pop() interface{} {
	old := *mh
	l := len(old)
	val := old[l-1]
	*mh = old[:l-1]
	return val
}

func prim(grid [][][2]int) int {
	for i, val := range grid {
		fmt.Printf("%d -> %v\n", i, val)
	}

	pq := make(minheap, len(grid)) 
	for i := range pq {
		pq[i] = [2]int{i, 1000}
	}
	pq[0] = [2]int{0,0}
	heap.Init(&pq)

	dist := make(map[int]int, 0)
	parent := make(map[int]int, 0)
	parent[0] = math.MaxInt

	for len(pq) > 0 {
		fmt.Printf("dist: %v\n", dist)
		least := heap.Pop(&pq).([2]int)
		if _, ok := dist[least[0]]; !ok {
			dist[least[0]]= least[1]
			for _, neighbours := range grid[least[0]] {
				if _, ok := dist[neighbours[0]]; !ok {
					parent[neighbours[0]] = least[0]
				}
				heap.Push(&pq, neighbours)
				heap.Fix(&pq, len(pq)-1)
			}
		}
	}

	fmt.Printf("parent: %v\n", parent)

	return 0
}