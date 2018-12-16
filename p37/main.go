package main

import (
	"fmt"
	"strings"
)

var values [][]string
var n int
var m int

type p struct {
	i, j int
}

func main() {
	fmt.Scan(&n)
	fmt.Scan(&m)

	values = make([][]string, n)
	for i := 0; i < n; i++ {
		values[i] = make([]string, m)
		var str string
		fmt.Scan(&str)
		chars := strings.Split(str, "")
		values[i] = chars
	}

	var dist [][]int
	dist = make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dist[i][j] = -1
		}
	}

	// find start
	var start p
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if values[i][j] == "S" {
				start = p{i, j}
				dist[i][j] = 0
				break
			}
		}
	}

	queue := make([]p, 0)
	queue = append(queue, start)

	for len(queue) > 0 {
		target := queue[0]
		queue = queue[1:]

		if values[target.i][target.j] == "G" {
			fmt.Println(dist[target.i][target.j])
			break
		}

		// next
		for _, d := range [4]p{p{-1, 0}, p{1, 0}, p{0, -1}, p{0, 1}} {
			nextx := target.i + d.i
			nexty := target.j + d.j
			if !(nextx >= 0 && nextx < n && nexty >= 0 && nexty < m) {
				continue
			}
			if values[nextx][nexty] != "#" && dist[nextx][nexty] < 0 {
				queue = append(queue, p{nextx, nexty})
				dist[nextx][nexty] = dist[target.i][target.j] + 1
			}
		}
	}

}
