package main

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n <= 0 {
		return head
	}
	first := head
	i := 0

	for ; i < n; i++ {
		if first == nil {
			return head
		}
		first = first.Next
	}
	fmt.Println(first)
	if first == nil {
		return head.Next
	}
	second := head
	for first.Next != nil {
		second = second.Next
		first = first.Next
	}
	second.Next = second.Next.Next
	return head
}

func main() {
	head := &ListNode{
		Val: 1,
	}
	cur := head
	for i := 2; i <= 2; i++ {
		node := &ListNode{
			Val: i,
		}
		cur.Next = node
		cur = cur.Next
	}
	fmt.Printf("%p %v %p %v \n", head, head, cur, cur)
	data := removeNthFromEnd(head, 42)
	for data != nil {
		fmt.Printf("after %p %v \n", data, data)
		data = data.Next
	}
}
