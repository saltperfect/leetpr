package sorting

import "fmt"

func mergesort(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}
	mid := len(nums) / 2
	first := nums[:mid]
	second := nums[mid:]
	fmt.Printf("first: %v, second: %v\n", first, second)
	firstsorted := mergesort(first)
	secondsorted := mergesort(second)
	return merge(firstsorted, secondsorted)
}

func merge(a []int, b []int) []int {
	res := []int{}
	ap := 0
	bp := 0
	for ap < len(a) && bp < len(b) {
		if a[ap] < b[bp] {
			res = append(res, a[ap])
			ap++
		} else {
			res = append(res, b[bp])
			bp++
		}
	}
	for ap == len(a) && bp < len(b){
		res = append(res, b[bp])
		bp++
	}
	for bp == len(b) && ap < len(a) {
		res = append(res, a[ap])
		ap++
	}
	return res
}
