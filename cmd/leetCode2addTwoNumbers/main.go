package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := 0
	i := 1
	for l1 != nil {
		sum += l1.Val * i
		i *= 10
		l1 = l1.Next
	}
	i = 1
	for l2 != nil {
		sum += l2.Val * i
		i *= 10
		l2 = l2.Next
	}
	l3 := &ListNode{sum % 10, nil}
	lr := l3
	i = 10
	for sum >= i {
		l3.Next = &ListNode{sum % (i * 10) / i, nil}
		l3 = l3.Next
		i *= 10
	}
	return lr
}

func main() {
	n1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	n2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	result := addTwoNumbers(n1, n2)
	i := 1
	num := 0
	for result != nil {
		num += result.Val * i
		result = result.Next
		i *= 10
	}
	fmt.Println(num)

}
