package main

import "fmt"

type pair struct {
	x, y int
}

var dir = []pair{{0, 1}, {0, -1}, {-1, 0}, {1, 0}, {-1, 1}, {-1, -1}, {1, -1}, {1, 1}}

func queensAttackTheKing(queens [][]int, king []int) [][]int {
	grid := make([][]int, 8)
	for i := 0; i < 8; i++ {
		grid[i] = make([]int, 8)
	}
	for _, queen := range queens {
		grid[queen[0]][queen[1]] = 1
	}
	res := [][]int{}
	x, y := king[0], king[1]
	for i := 0; i < 8; i++ {
		direction := dir[i]
		newX, newY := x+direction.x, y+direction.y
		for newX >= 0 && newX < 8 && newY >= 0 && newY < 8 {
			if grid[newX][newY] == 1 {
				res = append(res, []int{newX, newY})
				break
			}
			newX = newX + direction.x
			newY = newY + direction.y
		}
	}
	return res
}

func main() {
	queue := [][]int{{0, 1}, {1, 0}, {4, 0}, {0, 4}, {3, 3}, {2, 4}}
	king := []int{0, 0}
	fmt.Println(queensAttackTheKing(queue, king))
}
