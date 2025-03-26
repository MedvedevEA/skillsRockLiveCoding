package main

import (
	"fmt"
	"math"
)

var prices = []int{7, 1, 5, 3, 6, 4}

func maxProfit(prices []int) int {
	minPrice := math.MaxInt
	maxProf := 0
	for i := range prices {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		if maxProf < prices[i]-minPrice {
			maxProf = prices[i] - minPrice
		}

	}
	return maxProf
}

func main() {
	result := maxProfit(prices)
	fmt.Println(result)
}
