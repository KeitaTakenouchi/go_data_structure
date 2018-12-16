package main

import (
	"fmt"
	"strings"
)

var n, m int
var values [][]bool

func main() {
	fmt.Scan(&n)
	fmt.Scan(&m)

	values = make([][]bool, n)
	for i := 0; i < n; i++ {
		values[i] = make([]bool, m)
	}

	for i := 0; i < n; i++ {
		var str string
		fmt.Scan(&str)
		chars := strings.Split(str, "")
		for j := 0; j < m; j++ {
			if chars[j] == "*" {
				values[i][j] = true
			} else {
				values[i][j] = false
			}
		}
	}

	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if values[i][j] {
				fill(i, j)
				count++

				fmt.Println()
				for i := 0; i < n; i++ {
					for j := 0; j < m; j++ {
						if values[i][j] {
							fmt.Print("*")
						} else {
							fmt.Print("-")
						}
					}
					fmt.Println()
				}
			}
		}
	}
	fmt.Println(count)

}

func fill(i, j int) {
	if i < 0 || j < 0 || i >= n || j >= m {
		return
	}
	if !values[i][j] {
		return
	}
	values[i][j] = false

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			fill(i+dx, j+dy)
		}
	}
}
