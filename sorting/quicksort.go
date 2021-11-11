package sorting

func quicksort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	left, right := pivot(nums, nums[0])
	quicksort(nums[:left])
	quicksort(nums[right:])
	return nums
}

func pivot(nums []int, p int) (int, int) {
	left, right := 0, len(nums)
	for mid := 0 ; mid < right; {
		switch a := nums[mid]; {
		case a < p:
			nums[mid] = nums[left]
			nums[left] = a
			left++
			mid++
		case a == p:
			mid++
		default:
			nums[mid] = nums[right-1]
			nums[right-1] = a 
			right--
		}
	}
	return left, right
}

func insertionsort(nums []int) []int {
	// invariant nums[:j] is sorted 

	for j := 1; j < len(nums); j++ {
		key := nums[j]
		i := j - 1
		for i >= 0 &&  nums[i] > key {
			nums[i+1] = nums[i]
			i--
		}
		nums[i+1] = key
	}
	return nums
}