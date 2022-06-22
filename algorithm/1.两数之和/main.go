package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 3}
	fmt.Print(twoSum(nums, 6))
	fmt.Print(twoSumMap(nums, 6))
}

func twoSum(nums []int, target int) []int {
	answer := make([]int, 2)
	var i, j int
	for i = 0; i < len(nums); i++ {
		for j = i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				answer[0] = i
				answer[1] = j
				return answer
			}
		}
	}
	// answer[0] = 0
	// answer[1] = 0
	return answer
}

func twoSumMap(nums []int, target int) []int {
	datamap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		datamap[nums[i]] = i
	}
	for index1 := 0; index1 < len(nums); index1++ {
		index2, isOk := datamap[target-nums[index1]]
		if isOk && index2 != index1 {
			return []int{index1, index2}
		}
	}
	return nil
}
