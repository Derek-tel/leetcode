package main

import (
	"fmt"
	"strconv"
)

func decodeString(s string) string {
	if len(s) == 0 {
		return ""
	}
	stack := []string{}
	res := ""

	for _, v := range s {
		if len(stack) == 0 || (len(stack) > 0 && v != ']') {
			stack = append(stack, string(v))
		} else {
			tmp := ""
			for stack[len(stack)-1] != "[" {
				tmp = stack[len(stack)-1] + tmp
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			repeat := ""
			index := 0
			for index = len(stack) - 1; index >= 0; index-- {
				if stack[index] >= "0" && stack[index] <= "9" {
					repeat = stack[index] + repeat
				} else {
					break
				}
			}
			nums, _ := strconv.Atoi(repeat)
			copyTmp := tmp
			for i := 0; i < nums-1; i++ {
				tmp = tmp + copyTmp
			}
			for i := 0; i < len(repeat)-1; i++ {
				stack = stack[:len(stack)-1]
			}
			stack[index+1] = tmp
		}
	}
	for len(stack) > 0 {
		res = stack[len(stack)-1] + res
		stack = stack[:len(stack)-1]
	}
	return res
}

func test(s string) string {
	stack := []string{}

	for _, ch := range s {
		if len(stack) == 0 || (len(stack) != 0 && ch != ']') {
			stack = append(stack, string(ch))
		} else {
			temp := ""
			for stack[len(stack)-1] != "[" {
				temp = stack[len(stack)-1] + temp
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			number := ""
			index := 0
			for index = len(stack) - 1; index >= 0; index-- {
				if stack[index] >= "0" && stack[index] <= "9" {
					number = stack[index] + number
				} else {
					break
				}
			}
			num, _ := strconv.Atoi(number)
			add := temp
			for i := 0; i < num-1; i++ {
				add = add + temp
			}
			stack = stack[:index+1]
			stack = append(stack, add)
		}
	}
	ans := ""
	for _, v := range stack {
		ans = ans + v
	}
	return ans
}

func get(s string) string {
	var stack []string

	for _, ch := range s {
		if len(stack) == 0 || ch != ']' {
			stack = append(stack, string(ch))
		} else {
			temp := ""
			for stack[len(stack)-1] != "[" {
				temp = stack[len(stack)-1] + temp
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			num := ""
			index := 0
			for index = len(stack) - 1; index >= 0; index-- {
				if stack[index] >= "0" && stack[index] <= "9" {
					num = stack[index] + num
				} else {
					break
				}
			}
			number, _ := strconv.Atoi(num)
			add := temp
			for i := 0; i < number-1; i++ {
				add = add + temp
			}
			stack = stack[:index+1]
			stack = append(stack, add)
		}
	}
	ans := ""
	for _, v := range stack {
		ans = ans + v
	}
	return ans
}

func main() {
	s := "3[a]a2[bc]"
	fmt.Println(decodeString(s))
}
