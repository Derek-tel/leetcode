package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return ""
	}
	var sFreq, tFreq [256]int
	result, left, right, finalLeft, minWin, finalRight, count := "", 0, -1, -1, len(s)+1, -1, 0
	for i := 0; i < len(t); i++ {
		tFreq[t[i]-'a']++
	}

	for left < len(s) {
		if right+1 < len(s) && count < len(t) {
			sFreq[s[right+1]-'a']++
			if sFreq[s[right+1]-'a'] <= tFreq[s[right+1]-'a'] {
				count++
			}
			right++
		} else {
			if right-left+1 < minWin && count == len(t) {
				finalRight = right
				finalLeft = left
				minWin = right - left + 1
			}
			if sFreq[s[left]-'a'] == tFreq[s[left]-'a'] {
				count--
			}
			sFreq[s[left]-'a']--
			left++
		}
	}
	if finalLeft != -1 {
		result = string(s[finalLeft : finalRight+1])
	}

	return result

}

func test(s string, t string) string {
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return ""
	}
	var tFreq [256]int
	var sFreq [256]int
	for i := 0; i < len(t); i++ {
		tFreq[t[i]-'a']++
	}
	left, right := 0, 0
	finalLeft, finalRight := -1, -1
	count := 0
	minWin := math.MaxInt
	for left < len(s) {
		if right < len(s) && count < len(t) {
			sFreq[s[right]-'a']++
			if sFreq[s[right]-'a'] <= tFreq[s[right]-'a'] {
				count++
			}
			right++
		} else {
			if right-left < minWin && count == len(t) {
				finalLeft = left
				finalRight = right
				minWin = right - left
			}

			if sFreq[s[left]-'a'] == tFreq[s[left]-'a'] {
				count++
			}
			sFreq[s[left]-'a']--
			left++
		}
	}
	result := ""
	if finalRight != -1 {
		result = s[finalLeft : finalRight+1]
	}
	return result
}

func get(s, t string) string {
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return ""
	}
	var sFreq, tFreq [256]int
	for i := 0; i < len(t); i++ {
		tFreq[t[i]-'a']++
	}
	left, right := 0, 0
	finalLeft, finalRight := -1, -1
	count := 0
	minWin := math.MaxInt

	for left < len(s) {
		if right < len(s) && count < len(t) {
			sFreq[s[right]-'a']++
			if sFreq[s[right]-'a'] <= tFreq[s[right]-'a'] {
				count++
			}
			right++
		} else {
			if right-left < minWin && count == len(t) {
				finalLeft = left
				finalRight = right
				minWin = right - left
			}

			if sFreq[s[left]-'a'] <= tFreq[s[left]-'a'] {
				count--
			}
			sFreq[s[left]-'a']--
			left++
		}
	}
	result := ""
	if finalRight != -1 {
		result = s[finalLeft:finalRight]
	}
	return result
}

func main() {
	s := "aaaacddDAddddb"
	t := "adA"
	fmt.Println(minWindow(s, t))
}
