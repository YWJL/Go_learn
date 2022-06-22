//双指针法
package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")
	nums := []int{3}
	fmt.Println(removeElement(nums, 3))
}

func removeElement(nums []int, val int) int {
	l := 0
	r := len(nums) - 1
	if r == -1 {
		return 0
	}
	for l != r {
		for nums[l] != val {
			if l == r {
				break
			}
			l++
		}
		for nums[r] == val {
			if l == r {
				break
			}
			r--
		}
		//这个时候，nums[l] =val && nums[r]!=val
		tmp := nums[l]
		nums[l] = nums[r]
		nums[r] = tmp
		fmt.Println(nums)
	}
	if nums[l] == val { //函数走到这里，l=r
		return l
	} else {
		return l + 1
	}
}
