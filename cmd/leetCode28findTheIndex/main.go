package main

import "fmt"

var (
	haystack = "saadbutsad"
	needle   = "sad"
)

func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

func main() {

	result := strStr(haystack, needle)
	fmt.Println(result)

}
