package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func countZero(num []int) int {
	start := 0
	end := len(num)
	for start < end {
		mid := (start + end) / 2
		if num[mid] >= 0 {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return len(num) - end
}

func countNegatives(grid [][]int) int {
	var res int
	for i, row := range grid {
		fmt.Printf(" %d: %d \n", i, countZero(row))
		res += countZero(row)
	}
	return res
}

func main2() {
	f, err := os.OpenFile("notes.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	grid := GetGrid(f)

	n := countNegatives(grid)
	fmt.Println(n)
}

func GetGrid(r io.Reader) [][]int {
	var grid [][]int
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			break
		}
		var row []int
		slicOfNum := strings.Split(scanner.Text(), " ")
		for _, val := range slicOfNum {
			v, _ := strconv.Atoi(val)
			row = append(row, v)
		}
		grid = append(grid, row)
		//scanner.Text() // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	//fmt.Print(grid)
	return grid
}
