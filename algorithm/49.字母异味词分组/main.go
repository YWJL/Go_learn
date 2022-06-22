//   114/117 未能全部通过
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("hello")
	strs := []string{"bdddddddddd", "bbbbbbbbbbc"}
	// strs := []string{"acc", "acc"}
	fmt.Println(groupAnagrams(strs))
}

// 输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
// 输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

func groupAnagrams(strs []string) [][]string {
	mapDate := make(map[string][]string)
	originMap := make(map[string]int)
	result := [][]string{}
	key := ""
	for i := 0; i < 26; i++ {
		originMap[string(rune(97+i))] = 0
	}
	for i := 0; i < len(strs); i++ {
		key = ""
		strElement := strs[i]
		for j := 0; j < len(strElement); j++ {
			fmt.Print(string(strElement[j]))
			originMap[string(strElement[j])] = originMap[string(strElement[j])] + 1
		}
		fmt.Println()
		for i := 0; i < 26; i++ {
			key = key + strconv.Itoa(originMap[string(rune(97+i))])
			originMap[string(rune(97+i))] = 0
		}
		if rowList, isOK := mapDate[key]; isOK {
			rowList = append(rowList, strElement)
			mapDate[key] = rowList
		} else {
			rowList = append(rowList, strElement)
			mapDate[key] = rowList
		}
	}
	for _, v := range mapDate {
		result = append(result, v)
	}
	fmt.Println(mapDate)
	return result
}

/*
func groupAnagrams(strs []string) [][]string {
	// var mapCompare func(map1 map[byte]int, map2 map[byte]int) bool
	mapCompare := func(map1 map[byte]int, map2 map[byte]int) bool {
		// fmt.Println(map1, map2)
		for k, v := range map1 {
			if map2[k] != v {
				return false
			}
		}
		for k, v := range map2 {
			if map1[k] != v {
				return false
			}
		}
		return true
	}
	result := [][]string{}
	mainMap := make(map[int]map[byte]int)
	// subMap := map[string]bool{"a": true}
	// mainMap[0] = map[byte]bool{0: false}
	count := 0
	for i := 0; i < len(strs); i++ {
		// fmt.Println("----------------------------------------", i)
		strElement := strs[i]
		subMap := make(map[byte]int)
		for j := 0; j < len(strElement); j++ {
			value, isOk := subMap[strElement[j]]
			if isOk {
				subMap[strElement[j]] = value + 1
			} else {
				subMap[strElement[j]] = 1
			}
		}
		flag := true
		for flag {
			for k, _ := range mainMap {
				if mapCompare(subMap, mainMap[k]) {
					row := result[k]
					row = append(row, strElement)
					result[k] = row
					flag = false
					break
				}
			}
			if flag {
				row := []string{}
				row = append(row, strElement)
				result = append(result, row)
				flag = false
				mainMap[count] = subMap
				count++
				flag = false
			}
		}
		// fmt.Println(subMap)
		// fmt.Println(mainMap)
		// fmt.Println(result)
	}
	// for i := 0; i < count; i++ {
	// 	listElement := []string{}
	// 	if
	// }
	return result
}
*/
