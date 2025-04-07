// ?
package main

import "fmt"

var n = 10

func getMoneyAmount(n int) int {

	count := 0
	for n > 1 {
		count++
		n = n / 2
	}
	return count
}

func main() {
	result := getMoneyAmount(n)
	fmt.Println(result)
}
