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

var demo []int

type demoHeap []int

func (d demoHeap) Less(i, j int) bool {
	return demo[d[i]] > demo[d[j]]
}
func (d demoHeap) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
func (d demoHeap) Len() int {
	return len(d)
}
func (d *demoHeap) Push(x interface{}) {
	item := x.(int)
	*d = append(*d, item)
}
func (d *demoHeap) Pop() interface{} {
	length := len(*d)
	item := (*d)[length-1]
	*d = (*d)[:length-1]
	return item
}

func four(nums []int, k int) []int {
	demo = nums
	d := make(demoHeap, k)
	for i := 0; i < k; i++ {
		d[i] = i
	}
	heap.Init(&d)
	result := []int{}
	result = append(result, nums[d[0]])
	for i := k; i < len(nums); i++ {
		heap.Push(&d, i)
		for d[0] <= i-k {
			heap.Pop(&d)
		}
		result = append(result, nums[d[0]])
	}
	return result
}

var fiveTag []int

type FiveHead []int

func (f FiveHead) Less(i, j int) bool {
	return fiveTag[f[i]] > fiveTag[f[j]]
}

func (f FiveHead) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f FiveHead) Len() int {
	return len(f)
}

func (f *FiveHead) Push(x interface{}) {
	item := x.(int)
	*f = append(*f, item)
}

func (f *FiveHead) Pop() interface{} {
	length := len(*f)
	top := (*f)[length-1]
	*f = (*f)[:length-1]
	return top
}

func five(nums []int, k int) []int {
	fiveTag = nums
	handler := FiveHead{}
	for i := 0; i < k; i++ {
		handler = append(handler, i)
	}
	heap.Init(&handler)
	result := []int{nums[handler[0]]}
	for i := k; i < len(nums); i++ {
		heap.Push(&handler, i)
		for handler[0] <= i-k {
			heap.Pop(&handler)
		}
		result = append(result, nums[handler[0]])
	}
	return result
}

var sixTag []int

type sixHead []int

func (s sixHead) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sixHead) Less(i, j int) bool {
	return sixTag[s[i]] > sixTag[s[j]]
}

func (s sixHead) Len() int {
	return len(s)
}

func (s *sixHead) Push(x interface{}) {
	item := x.(int)
	*s = append(*s, item)
}
func (s *sixHead) Pop() interface{} {
	length := len(*s)
	v := (*s)[length-1]
	*s = (*s)[:length-1]
	return v
}

func six(nums []int, k int) []int {
	sixTag = nums

	handler := sixHead{}
	for i := 0; i < k; i++ {
		heap.Push(&handler, i)
	}

	result := []int{nums[handler[0]]}
	for i := k; i < len(nums); i++ {
		heap.Push(&handler, i)
		for handler[0] <= i-k {
			heap.Pop(&handler)
		}
		result = append(result, nums[handler[0]])
	}
	return result
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, -1, 3, 6, 7}, 3))
	d := make([]int, 2)
	for i := 0; i < 5; i++ {
		d = append(d, 1)
	}
	fmt.Println(d)

}
