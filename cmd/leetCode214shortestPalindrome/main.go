package main

import "fmt"

var s = "43212"

func shortestPalindrome(s string) string {
	pMaxLeft, pLeft, pRight := "", "", ""

	for i := range s {
		pLeft = string(s[i]) + pLeft
		if s[:i+1] == pLeft {
			pMaxLeft = s[:i+1]
		}
		pRight = pRight + string(s[len(s)-(i+1)])
	}
	return pRight[:len(s)-len(pMaxLeft)] + s

}

func main() {
	result := shortestPalindrome(s)
	fmt.Println(result)

}
