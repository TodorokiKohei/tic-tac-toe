package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Board struct {
	tokens []int
}

func (b *Board) put(x, y int, u string) error {
	if (0 <= x && x <= 2) || (0 <= y && y <= 2) {
		if (u != "x" && u != "o") || b.tokens[x+3*y] != 0 {
			// ox以外の入力, 既に記入されているマスへの入力
			return errors.New("Invalid operation")
		}
		if u == "o" {
			b.tokens[x+3*y] = 1
		} else if u == "x" {
			b.tokens[x+3*y] = -1
		}
		return nil
	}
	// 配列外参照
	return errors.New("Out of range")
}

func (b *Board) get(x, y int) string {
	if b.tokens[x+3*y] == 1 {
		return "o"
	} else if b.tokens[x+3*y] == -1 {
		return "x"
	}
	return "."
}

func (b *Board) judge() bool {
	for i := 0; i < 3; i++ {
		// 縦列判定
		sum := b.tokens[i] + b.tokens[i+3] + b.tokens[i+6]
		if sum == 3 || sum == -3 {
			return true
		}
		// 横列判定
		sum = b.tokens[i*3] + b.tokens[i*3+1] + b.tokens[i*3+2]
		if sum == 3 || sum == -3 {
			return true
		}
	}
	// 斜め判定
	sum := b.tokens[0] + b.tokens[4] + b.tokens[8]
	if sum == 3 || sum == -3 {
		return true
	}
	sum = b.tokens[2] + b.tokens[4] + b.tokens[6]
	if sum == 3 || sum == -3 {
		return true
	}
	return false
}

func (b *Board) print() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(b.get(i, j))
		}
		fmt.Println()
	}
}

func main() {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	var str string
	for {
		fmt.Print("Player 1: Input(x, y) ")
		fmt.Scan(&str)
		ope := strings.Split(str, ",")
		x, _ := strconv.Atoi(ope[0])
		y, _ := strconv.Atoi(ope[1])
		b.put(x, y, "o")
		b.print()
	}
}
