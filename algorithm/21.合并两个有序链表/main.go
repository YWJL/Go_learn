package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// fmt.Print("hello")
	data1 := []int{1, 2, 4}
	data2 := []int{1, 3, 4}
	list1 := new(ListNode)
	p1 := list1
	list2 := new(ListNode)
	p2 := list2
	for i := 0; i < len(data1); i++ {
		tmp := new(ListNode)
		tmp.Val = data1[i]
		p1.Next = tmp
		p1 = tmp
	}
	for i := 0; i < len(data2); i++ {
		tmp := new(ListNode)
		tmp.Val = data2[i]
		p2.Next = tmp
		p2 = tmp
	}
	list3 := mergeTwoLists(list1.Next, list2.Next)
	for list3 != nil {
		fmt.Print(list3.Val)
		list3 = list3.Next
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := new(ListNode)
	node := head
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			tmp := new(ListNode)
			tmp.Val = list1.Val
			node.Next = tmp
			node = tmp
			list1 = list1.Next
		} else if list1.Val > list2.Val {
			tmp := new(ListNode)
			tmp.Val = list2.Val
			node.Next = tmp
			node = tmp
			list2 = list2.Next
		}
	}
	for list1 != nil {
		tmp := new(ListNode)
		tmp.Val = list1.Val
		node.Next = tmp
		node = tmp
		list1 = list1.Next
	}
	for list2 != nil {
		tmp := new(ListNode)
		tmp.Val = list2.Val
		node.Next = tmp
		node = tmp
		list2 = list2.Next
	}
	return head.Next
}
