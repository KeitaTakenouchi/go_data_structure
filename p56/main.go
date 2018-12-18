package main

import (
	"fmt"
	"strings"
)

var n int
var m int

func main() {
	fmt.Scan(&n)
	fmt.Scan(&m)

	var str string
	fmt.Scan(&str)
	s := strings.Split(str, "")
	fmt.Scan(&str)
	t := strings.Split(str, "")

	v := rec(s, t)
	fmt.Println(v)
}

func rec(a []string, b []string) int {
	var ret int
	if len(a) == 0 || len(b) == 0 {
		ret = 0
	} else if a[0] == b[0] {
		ret = 1 + rec(a[1:], b[1:])
	} else {
		ret = max(rec(a[1:], b), rec(a, b[1:]))
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
