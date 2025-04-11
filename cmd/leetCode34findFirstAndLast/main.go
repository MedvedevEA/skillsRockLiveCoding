package main

import "fmt"

var (
	nums   []int = []int{5, 7, 7, 8, 8, 10}
	target int   = 8
)

func searchRange(nums []int, target int) []int {
	a := nums
	aLeft := 0
	for len(a) > 0 {
		c := len(a) / 2
		if a[c] < target {
			a = a[c+1:]
			aLeft += c + 1
		} else if a[c] == target && (c-1 < 0 || a[c-1] < target) {
			aLeft += c
			break
		} else {
			a = a[:c]
		}
	}
	if len(a) == 0 {
		return []int{-1, -1}
	}
	b := nums
	bLeft := 0
	for len(b) > 0 {
		c := len(b) / 2
		if b[c] > target {
			b = b[:c]
		} else if b[c] == target && (c+1 > len(b)-1 || b[c+1] > target) {
			bLeft += c
			break
		} else {
			b = b[c+1:]
			bLeft += c + 1

		}
	}
	return []int{aLeft, bLeft}
}

func main() {
	result := searchRange(nums, target)

	fmt.Println(result)
}
