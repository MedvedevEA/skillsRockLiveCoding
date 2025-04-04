package main

import "fmt"

var s = "43212"

func shortestPalindrome(s string) string {
	pMaxLeft, pMaxRight, pLeft, pRight := "", "", "", ""

	for i := range s {
		pLeft = string(s[i]) + pLeft
		if s[:i+1] == pLeft {
			pMaxLeft = s[:i+1]
		}
		pRight = pRight + string(s[len(s)-(i+1)])
		if s[len(s)-(i+1):] == pRight {
			pMaxRight = s[len(s)-(i+1):]
		}
	}
	if len(pMaxLeft) > len(pMaxRight) {
		return pRight[:len(s)-len(pMaxLeft)] + s
	}
	return s + pLeft[len(pMaxRight):]

}

func main() {
	result := shortestPalindrome(s)
	fmt.Println(result)

}
