package main

import "fmt"

type MountainArray struct {
	arr []int
}

func (this *MountainArray) get(index int) int {
	return this.arr[index]
}
func (this *MountainArray) length() int {
	return len(this.arr)
}

var (
	target      = 0
	mountainArr = MountainArray{arr: []int{1, 2, 5, 1}}
)

func findInMountainArray(target int, mountainArr *MountainArray) int {
	left, right := 1, mountainArr.length()-2
	var center int
	for left != right {
		center := (left + right) / 2
		if mountainArr.get(center) < mountainArr.get(center+1) {
			left = center + 1
		} else {
			right = center
		}
	}
	peak := left

	left, right = 0, peak
	for left < right {
		center = (right + left) / 2
		if mountainArr.get(center) < target {
			left = center + 1
		} else {
			right = center
		}
	}
	if mountainArr.get(left) == target {
		return left
	}

	left, right = peak, mountainArr.length()-1
	for left < right {
		center = (right + left) / 2
		if mountainArr.get(center) > target {
			left = center + 1
		} else {
			right = center
		}
	}
	if mountainArr.get(left) == target {
		return left
	}

	return -1
}

func main() {
	result := findInMountainArray(target, &mountainArr)
	fmt.Println(result)
}
