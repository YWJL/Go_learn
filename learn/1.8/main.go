package main

import "fmt"

func main() {
	fmt.Print(Signum(1))
}

func Signum(x int) int {
	switch {
	case x > 0:
		return +1
	case x < 0:
		return -1
	default:
		return 0
	}
}
