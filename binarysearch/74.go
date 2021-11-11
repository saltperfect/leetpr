package binarysearch

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	// try to find the row where the target exist and then do normal binary search

	// binary search of the first column will give the nearest lowest
	lo := 0
	hi := len(matrix) - 1

	for lo < hi {
		mid := (hi + lo) / 2 + 1
		if matrix[mid][0] == target {
			return true
		}
		if matrix[mid][0] < target {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	fmt.Printf("target row: %d\n", lo)
	targetnum := matrix[lo]

	lo = 0
	hi = len(targetnum) - 1

	for lo <= hi {
		mid := (lo + hi) / 2
		if targetnum[mid] == target {
			return true
		}
		if targetnum[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return false

}
