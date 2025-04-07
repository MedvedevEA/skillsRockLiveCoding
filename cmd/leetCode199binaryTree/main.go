package main

import "fmt"

var root = &TreeNode{1, &TreeNode{2, &TreeNode{4, &TreeNode{5, nil, nil}, nil}, nil}, &TreeNode{3, nil, nil}}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	return getRight([]*TreeNode{root})
}
func getRight(nodes []*TreeNode) []int {
	if len(nodes) == 0 {
		return []int{}
	}
	newNodes := []*TreeNode{}
	for i := range nodes {
		if nodes[i].Right != nil {
			newNodes = append(newNodes, nodes[i].Right)
		}
		if nodes[i].Left != nil {
			newNodes = append(newNodes, nodes[i].Left)
		}
	}
	return append([]int{nodes[0].Val}, getRight(newNodes)...)

}

func main() {
	result := rightSideView(root)
	fmt.Println(result)

}
