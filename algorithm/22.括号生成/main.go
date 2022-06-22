package main

import (
	"fmt"
)

func main() {

	fmt.Println("hello" + "(")
	fmt.Println(generateParenthesis(3))

}

func generateParenthesis(n int) []string {
	result := []string{}
	left := 0
	right := 0
	strNow := ""
	var gobackfunc func(n int, left int, right int, strNow string)
	gobackfunc = func(n int, left int, right int, strNow string) {
		if right > left {
			return
		}
		if left == right && left == n {
			result = append(result, strNow)
			return
		}
		if left < n {
			gobackfunc(n, left+1, right, strNow+"(")
		}
		if left > right {
			gobackfunc(n, left, right+1, strNow+")")
		}
	}
	gobackfunc(n, left, right, strNow)
	return result
}

/*
func generateParenthesis(n int) []string {
	//不能把result作为参数传过去，因为回溯后中，result是空的，也就是你append的都会被清掉
	result := []string{}
	left := 0
	right := 0
	strNow := ""
	var gobackfunc func(result []string, n int, left int, right int, strNow string)
	gobackfunc = func(result []string, n int, left int, right int, strNow string) {
		if right > left {
			return
		}
		if left == right && left == n {
			result = append(result, strNow)
			return
		}
		if left < n {
			gobackfunc(result, n, left+1, right, strNow+"(")
		}
		if left > right {
			gobackfunc(result, n, left, right+1, strNow+")")
		}
	}
	gobackfunc(result, n, left, right, strNow)
	return result
}
*/

/*
func generateParenthesis(n int) []string {
	var generate func(left, right int, tmp []byte)

	res := make([]string, 0)
	generate = func(left, right int, tmp []byte) {
		if left == 0 && right == 0 {
			res = append(res, string(tmp))
			return
		}
		// 左括号还能放就优先放
		if left != 0 {
			tmp = append(tmp, '(')
			generate(left-1, right, tmp)
			tmp = tmp[:len(tmp)-1]
		}
		// 右括号数量不能大于左括号的，要不然合不回来，就不是合法数据了
		if right > left {
			tmp = append(tmp, ')')
			generate(left, right-1, tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}
	generate(n, n, []byte{})
	return res
}
*/
