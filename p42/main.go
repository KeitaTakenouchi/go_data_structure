package main

import (
	"fmt"
)

var vs = [6]int{1, 5, 10, 50, 100, 500}
var ns [6]int
var total int

func main() {
	for i := 0; i < 6; i++ {
		fmt.Scan(&ns[i])
	}
	fmt.Scan(&total)

	count := 0
	for i := 5; i >= 0; i-- {
		num := min(total/vs[i], ns[i])
		total = total - num*vs[i]
		count = count + num
	}
	fmt.Println(count)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
