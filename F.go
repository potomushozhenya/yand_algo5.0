package main

import (
	"fmt"
)

func main() {
	var n, curr int
	var even bool
	fmt.Scan(&n)
	arr := make([]bool, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&curr)
		if curr%2 == 0 {
			even = true
		} else {
			even = false
		}
		arr[i] = even
	}
	prev := arr[0]
	for i := 1; i < n; i++ {
		if !prev && !arr[i] {
			fmt.Printf("%c", 'x')

		} else {
			if !prev || !arr[i] {
				prev = false
			}
			fmt.Printf("%c", '+')
		}
	}
}
