package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

var head = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}

func reorderList(head *ListNode) {
	h := head
	s := []*ListNode{}
	for h != nil {
		s = append(s, h)
		h = h.Next
	}
	if len(s) < 3 {
		return
	}
	for i := 0; i < len(s)/2; i++ {
		s[i].Next = s[len(s)-1-i]
		s[len(s)-1-i].Next = s[i+1]
	}
	s[len(s)/2].Next = nil
}
func main() {
	reorderList(head)
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}

}

/*
	l1 := head
	l2 := head
	l3 := head
	for l3 != nil {
		l2 = l2.Next
		l3 = l3.Next
		if l3 != nil {
			l3 = l3.Next
		}
	}
	for l2 != nil {
		l3 = l1.Next
		l1.Next, l2 = l2, l2.Next
		l1.Next.Next, l1 = l3, l3
	}

*/
