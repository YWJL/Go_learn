package main

import "fmt"

func main() {
	fmt.Println(isValid2("(())}[]"))
	fmt.Println(isValid("()[]"))
}

func isValid(s string) bool {
	//  (:40    [:91   {:123       ):41     ]:93   }:125
	length := len(s)
	if length == 0 {
		return true
	}
	// if length%2 == 1 {
	// 	return false
	// }
	stock := []int{}
	for i := 0; i < length; i++ {
		tmp := s[i]
		if tmp == 40 || tmp == 123 || tmp == 91 {
			stock = append(stock, int(tmp))
		} else if tmp == 41 {
			if len(stock) == 0 {
				return false
			}
			if stock[len(stock)-1] == 40 {
				stock = stock[:len(stock)-1]
			} else {
				stock = append(stock, int(tmp))
			}
		} else if tmp == 93 {
			if len(stock) == 0 {
				return false
			}
			if stock[len(stock)-1] == 91 {
				stock = stock[:len(stock)-1]
			} else {
				stock = append(stock, int(tmp))
			}
		} else if tmp == 125 {
			if len(stock) == 0 {
				return false
			}
			if stock[len(stock)-1] == 123 {
				stock = stock[:len(stock)-1]
			} else {
				stock = append(stock, int(tmp))
			}
		}
	}
	return len(stock) == 0
}

func isValid2(s string) bool {
	//  (:40    [:91   {:123       ):41     ]:93   }:125
	length := len(s)
	if length == 0 {
		return true
	}
	mapdata := make(map[int]int)
	mapdata[40] = 41
	mapdata[91] = 93
	mapdata[123] = 125
	stock := []int{}
	for i := 0; i < length; i++ {
		tmp := s[i]
		if tmp == 40 || tmp == 123 || tmp == 91 {
			stock = append(stock, int(tmp))
		} else {
			if len(stock) == 0 {
				return false
			} else if mapdata[stock[len(stock)-1]] == int(tmp) {
				stock = stock[:len(stock)-1]
			} else {
				stock = append(stock, int(tmp))
			}
		}
	}
	return len(stock) == 0
}
