package dikstra

import (
	"container/heap"
	"fmt"
	"math"
)

type vertex struct {
	point int
	distance int
}

type Item struct {
	value int
	priority int
	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(i interface{}) {
	n := len(*pq)
	item := i.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface {} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[:n-1]
	return item
}



func dikstra(graph [][]vertex, start int) {

	// numVertex := len(graph)

	var pq = make(PriorityQueue, 0)

	// results := make(map[int]int, 0)

	tempDist := make([]int, len(graph))
	for i := range tempDist {
		if i == start {
			tempDist[i] = 0
		}
		tempDist[i] = math.MaxInt
	}
	
	pq.Push(&Item{
		value: start,
		priority: 0,
	}) // marking 0 distance
	heap.Init(&pq)

	for len(pq) > 0 {
		item := heap.Pop(&pq)
		parent := item.(*Item)
		// results[parent.value] = parent.priority
		for _, child := range graph[parent.value] {
			// if _, ok := results[child.point]; !ok {
				if tempDist[child.point] > parent.priority + child.distance {
					tempDist[child.point] = parent.priority + child.distance
					heap.Push(&pq, &Item{child.point, tempDist[child.point], -1 })
				}
			// }
		}
	}

	fmt.Printf("resuts: %v\n", tempDist)

}