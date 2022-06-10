package main

import (
	"fmt"
)

type pair struct {
	x, y int
}

var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func exist(board [][]byte, word string) bool {
	h, w := len(board), len(board[0])
	vis := make([][]bool, h)
	for i := range vis {
		vis[i] = make([]bool, w)
	}
	var check func(i, j, k int) bool
	check = func(i, j, k int) bool {
		if board[i][j] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true
		}
		vis[i][j] = true
		defer func() { vis[i][j] = false }()
		for _, dir := range directions {
			if newI, newJ := i+dir.x, j+dir.y; newI >= 0 && newI < h && newJ >= 0 && newJ < w && !vis[newI][newJ] {
				if check(newI, newJ, k+1) {
					return true
				}
			}
		}
		return false
	}
	for i, row := range board {
		for j := range row {
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false
}

type pair1 struct {
	x, y int
}

var direction = []pair1{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

func exist1(board [][]byte, word string) bool {
	h, w := len(board), len(board[0])
	visit := make([][]bool, h)
	for i := 0; i < h; i++ {
		visit[i] = make([]bool, w)
	}

	var check func(i, j, k int) bool

	check = func(i, j, k int) bool {
		if board[i][j] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true
		}
		visit[i][j] = true
		defer func() { visit[i][j] = false }()
		for _, dir := range direction {
			if newI, newJ := i+dir.x, j+dir.y; newI >= 0 && newI < h && newJ >= 0 && newJ < w && !visit[newI][newJ] {
				if check(newI, newJ, k+1) {
					return true
				}
			}
		}

		return false
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func returnValues() int {
	var result int
	defer func() {
		result++
		fmt.Println("defer")
	}()
	return result
}

type Trie struct {
	IsWord   bool
	Children map[rune]*Trie
}

func main() {
	board := [][]byte{[]byte("ABCE"), []byte("SFCS"), []byte("ADEE")}
	word := "ABCCED"
	bol := exist(board, word)
	fmt.Println(bol)

	fmt.Println(returnValues())

	for _, r := range "hello,世界" {
		fmt.Printf("%T", r)
	}

}
