package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
func swapPairs(head *ListNode) *ListNode {
	res := new(ListNode)
	res.Next = head
	p := res.Next
	fmt.Print("changeOrgin:")
	for p != nil {
		fmt.Print(p.Val)
		p = p.Next
	}
	fmt.Println("")
	cur := res
	if head == nil {
		return nil
	}
	next := head.Next
	if next == nil {
		return head
	}
	tmp := next.Next
	if tmp == nil {
		cur.Next = next
		next.Next = head
		head.Next = nil
		return res.Next
	}
	cur.Next = next
	next.Next = head
	head.Next = tmp
	p = res.Next
	fmt.Print("changeSwap1:")
	for p != nil {
		fmt.Print(p.Val)
		p = p.Next
	}
	fmt.Println("")
	cur = head
	head = tmp
	next = head.Next
	tmp = next.Next
	cur.Next = next
	next.Next = head
	head.Next = tmp
	p = res.Next
	fmt.Print("changeSwap2:")
	for p != nil {
		fmt.Print(p.Val)
		p = p.Next
	}
	fmt.Println("")
	cur = head
	head = tmp
	next = head.Next
	tmp = next.Next
	cur.Next = next
	next.Next = head
	head.Next = tmp
	p = res

	fmt.Print("changeSwap3:")
	for p != nil {
		fmt.Print(p.Val)
		p = p.Next
	}
	fmt.Println("")
	return res.Next
}
*/
func swapPairs(head *ListNode) *ListNode {
	res := new(ListNode)
	res.Next = head
	// p := res.Next
	// fmt.Print("changeOrgin:")
	// for p != nil {
	// 	fmt.Print(p.Val)
	// 	p = p.Next
	// }
	// fmt.Println("")
	cur := res
	for head != nil && head.Next != nil && head.Next.Next == nil {
		head = cur.Next
		next := head.Next
		cur.Next = next
		next.Next = head
		head.Next = nil
		cur = head
	}
	for head != nil && head.Next != nil && head.Next.Next != nil {
		head = cur.Next
		next := head.Next
		tmp := next.Next
		cur.Next = next
		next.Next = head
		head.Next = tmp
		cur = head
		// p = res.Next
		// fmt.Print("changeSwap1:")
		// for p != nil {
		// 	fmt.Print(p.Val)
		// 	p = p.Next
		// }
		// fmt.Println("")
	}
	return res.Next
}
func main() {
	data := []int{1, 2}
	head := new(ListNode)
	p := head
	for i := 0; i < len(data); i++ {
		node := new(ListNode)
		node.Val = data[i]
		p.Next = node
		p = node
	}
	// p = head.Next

	p = swapPairs(head.Next)
	for p != nil {
		fmt.Print(p.Val)
		p = p.Next
	}
}
