package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	lr := &ListNode{0, nil}
	l3 := lr
	m := 0
	for l1 != nil || l2 != nil || m > 0 {
		l1value := 0
		if l1 != nil {
			l1value = l1.Val
			l1 = l1.Next
		}
		l2value := 0
		if l2 != nil {
			l2value = l2.Val
			l2 = l2.Next
		}

		l3value := (l1value + l2value + m) % 10
		m = (l1value + l2value + m) / 10
		l3.Next = &ListNode{l3value, nil}
		l3 = l3.Next

	}
	return lr.Next
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
