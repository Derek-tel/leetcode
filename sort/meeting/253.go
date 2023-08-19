package main

import (
	"fmt"
	"sort"
)

type Meet [][]int

func (m Meet) Less(i, j int) bool {
	return m[i][0] < m[j][0]
}

func (m Meet) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m Meet) Len() int {
	return len(m)
}

func meeting(interval [][]int) int {
	sort.Sort(Meet(interval))
	heap := make([]int, 0)

	for i := 0; i < len(interval); i++ {
		if len(heap) == 0 {
			heap = append(heap, interval[i][1])
		} else {
			if interval[i][0] > heap[0] { //本厂会议开始时间 早于  所有会议室里最小结束时间
				heap[0] = interval[i][1] //占上这个会议室
			} else {
				heap = append(heap, interval[i][1]) //新开一个吧
			}
		}
		buildMinHeap(heap, len(heap))
	}

	return len(heap)
}

func buildMinHeap(nums []int, heapSize int) {
	for i := heapSize / 2; i >= 0; i-- {
		heapify(nums, heapSize, i)
	}

}

func heapify(nums []int, heapSize int, i int) {
	left := 2*i + 1
	right := 2*i + 2
	small := i

	if left < heapSize && nums[i] > nums[left] {
		small = left
	}
	if right < heapSize && nums[i] > nums[right] {
		small = right
	}
	if small != i {
		nums[small], nums[i] = nums[i], nums[small]
		heapify(nums, heapSize, small)
	}
}

type rank [][]int

func (r rank) Less(i, j int) bool {
	return r[i][0] < r[j][0]
}

func (r rank) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r rank) Len() int {
	return len(r)
}

func meet(interval [][]int) int {
	sort.Sort(rank(interval))
	h := []int{}

	for i := 0; i < len(interval); i++ {
		if len(h) == 0 {
			h = append(h, interval[i][1])
		} else {
			if interval[i][0] > h[0] { //当前会议开始小时 比占用的会议室中最开释放的早，占用上
				h[0] = interval[i][1]
			} else {
				h = append(h, interval[i][1])
			}
		}

		heap(h)
	}
	return len(h)
}

func heap(nums []int) {
	heapSize := len(nums)
	for i := heapSize / 2; i >= 0; i-- {
		build(nums, i, heapSize)
	}
}

func build(nums []int, i int, size int) {
	left := i*2 + 1
	right := i*2 + 2
	small := i
	if left < size && nums[left] < nums[small] {
		small = left
	}
	if right < size && nums[right] < nums[small] {
		small = right
	}

	if small != i {
		nums[small], nums[i] = nums[i], nums[small]
		build(nums, small, size)
	}
}

type hack [][]int

func (h hack) Less(i, j int) bool {
	return h[i][0] < h[j][0]
}
func (h hack) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h hack) Len() int {
	return len(h)
}

func get(interval [][]int) int {
	sort.Sort(hack(interval))
	var ans []int

	for i := 0; i < len(interval); i++ {
		if len(ans) == 0 {
			ans = append(ans, interval[i][1])
		} else {
			if interval[i][0] < ans[0] {
				ans = append(ans, interval[i][1])
			} else {
				ans[0] = interval[i][1]
			}
		}
		helper(ans)
	}
	return len(ans)
}

func helper(nums []int) {
	size := len(nums)
	for i := size / 2; i >= 0; i-- {
		up(nums, i, size)
	}
}

func up(nums []int, i int, size int) {
	left := i<<1 + 1
	right := i<<2 + 2
	small := i
	if left < size && nums[left] < nums[small] {
		small = left
	}
	if right < size && nums[right] < nums[small] {
		small = right
	}
	if small != i {
		nums[small], nums[i] = nums[i], nums[small]
		up(nums, small, size)
	}
}

type demo [][]int

func (d demo) Less(i, j int) bool {
	return d[i][0] < d[j][0]
}

func (d demo) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
func (d demo) Len() int {
	return len(d)
}

