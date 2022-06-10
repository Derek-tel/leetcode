package main

import (
	"container/heap"
	"fmt"
)

var a []int

type maxHeap []int

func (h maxHeap) Less(i, j int) bool {
	return a[h[i]] > a[h[j]]
}
func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h maxHeap) Len() int {
	return len(h)
}

func (h *maxHeap) Pop() any {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

func (h *maxHeap) Push(v any) {
	*h = append(*h, v.(int))
}

func maxSlidingWindow(nums []int, k int) []int {
	a = nums
	q := make(maxHeap, k)
	for i := 0; i < k; i++ {
		q[i] = i
	}
	heap.Init(&q)

	length := len(nums)
	ans := make([]int, 1, length-k+1)
	ans[0] = nums[q[0]]

	for i := k; i < length; i++ {
		heap.Push(&q, i)
		for q[0] <= i-k {
			heap.Pop(&q)
		}
		ans = append(ans, nums[q[0]])
	}

	return ans
}

var number []int

type IndexHeap []int

func (h IndexHeap) Less(i, j int) bool {
	return number[h[i]] > number[h[j]]
}

func (h IndexHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h IndexHeap) Len() int {
	return len(h)
}

func (h *IndexHeap) Push(x any) {
	item := x.(int)
	*h = append(*h, item)
}

func (h *IndexHeap) Pop() any {
	length := len(*h)
	item := (*h)[length-1]
	*h = (*h)[:length-1]
	return item
}

func test(nums []int, k int) []int {
	number = nums

	h := make(IndexHeap, k)
	for i := 0; i < k; i++ {
		h[i] = i
	}
	heap.Init(&h)

	ans := []int{}
	ans = append(ans, nums[h[0]])

	for i := k; i < len(nums); i++ {
		heap.Push(&h, i)
		for h[0] <= i-k {
			heap.Pop(&h)
		}
		ans = append(ans, nums[h[0]])
	}
	return ans
}

var numbers []int

type GHeap []int //存储number 的 key，用number value来排序key

func (g GHeap) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

func (g GHeap) Less(i, j int) bool {
	return number[g[i]] > number[g[j]]
}

func (g GHeap) Len() int {
	return len(g)
}

func (g *GHeap) Push(x any) {
	item := x.(int)
	*g = append(*g, item)
}

func (g *GHeap) Pop() any {
	length := len(*g)
	item := (*g)[length-1]
	*g = (*g)[:length-1]
	return item
}

func get(nums []int, k int) []int {
	numbers = nums
	g := make(GHeap, k)
	for i := 0; i < k; i++ {
		g = append(g, i)
	}
	heap.Init(&g)
	ans := []int{}
	ans = append(ans, nums[g[0]])
	for i := k; i < len(nums); i++ {
		heap.Push(&g, i)
		for g[0] <= i-k {
			heap.Pop(&g)
		}
		ans = append(ans, nums[g[0]])
	}
	return ans
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, -1, 3, 6, 7}, 3))
}
