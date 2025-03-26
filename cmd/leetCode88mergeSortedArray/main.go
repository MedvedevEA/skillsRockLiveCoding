package main

import "fmt"

var (
	nums1 = []int{1, 2, 3, 0, 0, 0}
	nums2 = []int{2, 5, 6}
	m     = 3
	n     = 3
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums2) == 0 {
		return
	}
	for i := 0; i < m; i++ {
		if nums1[i] > nums2[0] {
			nums1[i], nums2[0] = nums2[0], nums1[i]
			for j := 0; j < len(nums2)-1; j++ {
				if nums2[j] < nums2[j+1] {
					break
				}
				nums2[j], nums2[j+1] = nums2[j+1], nums2[j]
			}

		}

	}
	for i := range nums2 {
		nums1[m+i] = nums2[i]
	}

}

func main() {

	merge(nums1, m, nums2, n)
	fmt.Println(nums1)

}
