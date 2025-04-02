package main

import "fmt"

var paths = [][]string{{"A", "B"}, {"B", "C"}, {"C", "D"}, {"D", "E"}}

func destCity(paths [][]string) string {
	m := make(map[string]int, len(paths)+1)
	for i := range paths {
		m[paths[i][0]]--
		m[paths[i][1]]++
	}
	for key := range m {
		if m[key] == 1 {
			return key
		}
	}
	return ""

}

func main() {
	result := destCity(paths)
	fmt.Println(result)

}
