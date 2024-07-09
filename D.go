package main

import (
	"fmt"
	"os"
)

func mark(chessboard *[10][10]byte, piece byte, posx int, posy int, counter *int) {
	var dx [4]int
	var dy [4]int
	switch piece {
	case 'B':
		dx = [4]int{1, 1, -1, -1}
		dy = [4]int{1, -1, 1, -1}
	case 'R':
		dx = [4]int{1, -1, 0, 0}
		dy = [4]int{0, 0, 1, -1}
	}
	for i := 0; i < 4; i++ {
		currx := posx
		curry := posy
		for {
			currx += dx[i]
			curry += dy[i]
			curr := chessboard[currx][curry]
			if curr != '0' && curr != '*' {
				break
			}
			if curr != '0' {
				*counter++
				chessboard[currx][curry] = '0'
			}
		}
	}
}
func fillBorder(chessboard *[10][10]byte) {
	for i := 0; i < 10; i += 9 {
		for j := 0; j < 10; j++ {
			chessboard[i][j] = '#'
			chessboard[j][i] = '#'
		}
	}
}
func printChessboard(chessboard *[10][10]byte) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("%c", chessboard[i][j])
		}
		fmt.Printf("%c", '\n')
	}
}
func main() {
	var chessboard [10][10]byte
	fillBorder(&chessboard)
	counter := 0
	for i := 1; i < 9; i++ {
		var input string
		fmt.Fscan(os.Stdin, &input)
		for j := 1; j < 9; j++ {
			chessboard[i][j] = input[j-1]
		}
	}
	for i := 1; i < 9; i++ {
		for j := 1; j < 9; j++ {
			if chessboard[i][j] == 'R' || chessboard[i][j] == 'B' {
				counter++
				mark(&chessboard, chessboard[i][j], i, j, &counter)
			}
		}
	}
	fmt.Println(64 - counter)
	printChessboard(&chessboard)
}
