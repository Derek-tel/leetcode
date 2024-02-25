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

type eightCounter struct {
	Value, Count int
}
type EightCounterList []*eightCounter

func (e EightCounterList) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
func (e EightCounterList) Less(i, j int) bool {
	return e[i].Count > e[j].Count
}
func (e EightCounterList) Len() int {
	return len(e)
}

func (e *EightCounterList) Push(x interface{}) {
	v := x.(*eightCounter)
	*e = append(*e, v)
}

func (e *EightCounterList) Pop() interface{} {
	v := (*e)[len(*e)-1]
	*e = (*e)[:len(*e)-1]
	return v
}

func eight(nums []int, k int) []int {
	count := make(map[int]int)
	for _, v := range nums {
		count[v]++
	}
	var e EightCounterList
	for i, v := range count {
		temp := &eightCounter{Value: i, Count: v}
		heap.Push(&e, temp)
	}
	var result []int
	for len(result) < k {
		v := heap.Pop(&e).(*eightCounter)
		result = append(result, v.Value)
	}
	return result
}

type nineCounter struct {
	Val, Count int
}

type nineList []nineCounter

func (n nineList) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
func (n nineList) Less(i, j int) bool {
	return n[i].Count > n[j].Count
}
func (n nineList) Len() int {
	return len(n)
}

func (n *nineList) Push(x interface{}) {
	tag := x.(nineCounter)
	*n = append(*n, tag)
}

func (n *nineList) Pop() interface{} {
	length := len(*n)
	tag := (*n)[length-1]
	*n = (*n)[:length-1]
	return tag
}

func nine(nums []int, k int) []int {
	var result []int
	dic := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		dic[nums[i]]++
	}
	var l nineList
	for i, c := range dic {
		tag := nineCounter{Val: i, Count: c}
		heap.Push(&l, tag)
	}

	for len(result) < k {
		tag := heap.Pop(&l).(nineCounter)
		result = append(result, tag.Val)
	}
	return result
}

type tenCounter struct {
	Val, Count int
}

type tenList []tenCounter

func (t tenList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
func (t tenList) Less(i, j int) bool {
	return t[i].Count > t[j].Count
}
func (t tenList) Len() int {
	return len(t)
}
func (t *tenList) Push(x interface{}) {
	tag := x.(tenCounter)
	*t = append(*t, tag)
}
func (t *tenList) Pop() interface{} {
	length := len(*t)
	tag := (*t)[length-1]
	*t = (*t)[:length-1]
	return tag
}

func ten(nums []int, k int) []int {
	var result []int
	dic := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		dic[nums[i]]++
	}
	var t tenList
	for i, c := range dic {
		tag := tenCounter{Val: i, Count: c}
		heap.Push(&t, tag)
	}

	for len(result) < k {
		tag := heap.Pop(&t).(tenCounter)
		result = append(result, tag.Val)
	}
	return result
}

type elevenCounter struct {
	Val, Count int
}

type elevenList []elevenCounter

func (t elevenList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t elevenList) Less(i, j int) bool {
	return t[i].Count > t[j].Count
}
func (t elevenList) Len() int {
	return len(t)
}

func (t *elevenList) Push(v interface{}) {
	tag := v.(elevenCounter)
	*t = append(*t, tag)
}

func (t *elevenList) Pop() interface{} {
	length := len(*t)
	tag := (*t)[length-1]
	*t = (*t)[:length-1]
	return tag
}

func eleven(nums []int, k int) []int {
	var result []int
	dic := make(map[int]int)
	for _, num := range nums {
		dic[num]++
	}
	var t elevenList
	for num, count := range dic {
		tag := elevenCounter{Val: num, Count: count}
		heap.Push(&t, tag)
	}
	for len(result) < k {
		tag := heap.Pop(&t).(elevenCounter)
		result = append(result, tag.Val)
	}
	return result
}

type twelveCounter struct {
	Val, Count int
}

type twelveList []twelveCounter

func (t twelveList) Less(i, j int) bool {
	return t[i].Count > t[j].Count
}

func (t twelveList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t twelveList) Len() int {
	return len(t)
}

func (t *twelveList) Push(v interface{}) {
	tag := v.(twelveCounter)
	*t = append(*t, tag)
}

func (t *twelveList) Pop() interface{} {
	tag := (*t)[len(*t)-1]
	*t = (*t)[:len(*t)-1]
	return tag
}

func twelve(nums []int, k int) []int {
	var result []int
	var dic = make(map[int]int)
	for _, num := range nums {
		dic[num]++
	}
	var listT twelveList
	for num, count := range dic {
		tag := twelveCounter{Val: num, Count: count}
		heap.Push(&listT, tag)
	}
	for len(result) < k {
		top := heap.Pop(&listT).(twelveCounter)
		result = append(result, top.Val)
	}
	return result
}

type thirteenCounter struct {
	Val, Count int
}

type thirteenList []thirteenCounter

func (t thirteenList) Less(i, j int) bool {
	return t[i].Count > t[j].Count
}

func (t thirteenList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t thirteenList) Len() int {
	return len(t)
}

func (t *thirteenList) Push(v interface{}) {
	tag := v.(thirteenCounter)
	*t = append(*t, tag)
}

func (t *thirteenList) Pop() interface{} {
	tag := (*t)[len(*t)-1]
	*t = (*t)[:len(*t)-1]
	return tag
}

func thirteen(nums []int, k int) []int {
	var result []int
	var dic = make(map[int]int)
	for _, v := range nums {
		dic[v]++
	}
	var listT thirteenList
	for num, count := range dic {
		tag := thirteenCounter{Val: num, Count: count}
		heap.Push(&listT, tag)
	}
	for len(result) < k {
		top := heap.Pop(&listT).(thirteenCounter)
		result = append(result, top.Val)
	}
	return result
}
