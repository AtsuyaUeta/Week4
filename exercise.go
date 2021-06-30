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
		b.tokens[x+3*y] = 0
	}
}

func (b *Board) get(x, y int) string {
	if b.tokens[x+3*y] == 1 {
		return "o"
	} else if b.tokens[x+3*y] == 0 {
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
	}
}
