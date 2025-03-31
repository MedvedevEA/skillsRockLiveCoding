package main

import (
	"fmt"
	"sort"
)

var (
	nums = []int{1, 2, 3, 4, 5, 6}
	k    = 2
)

func smallestRangeII(nums []int, k int) int {
	sort.Ints(nums)
	result := nums[len(nums)-1] - nums[0]

	for i := 0; i < len(nums)-1; i++ {
		high := max(nums[len(nums)-1]-k, nums[i]+k)
		low := min(nums[0]+k, nums[i+1]-k)
		if high-low < result {
			result = high - low
		}
	}

	return result
}

func main() {
	result := smallestRangeII(nums, k)
	fmt.Println(result)

}
