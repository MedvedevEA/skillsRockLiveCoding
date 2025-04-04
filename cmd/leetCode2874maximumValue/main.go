package main

import (
	"fmt"
)

var nums = []int{1, 10, 3, 4, 19}

func maximumTripletValue(nums []int) int64 {
	var maxValue int64 = 0
	iMax, zMax := nums[0], nums[1]
	for j := 1; j < len(nums)-1; j++ {
		if iMax < nums[j-1] {
			iMax = nums[j-1]
		}
		if zMax == nums[j] {
			zMax = nums[j+1]
			for z := j + 2; z < len(nums); z++ {
				if nums[z] > zMax {
					zMax = nums[z]
				}
			}
		}
		newValue := int64((iMax - nums[j]) * zMax)
		if newValue > maxValue {
			maxValue = newValue
		}
	}
	return maxValue
}
func main() {
	result := maximumTripletValue(nums)
	fmt.Println(result)

}
