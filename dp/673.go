package dp

import "fmt"

func findNumberOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	cnt := make([]int, len(nums))
	dp[0] = 1
	cnt[0] = 1
	globalmax := 0
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[i] == dp[j]+1 {
					cnt[i] += cnt[j]
				}
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
					cnt[i] = cnt[j]
				}
				
				globalmax = max(globalmax, dp[i])
			}
		}
	}
	fmt.Printf("len: %v, cnt: %v\n", dp, cnt)
	res := 0
	for i, val := range dp {
		if val == globalmax {
			res += cnt[i]
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
