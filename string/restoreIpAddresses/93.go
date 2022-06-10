package main

import (
	"fmt"
	"strconv"
)

var (
	res      []string
	segments []int
)

const SEG_COUNT = 4

func restoreIpAddresses(s string) []string {
	if len(s) == 0 {
		return []string{}
	}
	segments = make([]int, SEG_COUNT)
	res = []string{}
	dfs(s, 0, 0)
	return res
}

func dfs(s string, segId int, segStart int) {
	//如果找到了4端地址，并且s用完了 就是一个答案
	if segId == SEG_COUNT {
		if segStart == len(s) {
			addr := ""
			for i := 0; i < SEG_COUNT; i++ {
				addr = addr + strconv.Itoa(segments[i])
				if i < SEG_COUNT-1 {
					addr = addr + "."
				}
			}
			res = append(res, addr)

		} else {
			return
		}
	}
	//找完了，不符合，退出
	if segStart == len(s) {
		return
	}

	//0开头
	if s[segStart] == '0' {
		segments[segId] = 0
		dfs(s, segId+1, segStart+1)
	}

	//枚举每一种可能
	addr := 0
	for segEnd := segStart; segEnd < len(s); segEnd++ {
		addr = addr*10 + int(s[segEnd]-'0')
		if addr > 0 && addr <= 0xFF {
			segments[segId] = addr
			dfs(s, segId+1, segEnd+1)
		} else {
			break
		}
	}
}

const seg = 4

var nums = make([]int, seg)
var result = []string{}

func test(s string) []string {
	if len(s) == 0 {
		return []string{}
	}
	helper(s, 0, 0)
	return result
}

func helper(s string, sgeNum int, start int) {

	if sgeNum == seg {
		if start == len(s) {
			add := ""
			for i, num := range nums {
				add = add + strconv.Itoa(num)
				if i < seg-1 {
					add = add + "."
				}
			}
			result = append(result, add)
		} else {
			return
		}
	}
	if start == len(s) {
		return
	}

	if s[start] == '0' {
		nums[sgeNum] = 0
		helper(s, sgeNum+1, start+1)
	}

	address := 0

	for i := start; i < len(s); i++ {
		address = address*10 + int(s[i]-'0')
		if address > 0 && address <= 0xFF {
			nums[sgeNum] = address
			helper(s, sgeNum+1, i+1)
		} else {
			break
		}
	}
}

const segNum = 4

var resArr []string

func get(s string) []string {
	getter(s, 0, 0)
	return resArr
}

func getter(s string, segNum int, start int) {
	if segNum == seg {
		if start == len(s) {
			add := ""
			for i, num := range nums {
				add = add + strconv.Itoa(num)
				if i < seg {
					add += "."
				}
			}
			resArr = append(resArr, add)
		} else {
			return
		}
	}
	if start == len(s) {
		return
	}

	if s[start] == '0' {
		nums[segNum] = 0
		getter(s, segNum+1, start+1)
	}
	address := 0
	for i := start; i < len(s); i++ {
		address = address*10 + int(s[i]-'0')
		if address > 0 && address <= 0xFF {
			nums[segNum] = address
			getter(s, segNum+1, start+1)
		} else {
			break
		}
	}
}

func main() {
	test := "0000"
	fmt.Println(restoreIpAddresses(test))
}
