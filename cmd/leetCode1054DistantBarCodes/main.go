package main

import (
	"fmt"
	"sort"
)

var barcodes = []int{1, 1, 1, 2, 2, 2, 2}

func rearrangeBarcodes(barcodes []int) []int {
	barcodesCount := make(map[int]int, 0)
	for i := range barcodes {
		barcodesCount[barcodes[i]]++
	}
	heap := make([][]int, 0, len(barcodesCount))
	for key, value := range barcodesCount {
		heap = append(heap, []int{key, value})
	}
	sort.Slice(heap, func(i int, j int) bool { return heap[i][1] > heap[j][1] })
	fmt.Println(heap)
	result := []int{}
	for len(heap) > 0 && heap[0][1] > 0 {
		index := -1
		if len(result) == 0 || result[len(result)-1] != heap[0][0] {
			index = 0
		} else if len(heap) > 1 && heap[1][1] > 0 {
			index = 1
		} else {
			return []int{}
		}
		result = append(result, heap[index][0])
		heap[index][1]--
		for i := index; i < len(heap)-1; i++ {
			if heap[i][1] >= heap[i+1][1] {
				break
			}
			heap[i], heap[i+1] = heap[i+1], heap[i]
		}
	}
	return result
}

func main() {
	result := rearrangeBarcodes(barcodes)
	fmt.Println(result)

}
