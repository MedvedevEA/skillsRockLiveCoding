package main

import "fmt"

var arr = []int{1, 2, 3, 2, 1, 4, 5, 4, 7}

func minMovesToDelete(arr []int) int {
	if len(arr) < 3 {
		return len(arr)
	}
	min := len(arr)
	for i := 0; i < len(arr); i++ {

		j := 0
		for i-j > -1 && i+j < len(arr) && arr[i-j] == arr[i+j] {
			arr2 := []int{}
			arr2 = append(arr2, arr[:i-j]...)
			arr2 = append(arr2, arr[i+j+1:]...)

			min2 := minMovesToDelete(arr2)
			if min2+1 < min {
				min = min2 + 1
			}
			j++

		}

	}
	return min

}

func main() {
	result := minMovesToDelete(arr)
	fmt.Println(result)

}
