// 二分查找法
package main

import "fmt"

func main() {
	fmt.Println("hello")
	nums := []int{-5, -4, -3, -2, -1, 0, 1, 2, 4, 5}
	fmt.Println(searchInsert(nums, 3))
}

func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	m := (l + r) / 2
	flag := true
	for flag {
		//这里顺序要考究下
		if nums[r] < target {
			return r + 1
		}
		if nums[r] == target {
			return r
		}
		if nums[l] >= target {
			return l
		}
		//以上先考虑边界情况
		if nums[m] < target && target < nums[m+1] {
			return m + 1
		}
		if nums[m] == target {
			return m
		}
		if nums[m] < target {
			l = m
			m = (l + r) / 2
		}
		if nums[m] > target {
			r = m
			m = (l + r) / 2
		}
	}
	return m
}
