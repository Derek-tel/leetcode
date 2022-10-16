package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [127]byte
	result, left, right := 0, 0, -1

	for left < len(s) {
		if right+1 < len(s) && freq[s[right+1]] == 0 {
			freq[s[right+1]]++
			right++
		} else {
			freq[s[left]]--
			left++
		}
		result = max(result, right-left+1)
	}
	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func test(s string) int {
	var freq [256]byte
	left, right := 0, 0
	res := 0
	for left < len(s) {
		if right < len(s) && freq[s[right]] == 0 {
			freq[s[right]]++
			right++
		} else {
			freq[s[left]]--
			left++
		}

		res = max(res, right-left)
	}
	return res
}

func get(s string) int {
	var freq [256]byte
	left, right := 0, 0
	res := 0
	for right < len(s) {
		if freq[s[right]] == 0 {
			freq[s[right]]++
			right++
		} else {
			freq[s[left]]--
			left++
		}
		res = max(res, right-left)
	}
	return res
}

func five(s string) int {
	var freq [256]int
	left, right := 0, 0
	var result int
	for right < len(s) {
		if freq[s[right]] == 0 {
			freq[s[right]]++
			right++
		} else {
			freq[s[left]]--
			left++
		}
		result = max(result, right-left)
	}
	return result
}

func main() {
	s := "abcbcbbb"
	fmt.Println(test(s))
}
