package main

import (
	"fmt"
)

var (
	nums []int = []int{3, 5, 2, 6}
	k    int   = 2
)

func mostCompetitive(nums []int, k int) []int {

	iMin := 0
	i := 1
	for i < len(nums)-k+1 {
		if nums[i] == nums[iMin] {
			for j := 0; j < k; j++ {
				if nums[i+j] < nums[iMin+j] {
					iMin = i
					break
				}
			}
		}
		if nums[i] < nums[iMin] {
			iMin = i
		}
		i++

	}
	return nums[iMin : iMin+k]

}

func main() {
	result := mostCompetitive(nums, k)
	fmt.Println(result)
}
