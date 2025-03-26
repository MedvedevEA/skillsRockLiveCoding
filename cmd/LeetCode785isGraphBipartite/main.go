package main

import (
	"fmt"
)

// var graph = [][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}
// var graph = [][]int{{1}, {0}}
var graph = [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}

func isBipartite(graph [][]int) bool {
	result := true

	nodeMark := make([]int, len(graph))

	for i := range nodeMark {
		if nodeMark[i] == 0 {

			nodeStack := [][]int{{i, -1}}

			for len(nodeStack) > 0 && result {
				node := nodeStack[len(nodeStack)-1]
				nodeStack = nodeStack[:len(nodeStack)-1]
				if nodeMark[node[0]] == 0 {
					nodeMark[node[0]] = node[1]
					for i := range graph[node[0]] {
						nodeStack = append(nodeStack, []int{graph[node[0]][i], node[1] * -1})
					}
				} else if nodeMark[node[0]] != node[1] {
					result = false
				}
			}

		}
	}

	return result
}

func main() {
	result := isBipartite(graph)
	fmt.Println(result)

}
