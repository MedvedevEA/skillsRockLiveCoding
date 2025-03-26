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
	// Перебираем все вершины
	for i := range nodeMark {
		//Если находим вершину без маркировки, то начинаем обход
		if nodeMark[i] == 0 {
			//В стек вершин, которые необходимо обойти, состоит из индекса вершины и значения маркера
			nodeStack := [][]int{{i, -1}}

			for len(nodeStack) > 0 && result {
				// Берем вершину из стека
				node := nodeStack[len(nodeStack)-1]
				nodeStack = nodeStack[:len(nodeStack)-1]
				// Если вершина без маркера, то маркируем и добавляем связанные с ней вершины в стек
				if nodeMark[node[0]] == 0 {
					nodeMark[node[0]] = node[1]
					for i := range graph[node[0]] {
						nodeStack = append(nodeStack, []int{graph[node[0]][i], node[1] * -1})
					}
					// Если вершина уже промаркирована, то проверяем значение маркера
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
