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

func three(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	pair := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}
	var stack []byte
	for i := 0; i < len(s); i++ {
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

func four(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	pair := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}
	var stack []byte
	for i := 0; i < len(s); i++ {
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

func five(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	pair := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}
	var stack []byte
	for i := 0; i < len(s); i++ {
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

func six(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	pair := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}
	var stack []byte
	for i := 0; i < len(s); i++ {
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

func seven(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	pair := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}
	var stack []byte
	for i := 0; i < len(s); i++ {
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

func eight(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	pair := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}
	var stack []byte
	for i := 0; i < len(s); i++ {
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

func nine(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	pair := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}
	var stack []byte
	for i := 0; i < len(s); i++ {
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

func ten(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	pair := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}
	var stack []byte
	for i := 0; i < len(s); i++ {
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

func eleven(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	pair := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}
	var stack []byte
	for i := 0; i < len(s); i++ {
		if tag, ok := pair[s[i]]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != tag {
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

func twelve(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	pair := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}
	var stack []byte
	for i := 0; i < len(s); i++ {
		if tag, ok := pair[s[i]]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != tag {
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
