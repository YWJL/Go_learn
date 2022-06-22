package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}

func permute(nums []int) [][]int {
	result := [][]int{}
	length := len(nums)
	dataMap := make(map[int]bool)
	for i := 0; i < length; i++ {
		dataMap[nums[i]] = false
	}
	row := []int{}
	var backMethod func(nums []int, length int, row []int)
	backMethod = func(nums []int, length int, row []int) {
		for i := 0; i < length; i++ {
			// if dataMap[nums[i]] {
			// 	return //这里必须得注视掉，因为递归的时候（row[:low(row)-1]），会直接从这个出口出去了。
			// }
			if !dataMap[nums[i]] {
				row = append(row, nums[i])
				dataMap[nums[i]] = true
				backMethod(nums, length, row)
				row = row[:len(row)-1]
				dataMap[nums[i]] = false
			}
		}
		if length == len(row) {
			row2 := []int{}
			for i := 0; i < len(row); i++ {
				row2 = append(row2, row[i])
			}
			result = append(result, row2)
			// return
		}
	}
	backMethod(nums, length, row)
	return result
}
