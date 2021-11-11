package binarysearch

import (
	"fmt"
	"math"
	"os"
)

func shipWithinDays(weights []int, days int) int {
	var left, right int
	for _, w := range weights {
		left = int(math.Max(float64(left), float64(w)))
		right += w
	}
	for left < right {
		mid := (left + right) / 2
		cur, need := 0, 1
		for _, w := range weights {
			if cur+w > mid {
				need += 1
				cur = 0
			}
			cur += w
		}

		if need > days {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func main4() {
	f, err := os.OpenFile("notes.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	grid := GetGrid(f)
	var k int
	fmt.Scan(&k)
	n := shipWithinDays(grid[0], k)
	fmt.Println(n)
}
