package main

import (
	"fmt"
)

func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}
	result := []int{}
	counter := map[string]int{}
	for _, w := range words {
		counter[w]++
	}
	length, totalLength, tmpCounter := len(words[0]), len(words[0])*len(words), copyMap(counter)
	for i, start := 0, 0; i < len(s)-length+1 && start < len(s)-length+1; i++ {
		if tmpCounter[s[i:i+length]] >= 1 {
			tmpCounter[s[i:i+length]]--
			if check(tmpCounter) && i+length-start == totalLength {
				result = append(result, start)
				continue
			}
			i = i + length - 1
		} else {
			start++
			i = start - 1
			tmpCounter = copyMap(counter)
		}
	}
	return result
}

func check(s map[string]int) bool {
	flag := true
	for _, v := range s {
		if v > 0 {
			flag = false
			break
		}
	}
	return flag
}

func copyMap(s map[string]int) map[string]int {
	c := map[string]int{}
	for k, v := range s {
		c[k] = v
	}
	return c
}

func testFunc(s string, words []string) []int {
	length := len(s)
	dic := make(map[string]int)

	for _, word := range words {
		dic[word]++
	}
	wordLen := len(words[0])
	totalLen := wordLen * len(words)
	tempDic := copyMap(dic)
	var res []int
	for i, start := 0, 0; i < length-wordLen+1 && start < length-wordLen+1; i++ {
		if tempDic[s[i:i+length]] > 0 {
			tempDic[s[i:i+length]]--
			if check(tempDic) && i+length-start == totalLen {
				res = append(res, start)
				continue
			}
			i = i + length - 1
		} else {
			start++
			i = start - 1
			tempDic = copyMap(dic)
		}
	}
	return res
}

func get(s string, words []string) []int {
	length := len(s)
	dic := make(map[string]int)
	for _, word := range words {
		dic[word]++
	}
	wordLen := len(words[0])
	totalLen := wordLen * len(words)
	tempDic := copyMap(dic)
	var res []int
	for i, start := 0, 0; i+wordLen-1 < length && start+wordLen-1 < length; i++ {
		if tempDic[s[i:i+wordLen]] > 0 {
			tempDic[s[i:i+wordLen]]--
			if i+wordLen-start == totalLen && check(tempDic) {
				res = append(res, start)
				continue
			}
			i = i + wordLen - 1
		} else {
			start++
			i = start - 1
			tempDic = copyMap(dic)
		}
	}
	return res
}

func helper(s string, words []string) []int {
	wordLen := len(words[0])
	wordCount := len(words)
	length := len(s)
	if wordCount < 1 || length < wordLen*wordCount {
		return []int{}
	}
	var res []int
	dic := make(map[string]int)
	for _, v := range words {
		dic[v]++
	}
	for i := 0; i < wordLen; i++ {
		var count int
		var tempDic = make(map[string]int)
		for left, right := i, i; right <= length-wordLen; right += wordLen {
			word := s[right : right+wordLen]
			if num, ok := dic[word]; ok {
				for tempDic[word] >= num {
					tempDic[s[left:left+wordLen]]--
					count--
					left += wordLen
				}
				tempDic[word]++
				count++
			} else {
				left = right + wordLen
				count = 0
				tempDic = make(map[string]int)
			}
			if count == wordCount {
				res = append(res, left)
			}
		}
	}
	return res
}

func five(s string, words []string) []int {
	var result []int
	wordLen := len(words[0])
	wordCount := len(words)
	if wordCount < 1 || len(s) < wordLen*wordCount {
		return result
	}
	dic := make(map[string]int)
	for _, word := range words {
		dic[word]++
	}
	for i := 0; i < wordLen; i++ {
		var count int
		var tempDic = make(map[string]int)
		for left, right := i, i; right <= len(s)-wordLen; right += wordLen {
			word := s[right : right+wordLen]
			if num, ok := dic[word]; ok {
				for tempDic[word] >= num {
					tempDic[s[left:left+wordLen]]--
					count--
					left = left + wordLen
				}
				tempDic[word]++
				count++
			} else {
				left = right + wordLen
				count = 0
				tempDic = make(map[string]int)
			}
			if count == wordCount {
				result = append(result, left)
			}
		}
	}
	return result
}

func six(s string, words []string) []int {
	var result []int
	wordLen := len(words[0])
	wordCount := len(words)
	if wordCount < 1 || len(s) < wordLen*wordCount {
		return result
	}
	dic := make(map[string]int)
	for _, word := range words {
		dic[word]++
	}
	for i := 0; i < wordLen; i++ {
		count := 0
		var tempDic = make(map[string]int)
		for left, right := i, i; right <= len(s)-wordLen; right += wordLen {
			word := s[right : right+wordLen]
			if num, ok := dic[word]; ok {
				for tempDic[word] >= num {
					tempDic[s[left:left+wordLen]]--
					count--
					left = left + wordLen
				}
				tempDic[word]++
				count++
			} else {
				left = right + wordLen
				count = 0
				tempDic = make(map[string]int)
			}
			if count == wordCount {
				result = append(result, left)
			}
		}
	}
	return result
}

func seven(s string, words []string) []int {
	var result []int
	wordLen := len(words[0])
	wordCount := len(words)
	if wordCount < 1 || len(s) < wordLen*wordCount {
		return result
	}
	dic := make(map[string]int)
	for _, word := range words {
		dic[word]++
	}
	for i := 0; i < wordLen; i++ {
		count := 0
		var tempDic = make(map[string]int)
		for left, right := i, i; right <= len(s)-wordLen; right += wordLen {
			word := s[right : right+wordLen]
			if num, ok := dic[word]; ok {
				for tempDic[word] >= num {
					tempDic[s[left:left+wordLen]]--
					count--
					left = left + wordLen
				}
				tempDic[word]++
				count++
			} else {
				left = right + wordLen
				count = 0
				tempDic = make(map[string]int)
			}
			if count == wordCount {
				result = append(result, left)
			}
		}
	}
	return result
}

