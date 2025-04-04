package main

import "fmt"

var (
	nums = []int{0, 1, 2, 2, 3, 3, 3, 3}
	val  = 3
)

func removeElement(nums []int, val int) int {
	removeCount := 0
	index := 0
	for index < len(nums)-removeCount {
		if nums[index] != val {
			index++
			continue
		}
		nums[index] = nums[len(nums)-1-removeCount]
		removeCount++
	}
	return len(nums) - removeCount
}

func main() {
	result := removeElement(nums, val)
	fmt.Println(nums)
	fmt.Println(result)

}
