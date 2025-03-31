package main

import "fmt"

var (
	s = "00110"
	k = 2
)

func main() {
	result := hasAllCodes(s, k)
	fmt.Println(result)

}
func hasAllCodes(s string, k int) bool {
	col := 1 << k
	f := make(map[string]bool, col)
	for i := 0; i < len(s)-k+1; i++ {
		if !f[s[i:i+k]] {
			col--
		}
		f[s[i:i+k]] = true
		if col == 0 {
			return col == 0
		}
	}
	return false

}
