package main

import (
	"fmt"
	"sort"
)

func main() {
	index := sort.Search(100, func(i int) bool {
		return divide143(i)
	})

	fmt.Println("index=", index)
}

func divide143(n int) bool {
	for i := 2; i <= n; i++ {
		if 143%i == 0 {
			return true
		}
	}
	return false
}
