package main

import (
	"fmt"
	"maps"
	"slices"
)

var n = 3

func main() {
	result := generateParenthesis(n)
	fmt.Println(result)
}

func generateParenthesis(n int) []string {
	if n < 2 {
		return []string{"()"}
	}

	h := make(map[string]struct{})

	s := generateParenthesis(n - 1)
	for i := range s {
		for j := 0; j < len(s[i])+1; j++ {
			key := s[i][0:j] + "()" + s[i][j:]
			h[key] = struct{}{}
		}
	}

	return slices.Collect(maps.Keys(h))

}
