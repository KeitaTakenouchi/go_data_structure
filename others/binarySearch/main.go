package main

import (
	"fmt"
	"sort"
)

func main() {
	values := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

	for _, v := range values {
		fmt.Println(v, divide143(v))
	}

	index := sort.Search(len(values), func(i int) bool {
		return divide143(values[i])
	})

	fmt.Println("index=", index, "value=", values[index])

}

func divide143(n int) bool {
	for i := 2; i <= n; i++ {
		if 143%i == 0 {
			return true
		}
	}
	return false
}
