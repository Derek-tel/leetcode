package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	MaxLen := 0
	begin := -1
	for i := 0; i < len(s)-1; i++ {
		oddLen := expandString(s, i, i)
		evenLen := expandString(s, i, i+1)
		CurrLen := max(oddLen, evenLen)
		if CurrLen > MaxLen {
			MaxLen = CurrLen
			begin = i - (CurrLen-1)/2
		}
	}
	return s[begin : begin+MaxLen]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func test(s string) string {
	if len(s) < 2 {
		return s
	}
	begin := -1
	maxLen := 0

	for i := 0; i < len(s)-1; i++ {
		oddLen := expandString(s, i, i)
		evenLen := expandString(s, i, i+1)
		currLen := max(oddLen, evenLen)
		if currLen > maxLen {
			maxLen = currLen
			begin = i - (currLen-1)/2
		}
	}
	return s[begin : begin+maxLen]
}

func expandString(s string, i, j int) int {

	for i >= 0 && j < len(s) {
		if s[i] != s[j] {
			break
		}
		i--
		j++
	}
	return j - i - 1
}

func get(s string) string {
	begin := -1
	maxLen := 0
	for i := 0; i < len(s)-1; i++ {
		oddLen := expandString(s, i, i)
		evenLen := expandString(s, i, i+1)
		currLen := max(oddLen, evenLen)
		if currLen > maxLen {
			maxLen = currLen
			begin = i - (currLen-1)/2
		}
	}
	return s[begin : begin+maxLen]
}

func helper(s string, i, j int) int {
	for i >= 0 && j < len(s) {
		if s[i] != s[j] {
			break
		}
		i--
		j++
	}
	return j - i - 1
}

func main() {
	test := "ab"
	fmt.Println(longestPalindrome(test))
}
