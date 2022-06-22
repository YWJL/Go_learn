package main

import "fmt"

func main() {
	fmt.Println("hello")
	nums := []int{-2, 1}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	maxNum := nums[0]
	for i := 0; i < len(nums); i++ {
		now := nums[i]
		if now > maxNum {
			maxNum = now
		}
		for j := i + 1; j < len(nums); j++ {
			now = now + nums[j]
			if now > maxNum {
				maxNum = now
			}
		}
	}
	return maxNum
}
