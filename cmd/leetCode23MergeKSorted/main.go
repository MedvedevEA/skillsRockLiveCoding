package main

import "fmt"

var l1 = &ListNode{1, &ListNode{4, &ListNode{5, nil}}}
var l2 = &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
var l3 = &ListNode{2, &ListNode{6, nil}}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	for i := 1; i < len(lists); i++ {
		if lists[i] == nil {
			continue
		}
		if lists[0] == nil {
			lists[0] = lists[i]
			continue
		}
		if lists[i].Val < lists[0].Val {
			lists[0], lists[i] = lists[i], lists[0]
		}
		var l0pre *ListNode
		l0 := lists[0]
		li := lists[i]

		for l0 != nil && li != nil {
			if li.Val < l0.Val {
				l0pre.Next, l0pre, li, l0 = li, li, l0, li.Next

			} else {
				l0pre, l0 = l0, l0.Next
			}
		}
		if l0 == nil {
			l0pre.Next = li
		}

	}
	return lists[0]
}

func main() {
	result := mergeKLists([]*ListNode{l1, l2, l3})
	for result != nil {
		fmt.Printf("%d ", result.Val)
		result = result.Next
	}
	fmt.Printf("\n")

}
