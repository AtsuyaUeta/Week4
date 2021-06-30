package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Board struct {
	tokens []int
}

func (b *Board) put(x, y int, u string) {
	if u == "o" {
		b.tokens[x+3*y] = 1
	} else if u == "x" {
		b.tokens[x+3*y] = 10
	}
}

func (b *Board) get(x, y int) string {
	if b.tokens[x+3*y] == 1 {
		return "o"
	} else if b.tokens[x+3*y] == 10 {
		return "x"
	} else {
		return "."
	}
}

func printBoard(b *Board) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%s", b.get(i, j))
		}
		fmt.Printf("\n")
	}
}

func judge(b *Board) int {
	sums := [8]int{0, 0, 0, 0, 0, 0, 0, 0}
	sums[0] = b.tokens[0] + b.tokens[1] + b.tokens[2]
	sums[1] = b.tokens[3] + b.tokens[4] + b.tokens[5]
	sums[2] = b.tokens[6] + b.tokens[7] + b.tokens[8]
	sums[3] = b.tokens[0] + b.tokens[3] + b.tokens[6]
	sums[4] = b.tokens[1] + b.tokens[4] + b.tokens[7]
	sums[5] = b.tokens[2] + b.tokens[5] + b.tokens[8]
	sums[6] = b.tokens[0] + b.tokens[4] + b.tokens[8]
	sums[7] = b.tokens[2] + b.tokens[4] + b.tokens[6]
	for _, v := range sums {
		if v == 3 {
			return 1
		} else if v == 30 {
			return 2
		}
	}
	return 0

}

func main() {
	b := &Board{
		tokens: []int{2, 2, 2, 2, 2, 2, 2, 2, 2},
	}
	var place string
	for i := 0; i < 9; i++ {
		if i%2 == 0 {
			fmt.Printf("Player 1: Input (x,y) ")
			fmt.Scan(&place)
			p := strings.Split(place, "")
			x, _ := strconv.Atoi(p[0])
			y, _ := strconv.Atoi(p[2])
			b.put(x, y, "o")
		} else {
			fmt.Printf("Player 2: Input (x,y) ")
			fmt.Scan(&place)
			p := strings.Split(place, "")
			x, _ := strconv.Atoi(p[0])
			y, _ := strconv.Atoi(p[2])
			b.put(x, y, "x")
		}
		printBoard(b)

		v := judge(b)
		if v == 1 {
			fmt.Printf("Player 1 won\n")
			break
		} else if v == 2 {
			fmt.Printf("Player 2 won\n")
			break
		}
	}
	v := judge(b)
	if v == 0 {
		fmt.Printf("Draw\n")
	}
}
