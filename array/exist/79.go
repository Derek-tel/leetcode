package main

import (
	"fmt"
	"time"
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

type pair2 struct {
	x, y int
}

var dir = []pair2{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func exist2(board [][]byte, word []byte) bool {
	h, w := len(board), len(board[0])
	visit := make([][]bool, h)
	for i := range visit {
		visit[i] = make([]bool, w)
	}

	var check func(i, j, k int) bool
	check = func(i, j, k int) bool {
		if board[i][j] != word[k] {
			return false
		}
		if len(word)-1 == k {
			return true
		}
		visit[i][j] = true
		defer func() { visit[i][j] = false }()
		for _, d := range direction {
			if newX, newY := i+d.x, j+d.y; newX >= 0 && newX < h && newY >= 0 && newY < w && !visit[newX][newY] {
				if check(newX, newY, k+1) {
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

func exist3(board [][]byte, word []byte) bool {
	h, w := len(board), len(board[0])
	visit := make([][]bool, h)
	for i := range visit {
		visit[i] = make([]bool, w)
	}
	var check func(int, int, int) bool
	check = func(i int, j int, k int) bool {
		if board[i][j] != word[k] {
			return false
		}
		if len(word)-1 == k {
			return true
		}
		visit[i][j] = true
		defer func() { visit[i][j] = false }()
		for _, d := range directions {
			if newX, newY := i+d.x, j+d.y; newX >= 0 && newX < h && newY >= 0 && newY < w && !visit[newX][newY] {
				if check(newX, newY, k+1) {
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

type dirPair struct {
	x int
	y int
}

var dirFour = []dirPair{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func four(board [][]byte, word []byte) bool {
	h, w := len(board), len(board[0])
	visit := make([][]bool, h)
	for i := 0; i < h; i++ {
		visit[i] = make([]bool, w)
	}
	var handler func(int, int, int) bool
	handler = func(x int, y int, k int) bool {
		if board[x][y] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true
		}
		visit[x][y] = true
		defer func() { visit[x][y] = false }()
		for _, p := range dirFour {
			if newX, newY := x+p.x, y+p.y; newX >= 0 && newX < h && newY >= 0 && newY < w && !visit[newX][newY] {
				if handler(newX, newY, k+1) {
					return true
				}
			}
		}
		return false
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if handler(i, j, 0) {
				return true
			}
		}
	}
	return false
}

type dirFivePair struct {
	x, y int
}

var dirFive = []dirFivePair{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

func five(board [][]byte, word []byte) bool {
	h, w := len(board), len(board[0])
	visit := make([][]bool, h)
	for i := 0; i < h; i++ {
		visit[i] = make([]bool, w)
	}
	var handler func(int, int, int) bool
	handler = func(x int, y int, k int) bool {
		if board[x][y] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true
		}
		visit[x][y] = true
		defer func() { visit[x][y] = false }()
		for _, p := range dirFive {
			if newX, newY := x+p.x, y+p.y; newX >= 0 && newX < h && newY >= 0 && newY < w && !visit[newX][newY] {
				if handler(newX, newY, k+1) {
					return true
				}
			}
		}
		return false
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if handler(i, j, 0) {
				return true
			}
		}
	}
	return false
}

type dirPairSix struct {
	x, y int
}

var dirSix = []dirPairSix{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

func six(board [][]byte, word []byte) bool {
	h, w := len(board), len(board[0])
	visit := make([][]bool, h)
	for i := 0; i < h; i++ {
		visit[i] = make([]bool, w)
	}
	var handler func(int, int, int) bool
	handler = func(x, y int, k int) bool {
		if board[x][y] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true
		}
		visit[x][y] = true
		defer func() { visit[x][y] = false }()
		for _, p := range dirSix {
			if newX, newY := x+p.x, y+p.y; newX >= 0 && newX < h && newY >= 0 && newY < w && !visit[newX][newY] {
				if handler(newX, newY, k+1) {
					return true
				}
			}
		}
		return false
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if handler(i, j, 0) {
				return true
			}
		}
	}
	return false
}

type dirPairSeven struct {
	x, y int
}

var dirSeven = []dirPairSeven{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

func seven(board [][]byte, word []byte) bool {
	h, w := len(board), len(board[0])
	visit := make([][]bool, h)
	for i := 0; i < len(visit); i++ {
		visit[i] = make([]bool, w)
	}
	var handler func(int, int, int) bool
	handler = func(x int, y int, k int) bool {
		if board[x][y] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true
		}
		visit[x][y] = true
		defer func() { visit[x][y] = false }()
		for _, p := range dirSeven {
			if newX, newY := x+p.x, y+p.y; newX >= 0 && newX < h && newY >= 0 && newY < w && !visit[newX][newY] {
				if handler(newX, newY, k+1) {
					return true
				}
			}
		}
		return false
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if handler(i, j, 0) {
				return true
			}
		}
	}
	return false
}

type dirPairEight struct {
	x, y int
}

var dirEight = []dirPairEight{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

func eight(board [][]byte, word []byte) bool {
	h, w := len(board), len(board[0])
	visit := make([][]bool, h)
	for i := 0; i < h; i++ {
		visit[i] = make([]bool, w)
	}
	var handler func(int, int, int) bool
	handler = func(index int, x int, y int) bool {
		if board[x][y] != word[index] {
			return false
		}
		if index == len(word)-1 {
			return true
		}
		visit[x][y] = true
		defer func() { visit[x][y] = false }()
		for _, d := range dirEight {
			if newX, newY := x+d.x, y+d.y; newX >= 0 && newX < h && newY >= 0 && newY < w && !visit[newX][newY] {
				if handler(index+1, newX, newY) {
					return true
				}
			}
		}
		return false
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if handler(0, i, j) {
				return true
			}
		}
	}
	return false
}

type dirPairNine struct {
	x, y int
}

var dirNine = []dirPairNine{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

func nine(board [][]byte, word string) bool {
	h, w := len(board), len(board[0])
	visit := make([][]bool, h)
	for i := 0; i < h; i++ {
		visit[i] = make([]bool, w)
	}
	var handler func(int, int, int) bool
	handler = func(index int, x int, y int) bool {
		if board[x][y] != word[index] {
			return false
		}
		if index == len(word)-1 {
			return true
		}
		visit[x][y] = true
		defer func() { visit[x][y] = false }()
		for _, v := range dirNine {
			if newX, newY := x+v.x, y+v.y; newX >= 0 && newX < h && newY >= 0 && newY < w && visit[newX][newY] == false {
				if handler(index+1, newX, newY) {
					return true
				}
			}
		}
		return false
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if handler(0, i, j) {
				return true
			}
		}
	}
	return false
}

type dirPairTen struct {
	x, y int
}

var dirTen = []dirPairTen{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

func ten(board [][]byte, word string) bool {
	h, w := len(board), len(board[0])
	visit := make([][]bool, h)
	for i := 0; i < h; i++ {
		visit[i] = make([]bool, w)
	}
	var handler func(int, int, int) bool
	handler = func(index int, x int, y int) bool {
		if board[x][y] != word[index] {
			return false
		}
		if index == len(word)-1 {
			return true
		}
		visit[x][y] = true
		defer func() { visit[x][y] = false }()
		for _, v := range dirTen {
			if newX, newY := x+v.x, y+v.y; newX >= 0 && newX < h && newY >= 0 && newY < w && !visit[newX][newY] {
				if handler(index+1, newX, newY) {
					return true
				}
			}
		}
		return false
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if handler(0, i, j) {
				return true
			}
		}
	}
	return false
}

type elevenPair struct {
	x, y int
}

var dirEleven = []elevenPair{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

func eleven(board [][]byte, word string) bool {
	h, w := len(board), len(board[0])
	visit := make([][]bool, h)
	for i := 0; i < h; i++ {
		visit[i] = make([]bool, w)
	}
	var handler func(int, int, int) bool
	handler = func(index int, x int, y int) bool {
		if board[x][y] != word[index] {
			return false
		}
		if index == len(word)-1 {
			return true
		}
		visit[x][y] = true
		defer func() { visit[x][y] = false }()
		for _, v := range dirEleven {
			if newX, newY := x+v.x, y+v.y; newX >= 0 && newX < h && newY >= 0 && newY < w && !visit[newX][newY] {
				if handler(index+1, newX, newY) {
					return true
				}
			}
		}
		return false
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if handler(0, i, j) {
				return true
			}
		}
	}
	return false
}

type twelvePair struct {
	x, y int
}

var dirTwelve = []twelvePair{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

func twelve(board [][]byte, word string) bool {
	h, w := len(board), len(board[0])
	visit := make([][]bool, h)
	for i := 0; i < len(visit); i++ {
		visit[i] = make([]bool, w)
	}
	var handler func(int, int, int) bool
	handler = func(x int, y int, index int) bool {
		if board[x][y] != word[index] {
			return false
		}
		if index == len(word)-1 {
			return true
		}
		visit[x][y] = true
		defer func() { visit[x][y] = false }()
		for _, v := range dirTwelve {
			if newX, newY := x+v.x, y+v.y; newX >= 0 && newX < h && newY >= 0 && newY < w && !visit[newX][newY] {
				if handler(newX, newY, index+1) {
					return true
				}
			}
		}
		return false
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if handler(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func returnValues() (result int) {
	//var result int
	defer func() {
		result++
		fmt.Println("defer")
	}()
	return
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
	fmt.Println("----")
	current := time.Now()
	fmt.Println(current)
	updateTime := time.Unix(current.Unix(), 0)
	fmt.Println(updateTime)

	fmt.Println(updateTime.UnixNano() / 1e6)

	vid := map[int]int{
		1: 2,
		3: 4,
	}
	for i := range vid {
		fmt.Println(i)
	}

}
