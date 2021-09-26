package main

import (
	"fmt"
	"sort"
)

func main1() {
	nums := []int{-1,0,1,2,-1,-4}
	namo := threeSum(nums)
	fmt.Println(namo)
}


func threeSum(nums []int) [][]int {
    var tres [][]int 
    sort.Ints(nums)
    for i := 0; i < len(nums) ; i++ {
        target := -nums[i]
        front := i+1
        back := len(nums)
        for ; front < len(nums) && front < back; {
			sum := nums[front] + nums[back]
            switch {
                case sum < target:
                    front++
                case sum > target:
                    back--
                default:
                    res := []int{nums[i], nums[front], nums[back]}
                    tres = append( tres, res)
                    for ; nums[front] == nums[front+1] && front+1 < len(nums);  {
                        front++
                        res := []int{nums[i], nums[front], nums[back]}
                        tres = append( tres, res)
                    }
                    for ; nums[back] == nums[back-1] && back-1 <= i;  {
                        back++
                        res := []int{nums[i], nums[front], nums[back]}
                        tres = append( tres, res)
                    }
            }
        }
        for ; i+1 < len(nums) && nums[i] == nums[i+1]; {
            i++
            res := []int{nums[i], nums[front], nums[back]}
            tres = append( tres, res)
        }
    }
    return tres
}