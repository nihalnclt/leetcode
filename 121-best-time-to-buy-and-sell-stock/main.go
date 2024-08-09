package main

import (
	"fmt"
	"math"
)

func main() {
	prices := []int{1, 2}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	lsf := math.MaxInt
	op := 0
	pist := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < lsf {
			lsf = prices[i]
		}
		pist = prices[i] - lsf
		if op < pist {
			op = pist
		}
	}

	return op
}
