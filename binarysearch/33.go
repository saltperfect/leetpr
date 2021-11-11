package binarysearch

func search(nums []int, target int) int {
    lo := 0
    high := len(nums) - 1
    for lo < high {
        mid := (lo + high)/2
        if nums[mid] > nums[high] {
            lo = mid + 1
        }else{
            high = mid
        }
    }
    rot := lo
    lo, high = 0, len(nums) -1
    for lo <= high {
        mid := (lo+high)/2
        realmid := (mid+rot)%len(nums)
        if nums[realmid] == target {
            return realmid
        }
        if nums[realmid] < target {
            lo = mid + 1
        }
        if nums[realmid] > target {
            high = mid-1
        }
    }
    return -1
}