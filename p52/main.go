package main

import (
	"fmt"
)

var n int
var W int

var weights []int
var values []int

var memo [][]int

func main() {
	fmt.Scan(&n)
	weights = make([]int, n)
	values = make([]int, n)

	for i := 0; i < n; i++ {
		var w int
		fmt.Scan(&w)
		weights[i] = w
	}

	for i := 0; i < n; i++ {
		var v int
		fmt.Scan(&v)
		values[i] = v
	}
	fmt.Scan(&W)

	memo = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		memo[i] = make([]int, W+1)
		for k := 0; k <= W; k++ {
			memo[i][k] = -1
		}
	}
	for i := 0; i < n; i++ {
		fmt.Println(memo[i])
	}

	total := rec(0, W)
	fmt.Println(total)
}

func rec(i int, rest int) int {
	if memo[i][rest] > 0 {
		return memo[i][rest]
	}
	if i == n {
		return 0
	}
	if rest < weights[i] {
		return rec(i+1, rest)
	}
	memo[i][rest] = max(values[i]+rec(i+1, rest-weights[i]), rec(i+1, rest))
	return memo[i][rest]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
