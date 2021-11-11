package binarysearch

import (
	"fmt"
	"math"
	"os"
)

func smallestDivisor(nums []int, threshold int) int {
	left, right := 1, 0
	for _, n := range nums {
		right = int(math.Max(float64(right), float64(n)))
	}
	for left < right {
		mid := (left + right) / 2
		sum := 0
		for _, n := range nums {
			sum += (n + mid -1 )/ mid
		}
		if sum > threshold {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right
}

func main() {
	f, err := os.OpenFile("notes.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	grid := GetGrid(f)
	var k int
	fmt.Scan(&k)
	n := smallestDivisor(grid[0], k)
	fmt.Println(n)
}