func four(interval [][]int) int {
	sort.Sort(demo(interval))
	var buildHeap func([]int)
	var modifyHeap func([]int, int, int)
	buildHeap = func(ints []int) {
		size := len(ints)
		for i := size / 2; i >= 0; i-- {
			modifyHeap(ints, i, size)
		}
	}
	modifyHeap = func(ints []int, i int, size int) {
		left := i*2 + 1
		right := i*2 + 2
		small := i
		if left < size && ints[left] <= ints[small] {
			small = left
		}
		if right < size && ints[right] <= ints[small] {
			small = right
		}
		if small != i {
			ints[i], ints[small] = ints[small], ints[i]
			modifyHeap(ints, small, size)
		}
	}
	var result []int
	for i := 0; i < len(interval); i++ {
		if len(result) == 0 {
			result = append(result, interval[i][1])
		} else {
			if interval[i][0] < result[0] {
				result = append(result, interval[i][1])
			} else {
				result[0] = interval[i][1]
			}
		}
		buildHeap(result)
	}
	return len(result)
}

func five(interval [][]int) int {
	var heapModify func([]int, int, int)
	heapModify = func(ints []int, i int, size int) {
		left := i*2 + 1
		right := i*2 + 2
		smallest := i
		if left < size && ints[left] < ints[smallest] {
			smallest = left
		}
		if right < size && ints[right] < ints[smallest] {
			smallest = right
		}
		if i != smallest {
			ints[i], ints[smallest] = ints[smallest], ints[i]
			heapModify(ints, smallest, size)
		}
	}
	var buildHeap func([]int, int)
	buildHeap = func(ints []int, size int) {
		for i := size / 2; i >= 0; i-- {
			heapModify(ints, i, size)
		}
	}
	sort.Slice(interval, func(i, j int) bool {
		return interval[i][0] < interval[j][0]
	})
	var ends []int
	for i := 0; i < len(interval); i++ {
		if len(ends) == 0 {
			ends = append(ends, interval[i][1])
		} else {
			if interval[i][1] < ends[0] {
				ends = append(ends, interval[i][1])
			} else {
				ends[0] = interval[i][1]
			}
		}
		buildHeap(ends, len(ends))
	}
	return len(ends)
}

func six(interval [][]int) int {
	var heapModify func([]int, int, int)
	heapModify = func(ints []int, i int, size int) {
		left := i*2 + 1
		right := i*2 + 2
		smallest := i
		if left < size && ints[left] < ints[smallest] {
			smallest = left
		}
		if right < size && ints[right] < ints[smallest] {
			smallest = right
		}
		if smallest != i {
			ints[smallest], ints[i] = ints[i], ints[smallest]
			heapModify(ints, smallest, size)
		}
	}
	var buildHead func([]int, int)
	buildHead = func(ints []int, size int) {
		for i := len(ints) / 2; i >= 0; i-- {
			heapify(ints, i, size)
		}
	}
	sort.Slice(interval, func(i, j int) bool {
		return interval[i][0] < interval[j][0]
	})
	var ends []int
	for i := 0; i < len(interval); i++ {
		if len(ends) == 0 {
			ends = append(ends, interval[i][1])
		} else {
			if interval[i][1] < ends[0] {
				ends = append(ends, interval[i][1])
			} else {
				ends[0] = interval[i][1]
			}
		}
		buildHead(ends, len(ends))
	}
	return len(ends)
}

func seven(interval [][]int) int {
	sort.Slice(interval, func(i, j int) bool {
		return interval[i][0] < interval[j][0]
	})
	var ends []int
	for i := 0; i < len(interval); i++ {
		if len(ends) == 0 {
			ends = append(ends, interval[i][1])
		} else {
			if interval[i][0] < ends[0] {
				ends = append(ends, interval[i][1])
			} else {
				ends[0] = interval[i][1]
			}
		}
		heap(ends)
	}
	return len(ends)
}

func eight(interval [][]int) int {
	sort.Slice(interval, func(i, j int) bool {
		return interval[i][0] < interval[j][0]
	})
	var end []int
	for i := 0; i < len(interval); i++ {
		if len(end) == 0 {
			end = append(end, interval[i][1])
		} else {
			if interval[i][0] < end[0] {
				end = append(end, interval[i][1])
			} else {
				end[0] = interval[i][1]
			}
		}
		heap(end)
	}
	return len(end)
}

func nine(interval [][]int) int {
	sort.Slice(interval, func(i, j int) bool {
		return interval[i][0] < interval[j][0]
	})

	var end []int
	for i := 0; i < len(interval); i++ {
		if len(end) == 0 {
			end = append(end, interval[i][1])
		} else {
			if interval[i][0] < end[0] {
				end = append(end, interval[i][1])
			} else {
				end[0] = interval[i][1]
			}
		}
		heap(end)
	}
	return len(end)
}

func main() {

	test := [][]int{
		{0, 6},
		{9, 20},
		{5, 10},
	}
	fmt.Println(four(test))
}
