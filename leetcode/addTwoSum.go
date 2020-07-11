package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

var (
	arg1 string
	arg2 string
)

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	flag.StringVar(&arg1, "list1", "1", "list1")
	flag.StringVar(&arg2, "list2", "1", "list2")
	flag.Parse()
	array1 := strings.Split(arg1, ",")
	array2 := strings.Split(arg2, ",")

	head1 := new(ListNode)
	list1 := head1
	for _, v := range array1 {
		list1.Next = new(ListNode)
		list1 = list1.Next
		list1.Val, _ = strconv.Atoi(v)
	}
	head1 = head1.Next

	head2 := new(ListNode)
	list2 := head2
	for _, v := range array2 {
		list2.Next = new(ListNode)
		list2 = list2.Next
		list2.Val, _ = strconv.Atoi(v)
	}
	head2 = head2.Next

	/*
		sum := addTwoNumbers(head1, head2)
		for sum != nil {
			fmt.Printf("%v ", sum.Val)
			sum = sum.Next
		}
	*/
	fmt.Printf("\n")
	sumR := addTwoReverseNumbers(head1, head2)
	for sumR != nil {
		fmt.Printf("%v ", sumR.Val)
		sumR = sumR.Next
	}
	fmt.Printf("\n")

}

/*
	2->4->3
	5->6->4
	7->0->8
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := new(ListNode)
	carry := 0
	suml := result
	for l1 != nil || l2 != nil || carry != 0 {
		suml.Next = new(ListNode)
		suml = suml.Next
		l1val := 0
		l2val := 0
		if l1 != nil {
			l1val = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2val = l2.Val
			l2 = l2.Next
		}
		sum := l1val + l2val + carry
		suml.Val = sum % 10
		carry = sum / 10
	}

	return result.Next
}

/*
	4->4->3
	5->6->4
	1->0->0->7
	result stack   9 0 7
	carry stack  0 1 0
*/

type stack []int

func (s *stack) push(item int) {
	(*s) = append((*s), item)
}

func (s *stack) pop() int {
	item := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return item
}

func (s stack) depth() int {
	return len(s)
}

func addTwoReverseNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sumStack := make(stack, 0)
	carryStack := make(stack, 0)
	carry := 0
	for l1 != nil || l2 != nil {
		fmt.Printf("%v %v %v\n", l1, l2, carry)
		l1val := 0
		l2val := 0
		if l1 != nil {
			l1val = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2val = l2.Val
			l2 = l2.Next
		}
		sum := l1val + l2val
		carry = sum / 10
		sumStack.push(sum % 10)
		carryStack.push(carry)
	}
	sumList := new(ListNode)

	isFirst := true
	carry = 0
	fmt.Printf("%v\n", sumStack)
	fmt.Printf("%v\n", carryStack)
	for sumStack.depth() != 0 || carryStack.depth() != 0 || carry != 0 {
		fmt.Println(sumStack.depth(), carryStack.depth(), carry)
		curSum := 0
		curCarry := 0
		curNode := new(ListNode)
		if sumStack.depth() != 0 {
			curSum = sumStack.pop()
		}
		if !isFirst && carryStack.depth() != 0 {
			curCarry = carryStack.pop()
		}
		isFirst = false
		curVal := (curSum + curCarry + carry)
		carry = curVal / 10
		curNode.Val = curVal % 10
		curNode.Next = sumList.Next
		sumList.Next = curNode
	}
	return sumList.Next
}
