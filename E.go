package main

import (
	"fmt"
	"os"
)

func recStr(n int, k int, d int) {
	fmt.Print(n)
	for i := 1; i < d; i++ {
		fmt.Printf("%c", '0')
	}
}
func sol(n int, k int, d int) int {
	for i := 0; i <= 9; i++ {
		curr := n*10 + i
		if curr%k == 0 {
			recStr(curr, k, d)
			return 0
		}
	}
	return -1
}
func main() {
	var n, k, d int
	fmt.Fscan(os.Stdin, &n, &k, &d)
	if sol(n, k, d) != 0 {
		fmt.Println(-1)
	}
}
