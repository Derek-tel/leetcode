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

			if sFreq[s[left]-'a'] <= tFreq[s[left]-'a'] {
				count++
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

func four(s, t string) string {
	if len(s) == 0 || len(s) == 0 || len(s) < len(t) {
		return ""
	}
	var freqS, freqT [256]int
	for i := 0; i < len(t); i++ {
		freqT[t[i]-'a']++
	}
	left, right := 0, 0
	count := 0
	minWin := math.MaxInt
	finalLeft, finalRight := -1, -1
	for left < len(s) {
		if right < len(s) && count < len(t) {
			freqS[s[right]-'a']++
			if freqS[s[right]-'a'] <= freqT[s[right]-'a'] {
				count++
			}
			right++
		} else {
			if count == len(t) && right-left+1 <= minWin {
				fmt.Println(left, right)
				minWin = right - left + 1
				finalLeft = left
				finalRight = right
			}
			if freqS[s[left]-'a'] <= freqT[s[left]-'a'] {
				count--
			}
			freqS[s[left]-'a']--
			left++
		}
	}
	var result string
	if finalRight != -1 {
		result = s[finalLeft:finalRight]
	}
	return result
}

func five(s string, t string) string {
	var result string
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return result
	}
	var freqS, freqT [256]int
	for i := 0; i < len(t); i++ {
		freqT[t[i]-'a']++
	}
	left, right := 0, 0
	count := 0
	minLen := math.MaxInt
	finalLeft, finalRight := -1, -1
	for left < len(s) {
		if right < len(s) && count < len(t) {
			freqS[s[right]-'a']++
			if freqS[s[right]-'a'] <= freqT[s[right]-'a'] {
				count++
			}
			right++
		} else {
			if count == len(t) && right-left+1 < minLen {
				finalLeft = left
				finalRight = right
				minLen = right - left + 1
			}
			if freqS[s[left]-'a'] <= freqT[s[left]-'a'] {
				count--
			}
			freqS[s[left]-'a']--
			left++
		}
	}
	if finalRight != -1 {
		result = s[finalLeft:finalRight]
	}
	return result
}

func six(s string, t string) string {
	var result string
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return result
	}
	var freqS, freqT [256]int
	for i := 0; i < len(t); i++ {
		freqT[t[i]-'a']++
	}
	left, right := 0, 0
	count := 0
	minLen := math.MaxInt
	finalLeft, finalRight := -1, -1
	for left < len(s) {
		if right < len(s) && count < len(t) {
			freqS[s[right]-'a']++
			if freqS[s[right]-'a'] <= freqT[s[right]-'a'] {
				count++
			}
			right++
		} else {
			if count == len(t) && right-left+1 < minLen {
				finalLeft = left
				finalRight = right
				minLen = right - left + 1
			}
			if freqS[s[left]-'a'] <= freqT[s[left]-'a'] {
				count--
			}
			freqS[s[left]-'a']--
			left++
		}
	}
	if finalRight != -1 {
		result = s[finalLeft:finalRight]
	}
	return result
}

func seven(s string, t string) string {
	var result string
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return result
	}
	var freqS, freqT [256]int
	for i := 0; i < len(t); i++ {
		freqT[t[i]-'a']++
	}
	left, right, count, minLen, finalLeft, finalRight := 0, 0, 0, math.MaxInt, -1, -1
	for left < len(s) {
		if right < len(s) && count < len(t) {
			freqS[s[right]-'a']++
			if freqS[s[right]-'a'] <= freqT[s[right]-'a'] {
				count++
			}
			right++
		} else {
			if count == len(t) && right-left+1 < minLen {
				finalLeft = left
				finalRight = right
				minLen = right - left + 1
			}
			if freqS[s[left]-'a'] <= freqT[s[left]-'a'] {
				count--
			}
			freqS[s[left]-'a']--
			left++
		}
	}
	if finalRight != -1 {
		result = s[finalLeft:finalRight]
	}
	return result
}

func eight(s string, t string) string {
	var result string
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return result
	}
	var freqS, freqT [256]int
	for i := 0; i < len(t); i++ {
		freqT[t[i]-'a']++
	}
	left, right, count, finalLeft, finalRight, minLen := 0, 0, 0, -1, -1, math.MaxInt
	for left < len(s) {
		if right < len(s) && count < len(t) {
			freqS[s[right]-'a']++
			if freqS[s[right]-'a'] <= freqT[s[right]-'a'] {
				count++
			}
			right++
		} else {
			if count == len(t) && right-left+1 < minLen {
				finalLeft = left
				finalRight = right
				minLen = right - left + 1
			}
			if freqS[s[left]-'a'] <= freqT[s[left]-'a'] {
				count--
			}
			freqS[s[left]-'a']--
			left++
		}
	}
	if finalRight != -1 {
		result = s[finalLeft:finalRight]
	}
	return result
}

