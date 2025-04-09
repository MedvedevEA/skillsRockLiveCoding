package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var root = &TreeNode{5, &TreeNode{2, nil, nil}, &TreeNode{-5, nil, nil}}

func findFrequentTreeSum(root *TreeNode) []int {
	m := map[int]int{}
	addSum(root, m)
	fmt.Println(m)
	result := []int{0}
	maxCount := 0
	for key := range m {
		if m[key] == maxCount {
			result = append(result, key)
		}
		if m[key] > maxCount {
			maxCount = m[key]
			result[0] = key
			result = result[:1]
		}
	}
	return result
}
func addSum(root *TreeNode, m map[int]int) int {
	if root == nil {
		return 0
	}
	sum := root.Val + addSum(root.Left, m) + addSum(root.Right, m)
	m[sum]++
	return sum
}

func main() {
	result := findFrequentTreeSum(root)
	fmt.Println(result)

}
