package main

import (
	"fmt"
)

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

func four(s string) string {
	begin := -1
	maxLen := 0
	var handler func(string, int, int) int
	handler = func(str string, i int, j int) int {
		for i >= 0 && j < len(str) {
			if str[i] != str[j] {
				break
			}
			i--
			j++
		}
		return j - i - 1
	}
	for i := 0; i < len(s); i++ {
		oddLen := handler(s, i, i)
		evenLen := handler(s, i, i+1)
		current := max(oddLen, evenLen)
		if current > maxLen {
			maxLen = current
			begin = i - (current-1)/2
		}
	}
	return s[begin : begin+maxLen]
}

func five(s string) string {
	if len(s) < 2 {
		return s
	}
	begin := -1
	maxLen := 0
	var handler func(string, int, int) int
	handler = func(str string, i int, j int) int {
		for i >= 0 && j < len(str) {
			if str[i] != str[j] {
				break
			}
			i--
			j++
		}
		return j - i - 1
	}
	for i := 0; i < len(s)-1; i++ {
		oddLen := handler(s, i, i)
		evenLen := handler(s, i, i+1)
		current := max(oddLen, evenLen)
		if current > maxLen {
			maxLen = current
			begin = i - (current-1)/2
		}
	}
	return s[begin : begin+maxLen]
}

func six(s string) string {
	if len(s) < 2 {
		return s
	}
	start, end := 0, 0
	handler := func(str string, i int, j int) (int, int) {
		for i >= 0 && j < len(str) {
			if str[i] != str[j] {
				break
			}
			i--
			j++
		}
		return i + 1, j - 1
	}
	for i := 0; i < len(s)-1; i++ {
		left1, right1 := handler(s, i, i)
		left2, right2 := handler(s, i, i+1)
		if right1-left1 > end-start {
			start = left1
			end = right1
		}

		if right2-left2 > end-start {
			start = left2
			end = right2
		}

	}
	return s[start : end+1]
}

func seven(s string) string {
	if len(s) < 2 {
		return s
	}
	start, end := 0, 0
	handler := func(str string, i int, j int) (int, int) {
		for i >= 0 && j < len(str) {
			if str[i] != str[j] {
				break
			}
			i--
			j++
		}
		return i + 1, j - 1
	}
	for i := 0; i < len(s)-1; i++ {
		left1, right1 := handler(s, i, i)
		left2, right2 := handler(s, i, i+1)
		if right1-left1 > end-start {
			start = left1
			end = right1
		}
		if right2-left2 > end-start {
			start = left2
			end = right2
		}
	}
	return s[start : end+1]
}

func eight(s string) string {
	if len(s) < 2 {
		return s
	}
	handler := func(str string, i, j int) (int, int) {
		for i >= 0 && j < len(str) {
			if str[i] != str[j] {
				break
			}
			i--
			j++
		}
		return i + 1, j - 1
	}
	start, end := 0, 0
	for i := 0; i < len(s)-1; i++ {
		left1, right1 := handler(s, i, i)
		left2, right2 := handler(s, i, i+1)
		if right1-left1 > end-start {
			start = left1
			end = right1
		}
		if right2-left2 > end-start {
			start = left2
			end = right2
		}
	}
	return s[start : end+1]
}

func nine(s string) string {
	if len(s) < 2 {
		return s
	}
	handler := func(str string, i, j int) (int, int) {
		for i >= 0 && j < len(str) {
			if str[i] != str[j] {
				break
			}
			i--
			j++
		}
		return i + 1, j - 1
	}
	start, end := 0, 0
	for i := 0; i < len(s)-1; i++ {
		left1, right1 := handler(s, i, i)
		left2, right2 := handler(s, i, i+1)
		if right1-left1 > end-start {
			start = left1
			end = right1
		}
		if right2-left2 > end-start {
			start = left2
			end = right2
		}
	}
	return s[start : end+1]
}

func ten(s string) string {
	if len(s) < 2 {
		return s
	}
	handler := func(str string, i, j int) (int, int) {
		for i >= 0 && j < len(str) {
			if str[i] != str[j] {
				break
			}
			i--
			j++
		}
		return i + 1, j - 1
	}

	start, end := 0, 0
	for i := 0; i < len(s)-1; i++ {
		left1, right1 := handler(s, i, i)
		left2, right2 := handler(s, i, i+1)
		if right1-left1 > end-start {
			start = left1
			end = right1
		}
		if right2-left2 > end-start {
			start = left2
			end = right2
		}
	}
	return s[start : end+1]
}

func eleven(s string) string {
	if len(s) < 2 {
		return s
	}
	handler := func(str string, i, j int) (int, int) {
		for i >= 0 && j < len(s) {
			if str[i] != str[j] {
				break
			}
			i--
			j++
		}
		return i + 1, j - 1
	}

	start, end := 0, 0
	for i := 0; i < len(s)-1; i++ {
		left1, right1 := handler(s, i, i)
		left2, right2 := handler(s, i, i+1)
		if right1-left1 > end-start {
			start = left1
			end = right1
		}
		if right2-left2 > end-start {
			start = left2
			end = right2
		}
	}
	return s[start : end+1]
}

func twelve(s string) string {
	if len(s) < 2 {
		return s
	}
	handler := func(str string, i, j int) (int, int) {
		for i >= 0 && j < len(str) {
			if str[i] != str[j] {
				break
			}
			i--
			j++
		}
		return i + 1, j - 1
	}
	start, end := 0, 0
	for i := 0; i < len(s)-1; i++ {
		left, right := handler(s, i, i)
		if right-left > end-start {
			start = left
			end = right
		}
		left, right = handler(s, i, i+1)
		if right-left > end-start {
			start = left
			end = right
		}
	}
	return s[start : end+1]
}

func thirteen(s string) string {
	if len(s) < 2 {
		return s
	}
	handler := func(str string, i, j int) (int, int) {
		for i >= 0 && j < len(str) {
			if str[i] != str[j] {
				break
			}
			i--
			j++
		}
		return i + 1, j - 1
	}
	start, end := 0, 0
	for i := 0; i < len(s)-1; i++ {
		left, right := handler(s, i, i)
		if right-left > end-start {
			start = left
			end = right
		}
		left, right = handler(s, i, i+1)
		if right-left > end-start {
			start = left
			end = right
		}
	}
	return s[start : end+1]
}

func main() {
	test := "ab"
	fmt.Println(longestPalindrome(test))
}
