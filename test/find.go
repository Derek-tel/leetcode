package main

import "fmt"

type Pair struct {
	Val, Idx int
}

func demo1(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	var pairs []Pair
	for i, num := range nums {
		pairs = append(pairs, Pair{Idx: i, Val: num})
	}
	size := len(pairs)
	build(pairs, size)
	k := 1
	for i := 0; i < k; i++ {
		pairs[0], pairs[len(pairs)-1] = pairs[len(pairs)-1], pairs[0]
		size--
		modify(pairs, 0, size)
	}
	return pairs[0].Idx
}

func build(pairs []Pair, size int) {
	for i := size / 2; i >= 0; i-- {
		modify(pairs, i, size)
	}
}

func modify(pairs []Pair, index int, size int) {
	left := index * 2
	right := index*2 + 1
	largest := index
	if left < size && pairs[left].Val > pairs[largest].Val {
		largest = left
	}
	if right < size && pairs[right].Val > pairs[largest].Val {
		largest = right
	}
	if largest != index {
		pairs[largest], pairs[index] = pairs[index], pairs[largest]
		modify(pairs, largest, size)
	}
}

func main() {
	unsortSlice := []int{13, 5, 7, 9}
	fmt.Println(demo1(unsortSlice))
}
