package main

import (
	"fmt"
)

func isValid(s string) bool {
	length := len(s)
	if length%2 == 1 {
		return false
	}
	pair := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := []byte{}
	for i := 0; i < length; i++ {
		if _, ok := pair[s[i]]; ok {
			if len(stack) == 0 || pair[s[i]] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

func get(s string) bool {
	length := len(s)
	if length%2 == 1 {
		return false
	}
	pair := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}
	var stack []byte
	for i := 0; i < length; i++ {
		if tag, ok := pair[s[i]]; ok {
			if len(stack) == 0 || tag != stack[len(stack)-1] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

func main() {
	test := "()"
	fmt.Println(isValid(test))
}