func eight(s string, words []string) []int {
	var result []int
	wordLen, wordCount := len(words[0]), len(words)
	if wordCount < 1 || len(s) < wordLen*wordCount {
		return result
	}
	dic := make(map[string]int)
	for _, word := range words {
		dic[word]++
	}
	for i := 0; i < wordLen; i++ {
		count := 0
		var temp = make(map[string]int)
		for left, right := i, i; right <= len(s)-wordLen; right += wordLen {
			word := s[right : right+wordLen]
			if num, ok := dic[word]; ok {
				for temp[word] >= num {
					temp[s[left:left+wordLen]]--
					count--
					left = left + wordLen
				}
				temp[word]++
				count++
			} else {
				left = right + wordLen
				count = 0
				temp = make(map[string]int)
			}
			if count == wordCount {
				result = append(result, left)
			}
		}
	}
	return result
}

func nine(s string, words []string) []int {
	var result []int
	wordLen, wordCount := len(words[0]), len(words)
	if wordCount < 1 || len(s) < wordLen*wordCount {
		return result
	}
	dic := make(map[string]int)
	for _, word := range words {
		dic[word]++
	}
	for i := 0; i < wordLen; i++ {
		count := 0
		var temp = make(map[string]int)
		for left, right := i, i; right <= len(s)-wordLen; right += wordLen {
			word := s[right : right+wordLen]
			if num, ok := dic[word]; ok {
				for temp[word] >= num {
					temp[s[left:left+wordLen]]--
					count--
					left = left + wordLen
				}
				temp[word]++
				count++
			} else {
				left = right + wordLen
				count = 0
				temp = make(map[string]int)
			}
			if count == wordCount {
				result = append(result, left)
			}
		}
	}
	return result
}

func ten(s string, words []string) []int {
	var result []int
	wordLen, wordCount := len(words[0]), len(words)
	if wordCount < 1 || len(s) < wordLen*wordCount {
		return result
	}
	dic := make(map[string]int)
	for _, word := range words {
		dic[word]++
	}
	for i := 0; i < wordLen; i++ {
		count := 0
		temp := make(map[string]int)
		for left, right := i, i; right+wordLen <= len(s); right += wordLen {
			word := s[right : right+wordLen]
			if num, ok := dic[word]; ok {
				for temp[word] >= num {
					temp[s[left:left+wordLen]]--
					count--
					left = left + wordLen
				}
				temp[word]++
				count++
			} else {
				left = right + wordLen
				count = 0
				temp = make(map[string]int)
			}
			if count == wordCount {
				result = append(result, left)
			}
		}
	}
	return result
}

func eleven(s string, words []string) []int {
	var result []int
	wordLen, wordCount := len(words[0]), len(words)
	if wordCount < 1 || len(s) < wordLen*wordCount {
		return result
	}
	dic := make(map[string]int)
	for _, word := range words {
		dic[word]++
	}
	for i := 0; i < wordLen; i++ {
		temp := make(map[string]int)
		count := 0
		for left, right := i, i; right+wordLen <= len(s); right += wordLen {
			word := s[right : right+wordLen]
			if num, ok := dic[word]; ok {
				for temp[word] >= num {
					temp[s[left:left+wordLen]]--
					count--
					left = left + wordLen
				}
				temp[word]++
				count++
			} else {
				left = right + wordLen
				count = 0
				temp = make(map[string]int)
			}
			if count == wordCount {
				result = append(result, left)
			}
		}
	}
	return result
}

func twelve(s string, words []string) []int {
	var result []int
	wordLen, wordCount := len(words[0]), len(words)
	if wordCount < 1 || len(s) < wordLen*wordCount {
		return result
	}
	dic := make(map[string]int)
	for _, word := range words {
		dic[word]++
	}
	for i := 0; i < wordLen; i++ {
		temp := make(map[string]int)
		count := 0
		for left, right := i, i; right+wordLen <= len(s); right += wordLen {
			if num, ok := dic[s[right:right+wordLen]]; ok {
				for temp[s[right:right+wordLen]] >= num {
					temp[s[left:left+wordLen]]--
					count--
					left = left + wordLen
				}
				temp[s[right:right+wordLen]]++
				count++
			} else {
				left = right + wordLen
				count = 0
				temp = make(map[string]int)
			}
			if count == wordCount {
				result = append(result, left)
			}
		}
	}
	return result
}

func thirteen(s string, words []string) []int {
	var result []int
	wordLen, wordCount := len(words[0]), len(words)
	if wordCount < 1 || len(s) < wordLen*wordCount {
		return result
	}
	dic := make(map[string]int)
	for _, word := range words {
		dic[word]++
	}
	for i := 0; i < wordLen; i++ {
		temp := make(map[string]int)
		count := 0
		for left, right := i, i; right+wordLen <= len(s); right += wordLen {
			if num, ok := dic[s[right:right+wordLen]]; ok {
				for temp[s[right:right+wordLen]] >= num {
					temp[s[left:left+wordLen]]--
					count--
					left += wordLen
				}
				temp[s[right:right+wordLen]]++
				count++
			} else {
				temp = make(map[string]int)
				count = 0
				left = right + wordLen
			}
			if count == wordCount {
				result = append(result, left)
			}
		}
	}
	return result
}

func main() {
	test := "wordgoodgoodgoodbestword"
	words := []string{"word", "good", "best", "good"}
	fmt.Printf("result = %+v", findSubstring(test, words))
}
