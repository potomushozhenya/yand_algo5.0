package main

import (
	"fmt"
	"os"
)

func main() {
	var g1, g2, h1, h2, home int
	var input string
	fmt.Fscan(os.Stdin, &input)
	g1 = int(input[0]) - 48
	g2 = int(input[2]) - 48
	fmt.Fscan(os.Stdin, &input)
	h1 = int(input[0]) - 48
	h2 = int(input[2]) - 48
	fmt.Fscan(os.Stdin, &home)
	curr := g2 + h2 - h1 - g1
	if curr < 0 {
		fmt.Print(0)
	} else {
		if (home == 1 && (curr+h1) <= g2) || (home == 2 && g1 <= h2) {
			curr += 1
		}
		fmt.Print(curr)
	}
}
