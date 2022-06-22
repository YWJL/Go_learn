package main

import "fmt"

func main() {
	fmt.Println("hello")
	nums := []string{"1", "hel", "袁"}
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
}
