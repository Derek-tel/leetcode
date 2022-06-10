package main

import (
	"container/heap"
)

type item struct {
	val   int
	count int
}

type IHeap []*item

func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h IHeap) Less(i, j int) bool {
	return h[i].count > h[j].count
}

func (h IHeap) Len() int {
	return len(h)
}

func (h *IHeap) Push(x any) {
	item := x.(*item)
	*h = append(*h, item)
}

func (h *IHeap) Pop() any {
	len := len(*h)
	item := (*h)[len-1]
	*h = (*h)[:len-1]
	return item
}

func topKFrequent(nums []int, k int) []int {
	countNum := make(map[int]int)

	for _, num := range nums {
		countNum[num]++
	}
	iHeap := IHeap{}
	for num, count := range countNum {
		heap.Push(&iHeap, &item{val: num, count: count})
	}
	var result []int
	for len(result) < k {
		item := heap.Pop(&iHeap).(*item)
		result = append(result, item.val)
	}

	return result
}

type coupe struct {
	val   int
	count int
}

type CHeap []*coupe

func (h CHeap) Less(i, j int) bool {
	return h[i].count > h[j].count
}

func (h CHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h CHeap) Len() int {
	return len(h)
}

func (h *CHeap) Push(x any) {
	item := x.(*coupe)
	*h = append(*h, item)
}

func (h *CHeap) Pop() any {
	len := len(*h)
	item := (*h)[len-1]
	*h = (*h)[:len-1]
	return item
}

func test(nums []int, k int) []int {
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
	}
	h := CHeap{}
	for num, count := range counter {
		heap.Push(&h, &coupe{num, count})
	}
	var res []int
	for len(res) < k {
		item := heap.Pop(&h).(*coupe)
		res = append(res, item.val)
	}
	return res
}

type counter struct {
	value int
	count int
}

type GHeap []*counter

func (g GHeap) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}
func (g GHeap) Less(i, j int) bool {
	return g[i].count > g[j].count
}

func (g GHeap) Len() int {
	return len(g)
}
func (g *GHeap) Push(x any) {
	item := x.(*counter)
	*g = append(*g, item)
}

func (g *GHeap) Pop() any {
	length := len(*g)
	item := (*g)[length-1]
	*g = (*g)[:length-1]
	return item
}

func get(nums []int, k int) []int {
	count := make(map[int]int)
	for _, i := range nums {
		count[i]++
	}
	g := GHeap{}
	for num, c := range count {
		heap.Push(&g, &counter{value: num, count: c})
	}

	res := []int{}
	for len(res) < k {
		item := heap.Pop(&g).(*counter)
		res = append(res, item.value)
	}
	return res
}
