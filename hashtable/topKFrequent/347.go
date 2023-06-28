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

type couple struct {
	value int
	count int
}

type helper []*couple

func (h helper) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h helper) Less(i, j int) bool {
	return h[i].count > h[j].count
}

func (h helper) Len() int {
	return len(h)
}

func (h *helper) Push(x any) {
	item := x.(*couple)
	*h = append(*h, item)
}

func (h *helper) Pop() any {
	length := len(*h)
	item := (*h)[length-1]
	*h = (*h)[:length-1]
	return item
}

func four(nums []int, k int) []int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}
	h := helper{}
	for val, c := range count {
		heap.Push(&h, &couple{value: val, count: c})
	}
	result := []int{}
	for len(result) < k {
		item := heap.Pop(&h).(*couple)
		result = append(result, item.value)
	}
	return result
}

type Counter struct {
	Value int
	Count int
}

type Counters []*Counter

func (c Counters) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c Counters) Less(i, j int) bool {
	return c[i].Count > c[j].Count
}

func (c Counters) Len() int {
	return len(c)
}

func (c *Counters) Push(x any) {
	v := x.(*Counter)
	*c = append(*c, v)
}

func (c *Counters) Pop() any {
	v := (*c)[len(*c)-1]
	*c = (*c)[:len(*c)-1]
	return v
}

func five(nums []int, k int) []int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}
	var c Counters
	for i, v := range count {
		temp := &Counter{Value: i, Count: v}
		heap.Push(&c, temp)
	}
	var result []int

	for len(result) < k {
		item := heap.Pop(&c).(*Counter)
		result = append(result, item.Value)
	}
	return result
}

type SixCounter struct {
	Value, Count int
}
type sixCounters []*SixCounter

func (s sixCounters) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s sixCounters) Less(i, j int) bool {
	return s[i].Count > s[j].Count
}
func (s sixCounters) Len() int {
	return len(s)
}
func (s *sixCounters) Push(x interface{}) {
	v := x.(*SixCounter)
	*s = append(*s, v)
}
func (s *sixCounters) Pop() interface{} {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func six(nums []int, k int) []int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}
	var s sixCounters
	for i, v := range count {
		temp := &SixCounter{Value: i, Count: v}
		heap.Push(&s, temp)
	}
	var result []int
	for len(result) < k {
		v := heap.Pop(&s).(*SixCounter)
		result = append(result, v.Value)
	}
	return result
}

type SevenCounter struct {
	Value, Count int
}
type SevenCounters []*SevenCounter

func (s SevenCounters) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SevenCounters) Less(i, j int) bool {
	return s[i].Count > s[j].Count
}

func (s SevenCounters) Len() int {
	return len(s)
}

func (s *SevenCounters) Push(x interface{}) {
	v := x.(*SevenCounter)
	*s = append(*s, v)
}
func (s *SevenCounters) Pop() interface{} {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func seven(nums []int, k int) []int {
	count := make(map[int]int)
	for _, n := range nums {
		count[n]++
	}
	var s SevenCounters
	for i, v := range count {
		temp := &SevenCounter{Value: i, Count: v}
		heap.Push(&s, temp)
	}
	var result []int
	for len(result) < k {
		v := heap.Pop(&s).(*SixCounter)
		result = append(result, v.Value)
	}
	return result
}
