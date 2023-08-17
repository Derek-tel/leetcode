package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/**

n个包厢  长度 n 每个包厢大小数组 a []int    [3，4,7,9]

m组的客人   b 长度m  每组人数  b []int    [3，10，4，8]

m组的客人   c 长度m  每组花的钱  c []int  [13，12，5，5]

*/

var price []int
var numbers []int

type priceSort []int

func (p priceSort) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p priceSort) Less(i, j int) bool {
	return price[p[i]] > price[p[j]] || (price[p[i]] == price[p[j]] && numbers[p[i]] < numbers[p[j]])
}
func (p priceSort) Len() int {
	return len(p)
}

func (p *priceSort) Push(x interface{}) {
	v := x.(int)
	*p = append(*p, v)
}
func (p *priceSort) Pop() interface{} {
	v := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return v
}

func maxPrice(a []int, b []int, c []int) int {
	price = c
	numbers = b
	sort.Ints(a)

	sh := priceSort{}
	for i := 0; i < len(c); i++ {
		heap.Push(&sh, i)
	}

	visit := make([]bool, len(a))
	result := 0
	for len(sh) > 0 {
		top := heap.Pop(&sh).(int)

		p := c[top]
		num := b[top]

		for i := 0; i < len(a); i++ {
			if !visit[i] && a[i] >= num {
				result = result + p
				visit[i] = true
				break
			}
		}
	}

	return result
}

func main() {
	p := []int{3, 4, 7, 9}
	nums := []int{3, 10, 4, 8}
	prices := []int{13, 12, 5, 9}

	fmt.Println(maxPrice(p, nums, prices))
}
