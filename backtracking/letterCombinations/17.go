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

func main() {
	fmt.Println(letterCombinations("23"))
}
