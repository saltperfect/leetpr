package main

import (
	"fmt"
	"os"
)

func getSoldiers(nums []int) int {
	start, end := 0, len(nums)
	for start < end {
		mid := (start + end) / 2
		if nums[mid] > 0 {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return start
}

// func kWeakestRows(mat [][]int, k int) []int {
// 	dat := make([][2]int, len(mat))

// 	for i, v := range mat {
// 		dat[i] = [2]int{getSoldiers(v), i}
// 	}
// 	sort.Slice(dat, func(i, j int) bool {
// 		if dat[i][0] < dat[j][0] {
// 			return true
// 		} else if dat[i][0] == dat[j][0] {
// 			return dat[i][1] < dat[j][1]
// 		} else {
// 			return false
// 		}
// 	})
// 	res := make([]int, k)
// 	for i:=0; i< k; i++ {
// 		res[i] = dat[i][1]
// 	}
// 	return res
// }
func insert(num [][2]int, k [2]int) [][2]int {
	start, end := 0, len(num)
	for start < end {
		mid := (start + end) / 2
		if num[mid][0] < k[0] {
			start = mid + 1
		} else if num[mid][0] == k[0] {
			if num[mid][1] < k[1] {
				start = mid + 1
			} else {
				end = mid
			}
		} else {
			end = mid
		}
	}
	fmt.Printf("target point:%d\n", start)
	return append(num[:start], append([][2]int{k}, num[start:]...)...)
}

func kWeakestRows(mat [][]int, k int) []int {
	var dat [][2]int

	for i, v := range mat {
		dat = insert(dat, [2]int{getSoldiers(v), i})
	}

	var res []int
	for i := 0; i < k; i++ {
		res = append(res, dat[i][1])
	}
	return res
}

func main3() {
	f, err := os.OpenFile("notes.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	grid := GetGrid(f)
	var k int
	fmt.Scan(&k)
	n := kWeakestRows(grid, k)
	fmt.Println(n)
}
