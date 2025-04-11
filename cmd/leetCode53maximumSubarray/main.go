package main

import "fmt"

var nums []int = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

func maxSubArray(nums []int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	l:=0
	r:=len(num)-1

	for l != r {
		if nums[l] <= 0 {
			l++
			continue
		}
		if num[r]<=0 {
			r--
			continue
		}
		for i:=i;j<len()
	}

}

func main() {
	result := maxSubArray()
	fmt.Println(result)

}
