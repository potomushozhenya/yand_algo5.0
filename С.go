package main

import (
	"fmt"
	"os"
)

func main() {
	var n, curr int
	result := 0
	fmt.Fscan(os.Stdin, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(os.Stdin, &curr)
		result += curr / 4
		curr = curr % 4
		switch curr {
		case 1:
			result += 1
		case 2, 3:
			result += 2
		}
	}
	fmt.Println(result)
}
