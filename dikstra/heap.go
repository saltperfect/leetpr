package dikstra

import (
	"container/heap"
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

var result []*Item

func visiualHeap() {
	pq := make(PriorityQueue, 0)
	pq = append(pq, &Item{0, 12, -1}, &Item{1, 17, -1}, &Item{2, 14, -1}, &Item{3, 20, -1}, &Item{4, 10, -1},
		&Item{5, 11, -1}, &Item{6, 13, -1})

	spew.Dump(pq)

	heap.Init(&pq)

	fmt.Println("after heapify")

	spew.Dump(pq)

	pq[4].priority = 9

	spew.Dump(pq)

	heap.Fix(&pq, 4)

	fmt.Println("after fix")

	spew.Dump(pq)

	pq.updatePriority(6, 13, 8)

	spew.Dump(pq)
}

func (pq *PriorityQueue) findElem(index int, val int, priority int) {
	priorityqueue := *pq
	item := priorityqueue[index]
	if item.value == val {
		result = make([]*Item, 0)
		item.index = index
		result = append(result, item)
	}
	left := 2*index + 1
	right := 2*index + 2
	if left < pq.Len() && item.priority < priority {
		fmt.Printf("checking in left at index: %d\n", left)
		pq.findElem(left, val, priority)
	}
	if right < pq.Len() && item.priority < priority {
		fmt.Printf("checking in right at index: %d\n", right)
		pq.findElem(right, val, priority)
	}
}

func (pq PriorityQueue) updatePriority(val, priority, updatePriority int) {
	pq.findElem(0, val, priority)
	if len(result) > 0 {
		item := result[0]
		pq[item.index].priority = updatePriority

		heap.Fix(&pq, item.index)
	}
}
