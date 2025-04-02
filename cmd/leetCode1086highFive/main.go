package main

import (
	"fmt"
	"sort"
)

var items = [][]int{{1, 50}, {1, 100}, {1, 200}, {1, 300}, {1, 400}, {1, 200}, {2, 1000}, {2, 1000}, {2, 1000}, {2, 1000}, {2, 1000}, {2, 500}, {3, 60}, {3, 50}, {3, 40}, {3, 30}, {3, 20}, {3, 10}}

func HighFive(items [][]int) [][]int {
	sort.Slice(items, func(i int, j int) bool {
		return items[i][0] < items[j][0] || items[i][0] == items[j][0] && items[i][1] > items[j][1]
	})
	m := make(map[int][]int)

	for i := 0; i < len(items); i++ {
		if _, ok := m[items[i][0]]; !ok {
			m[items[i][0]] = []int{0, 0}
		}
		if m[items[i][0]][0] < 5 {
			m[items[i][0]][0]++
			m[items[i][0]][1] += items[i][1]
		}
	}
	result := [][]int{}
	for key := range m {
		result = append(result, []int{key, m[key][1] / m[key][0]})
	}

	return result

}

func main() {
	result := HighFive(items)
	fmt.Println(result)

}
