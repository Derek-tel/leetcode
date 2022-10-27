package main

import "fmt"

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

func main() {
	fmt.Println(letterCombinations("23"))
}