func nine(s string, t string) string {
	var result string
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return result
	}
	var freqS, freqT [256]int
	for i := 0; i < len(t); i++ {
		freqT[t[i]-'a']++
	}
	left, right, count := 0, 0, 0
	finalLeft, finalRight, minLen := -1, -1, math.MaxInt
	for left < len(s) {
		if right < len(s) && count < len(t) {
			freqS[s[right]-'a']++
			if freqS[s[right]-'a'] <= freqT[s[right]-'a'] {
				count++
			}
			right++
		} else {
			if count == len(t) && right-left+1 < minLen {
				finalLeft = left
				finalRight = right
				minLen = right - left + 1
			}
			if freqS[s[left]-'a'] <= freqT[s[left]-'a'] {
				count--
			}
			freqS[s[left]-'a']--
			left++
		}
	}
	if finalRight != -1 {
		result = s[finalLeft:finalRight]
	}
	return result
}

func ten(s string, t string) string {
	var result string
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return result
	}
	var freqS, freqT [256]int
	for i := 0; i < len(t); i++ {
		freqT[t[i]-'a']++
	}
	left, right, count := 0, 0, 0
	finalLeft, finalRight, minLen := -1, -1, math.MaxInt
	for left < len(s) {
		if right < len(s) && count < len(t) {
			freqS[s[right]-'a']++
			if freqS[s[right]-'a'] <= freqT[s[right]-'a'] {
				count++
			}
			right++
		} else {
			if count == len(t) && right-left+1 < minLen {
				finalLeft = left
				finalRight = right
				minLen = right - left + 1
			}
			if freqS[s[left]-'a'] <= freqT[s[left]-'a'] {
				count--
			}
			freqS[s[left]-'a']--
			left++
		}
	}
	if finalRight != -1 {
		result = s[finalLeft:finalRight]
	}
	return result
}

func eleven(s string, t string) string {
	var result string
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return result
	}
	var freqS, freqT [256]int
	for i := 0; i < len(t); i++ {
		freqT[t[i]-'a']++
	}
	left, right, count := 0, 0, 0
	finalLeft, finalRight, minLen := -1, -1, math.MaxInt
	for left < len(s) {
		if right < len(s) && count < len(t) {
			freqS[s[right]-'a']++
			if freqS[s[right]-'a'] <= freqT[s[right]-'a'] {
				count++
			}
			right++
		} else {
			if count == len(t) && right-left+1 < minLen {
				finalLeft = left
				finalRight = right
				minLen = right - left + 1
			}
			if freqS[s[left]-'a'] <= freqT[s[left]-'a'] {
				count--
			}
			freqS[s[left]-'a']--
			left++
		}
	}
	if finalRight != -1 {
		result = s[finalLeft:finalRight]
	}
	return result
}

func twelve(s string, t string) string {
	var result string
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return result
	}
	var freqS, freqT [256]int
	for i := 0; i < len(t); i++ {
		freqT[t[i]-'a']++
	}
	left, right, count := 0, 0, 0
	start, end, minLen := -1, -1, math.MaxInt
	for left < len(s) {
		if right < len(s) && count < len(t) {
			freqS[s[right]-'a']++
			if freqS[s[right]-'a'] <= freqT[s[right]-'a'] {
				count++
			}
			right++
		} else {
			if count == len(t) && right-left+1 < minLen {
				start = left
				end = right
				minLen = right - left + 1
			}
			if freqS[s[left]-'a'] <= freqT[s[left]-'a'] {
				count--
			}
			freqS[s[left]-'a']--
			left++
		}
	}
	if end != -1 {
		result = s[start:end]
	}
	return result
}

func thirteen(s string, t string) string {
	var result string
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return result
	}
	var freqS, freqT [256]int
	for i := 0; i < len(t); i++ {
		freqT[t[i]-'a']++
	}
	left, right, count := 0, 0, 0
	start, end, minLen := -1, -1, math.MaxInt
	for left < len(s) {
		if right < len(s) && count < len(t) {
			freqS[s[right]-'a']++
			if freqS[s[right]-'a'] <= freqT[s[right]-'a'] {
				count++
			}
			right++
		} else {
			if count == len(t) && right-left+1 < minLen {
				start = left
				end = right
				minLen = right - left + 1
			}
			if freqS[s[left]-'a'] <= freqT[s[left]-'a'] {
				count--
			}
			freqS[s[left]-'a']--
			left++
		}
	}
	if end != -1 {
		result = s[start:end]
	}
	return result
}

func fourteen(s string, t string) string {
	var result string
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return result
	}
	var freqS, freqT [256]int
	for i := 0; i < len(t); i++ {
		freqT[t[i]-'a']++
	}
	left, right, count := 0, 0, 0
	start, end, minLen := -1, -1, math.MaxInt
	for left < len(s) {
		if right < len(s) && count < len(t) {
			freqS[s[right]-'a']++
			if freqS[s[right]-'a'] <= freqT[s[right]-'a'] {
				count++
			}
			right++
		} else {
			if count == len(t) && right-left < minLen {
				start = left
				end = right
				minLen = right - left
			}
			if freqS[s[left]-'a'] <= freqT[s[left]-'a'] {
				count--
			}
			freqS[s[left]-'a']--
			left++
		}
	}
	if end != -1 {
		result = s[start:end]
	}
	return result
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(eleven(s, t))
}
