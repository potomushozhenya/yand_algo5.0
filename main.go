package main

import (
	"fmt"
	"os"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var p, q, v, m int
	fmt.Fscan(os.Stdin, &p, &q, &v, &m)
	l1 := p - q
	r1 := p + q
	l2 := v - m
	r2 := v + m
	if max(l1, l2) <= min(r1, r2) {
		fmt.Print(Abs(max(r1, r2)-min(l1, l2)) + 1)
	} else {
		fmt.Print(Abs(r1-l1) + Abs(r2-l2) + 2)
	}
}
