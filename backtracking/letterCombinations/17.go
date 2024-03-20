package main

import (
	"fmt"
)

func letterCombinations(digits string) []string {
	if len(digits) < 1 {
		return []string{}
	}
	res := []string{}
	backTracking(digits, 0, "", &res)
	return res
}

func backTracking(digits string, index int, str string, res *[]string) {
	if index == len(digits) {
		*res = append(*res, str)
		return
	}
	digit := digits[index]
	chars := phoneMap[string(digit)]
	for _, char := range chars {
		backTracking(digits, index+1, str+string(char), res)
	}
}

func two(digits string) []string {
	if len(digits) < 1 {
		return []string{}
	}
	var result []string
	var helper func(string, int, string, *[]string)
	helper = func(letters string, index int, str string, res *[]string) {
		if index == len(letters) {
			*res = append(*res, str)
			return
		}
		digit := digits[index]
		chars := phoneMap[string(digit)]
		for _, char := range chars {
			helper(digits, index+1, str+string(char), res)
		}

	}
	helper(digits, 0, "", &result)
	return result
}

var phone map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func three(digits string) []string {
	if len(digits) < 1 {
		return []string{}
	}
	var result []string
	var handler func(int, string)
	handler = func(i int, s string) {
		if i == len(digits) {
			result = append(result, s)
			return
		}
		dig := digits[i]
		char := phone[string(dig)]
		for _, v := range char {
			handler(i+1, s+string(v))
		}
	}
	handler(0, "")
	return result
}

var letters map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func four(digits string) []string {
	var result []string
	if len(digits) < 1 {
		return result
	}
	var handler func(int, string)
	handler = func(i int, s string) {
		if i == len(digits) {
			result = append(result, s)
			return
		}
		dig := digits[i]
		char := letters[string(dig)]
		for _, v := range char {
			handler(i+1, s+string(v))
		}
	}
	handler(0, "")
	return result
}

func five(digits string) []string {
	var result []string
	if len(digits) < 1 {
		return result
	}
	var handler func(int, string)
	handler = func(index int, s string) {
		if index == len(digits) {
			result = append(result, s)
			return
		}
		dig := digits[index]
		char := letters[string(dig)]
		for _, v := range char {
			handler(index+1, s+string(v))
		}
	}
	handler(0, "")
	return result
}

func six(digits string) []string {
	var result []string
	if len(digits) < 1 {
		return result
	}
	var handler func(int, string)
	handler = func(index int, s string) {
		if index == len(digits) {
			result = append(result, s)
			return
		}
		dig := digits[index]
		char := letters[string(dig)]
		for _, ch := range char {
			handler(index+1, s+string(ch))
		}
	}
	handler(0, "")
	return result
}

func seven(digits string) []string {
	var result []string
	if len(digits) < 1 {
		return result
	}
	var handler func(int, string)
	handler = func(index int, str string) {
		if index == len(digits) {
			result = append(result, str)
			return
		}
		dig := digits[index]
		char := letters[string(dig)]
		for _, ch := range char {
			handler(index+1, str+string(ch))
		}
	}
	handler(0, "")
	return result
}

func eight(digits string) []string {
	var result []string
	if len(digits) < 1 {
		return result
	}
	var handler func(int, string)
	handler = func(index int, str string) {
		if index == len(digits) {
			result = append(result, str)
			return
		}
		dig := digits[index]
		char := letters[string(dig)]
		for _, ch := range char {
			handler(index+1, str+string(ch))
		}
	}
	handler(0, "")
	return result
}

func nine(digits string) []string {
	var result []string
	if len(digits) < 1 {
		return result
	}
	var handler func(int, string)
	handler = func(index int, str string) {
		if index == len(digits) {
			result = append(result, str)
			return
		}
		dig := digits[index]
		char := letters[string(dig)]
		for _, v := range char {
			handler(index+1, str+string(v))
		}
	}
	handler(0, "")
	return result
}

func ten(digits string) []string {
	var result []string
	if len(digits) < 1 {
		return result
	}
	var handler func(int, string)
	handler = func(index int, s string) {
		if index == len(digits) {
			result = append(result, s)
			return
		}
		ch := digits[index]
		char := letters[string(ch)]
		for _, v := range char {
			handler(index+1, s+string(v))
		}
	}
	handler(0, "")
	return result
}

func eleven(digits string) []string {
	var result []string
	if len(digits) < 1 {
		return result
	}
	var handler func(int, string)
	handler = func(index int, s string) {
		if index == len(digits) {
			result = append(result, s)
			return
		}
		ch := digits[index]
		char := letters[string(ch)]
		for _, v := range char {
			handler(index+1, s+string(v))
		}
	}
	handler(0, "")
	return result
}

func twelve(digits string) []string {
	var result []string
	if len(digits) < 1 {
		return result
	}
	var handler func(int, string)
	handler = func(i int, s string) {
		if i == len(digits) {
			result = append(result, s)
			return
		}
		ch := digits[i]
		char := letters[string(ch)]
		for _, v := range char {
			handler(i+1, s+string(v))
		}
	}
	handler(0, "")
	return result
}

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func main() {
	fmt.Println(letterCombinations("23"))
}
