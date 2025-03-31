package main

import "fmt"

var (
	arr = []int{2, 3, 8, 9, 10, 11}
	k   = 2
)

func main() {
	result := findKthPositive(arr, k)
	fmt.Println(result)

}
func findKthPositive(arr []int, k int) int {
	if k <= arr[0]-1 {
		return k
	}
	k -= arr[0] - 1
	n := len(arr)
	for i := 0; i < n-1; i++ {
		currMissing := arr[i+1] - arr[i] - 1
		if k <= currMissing {
			return arr[i] + k
		}
		k -= currMissing
	}
	return arr[n-1] + k
}
