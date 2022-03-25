package sorting

import (
	"fmt"
	"sort"
)

func mymerge(nums *[]int, left, mid, right int) {
	temp := make([]int, right+1)
	copy(temp[left:], (*nums)[left:right+1])
	fmt.Println(*nums)
	fmt.Println(temp)
	mid = (right + left) / 2
	i, j, curr := left, mid+1, left
	for i <= mid && j <= right {
		fmt.Printf("i: %d, j: %d, curr: %d\n", temp[i], temp[j], curr)
		if temp[i] < temp[j] {
			(*nums)[curr] = temp[i]
			curr++
			i++
		} else {
			(*nums)[curr] = temp[j]
			curr++
			j++
		}
	}
	if j <= right-left {
		copy((*nums)[curr:right+1], temp[j:right+1])
	} else {
		copy((*nums)[curr:right+1], temp[i:mid+1])
	}
	fmt.Println(*nums)
}

var ref [][2]int
var res []int

func countSmaller(nums []int) []int {
	ref = make([][2]int, len(nums))
	for i, val := range nums {
		ref[i] = [2]int{i, val}
	}
	res = make([]int, len(nums))
	mergeSort(0, len(nums)-1)
	return res
}

func mergeSort(left, right int) {
	if left >= right {
		return
	}
	mid := left + (right-left)/2

	mergeSort(left, mid)
	mergeSort(mid+1, right)

	for i, j := left, mid+1; i <= mid; i++ {
		for j <= right && ref[j][1] < ref[i][1] {
			j++
		}
		res[ref[i][0]] += j - (mid + 1)
	}
	fmt.Printf("left: %d, right :%d\n", left, right)

	sort.Slice(ref[left:right+1], func(i, j int) bool {
		return ref[i][1] < ref[j][1]
	})

	fmt.Println(res)
	fmt.Println(ref)
}
