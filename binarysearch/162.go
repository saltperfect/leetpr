package binarysearch

func findPeakElement(nums []int) int {
    lo := 0
    hi := len(nums) - 1
    for lo < hi {
        mid := (lo + hi)/2
        if mid + 1 <= hi && nums[mid] < nums[mid+1] {
            lo = mid + 1
        } else if mid - 1 >= lo && nums[mid] > nums[mid+1] {
            hi = mid - 1
        } else {
			return mid
		}
    }
	return -1
}