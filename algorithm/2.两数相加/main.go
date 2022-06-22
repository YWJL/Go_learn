package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	data1 := []int{9, 9, 9, 9, 9, 9, 9}
	data2 := []int{9, 9, 9, 9}
	headA := new(ListNode)
	headA.Val = data1[0]
	headB := new(ListNode)
	headB.Val = data2[0]
	t := headA
	// data2 := []int{5, 6, 4}
	for i := 1; i < len(data1); i++ {
		node := new(ListNode)
		node.Val = data1[i]
		t.Next = node
		t = node
	}
	t = headB
	for i := 1; i < len(data2); i++ {
		node := new(ListNode)
		node.Val = data2[i]
		t.Next = node
		t = node
	}
	t = addTwoNumbers(headA, headB)
	for t != nil {
		fmt.Print(t.Val)
		t = t.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := new(ListNode)
	p := l3
	flag := 0
	for l1 != nil && l2 != nil {
		node := new(ListNode)
		node.Val = (l1.Val + l2.Val + flag)
		if flag == 1 {
			flag = 0
		}
		if node.Val >= 10 {
			node.Val = node.Val - 10
			flag = 1
		}
		// fmt.Print(node.Val)
		p.Next = node
		p = node
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 == nil && l2 != nil {
		node := new(ListNode)
		node.Val = l2.Val + flag
		if flag == 1 {
			flag = 0
		}
		if node.Val >= 10 {
			node.Val = node.Val - 10
			flag = 1
		}
		p.Next = node
		p = node
		l2 = l2.Next
		// fmt.Print(node.Val)
	}
	for l2 == nil && l1 != nil {
		node := new(ListNode)
		node.Val = l1.Val + flag
		if flag == 1 {
			flag = 0
		}
		if node.Val >= 10 {
			node.Val = node.Val - 10
			flag = 1
		}
		p.Next = node
		p = node
		l1 = l1.Next
		// fmt.Print(node.Val)
	}
	if l1 == nil && l2 == nil && flag == 1 {
		node := new(ListNode)
		node.Val = 1
		p.Next = node
	}
	return l3.Next
	//这里.Next是符合力扣
}
