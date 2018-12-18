package main

import "fmt"

var n int
var W int
var w []int
var v []int

var memo []int

func main() {
	fmt.Scan(&n)
	fmt.Scan(&W)

	w = make([]int, n)
	for i := 0; i < n; i++ {
		var k int
		fmt.Scan(&k)
		w[i] = k
	}

	v = make([]int, n)
	for i := 0; i < n; i++ {
		var k int
		fmt.Scan(&k)
		v[i] = k
	}

	memo = make([]int, W+1)
	for i := 0; i < len(memo); i++ {
		memo[i] = -1
	}
	v := rec(W)
	fmt.Println(v)
}

func rec(rest int) int {
	if memo[rest] > 0 {
		return memo[rest]
	}
	m := 0
	for i := 0; i < n; i++ {
		if rest >= w[i] {
			m = max(m, v[i]+rec(rest-w[i]))
		}
	}
	memo[rest] = m
	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
