package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	stack := []string{}
	for _, str := range strings.Split(path, "/") {
		if str == ".." {
			if len(stack) > 0 {
				stack = stack[0 : len(stack)-1]
			}
		} else if str != "." && str != "" {
			stack = append(stack, str)
		}
	}
	return "/" + strings.Join(stack, "/")
}

func test(path string) string {
	stack := []string{}

	for _, v := range strings.Split(path, "/") {
		if v == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if v != "." && v != "" {
			stack = append(stack, v)
		}
	}
	return "/" + strings.Join(stack, "/")
}

func get(path string) string {
	var stack []string

	for _, v := range strings.Split(path, "/") {
		if v == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if v != "." && v != "" {
			stack = append(stack, v)
		}
	}
	return "/" + strings.Join(stack, "/")
}

func four(path string) string {
	var stack []string

	for _, v := range strings.Split(path, "/") {
		if v == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if v != "." && v != "" {
			stack = append(stack, v)
		}
	}
	return "/" + strings.Join(stack, "/")
}

func five(path string) string {
	var stack []string

	for _, v := range strings.Split(path, "/") {
		if v == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if v != "." && v != "" {
			stack = append(stack, v)
		}
	}
	return "/" + strings.Join(stack, "/")
}

func six(path string) string {
	var stack []string
	for _, v := range strings.Split(path, "/") {
		if v == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if v != "." && v != "" {
			stack = append(stack, v)
		}
	}
	return "/" + strings.Join(stack, "/")
}

func seven(path string) string {
	var stack []string
	for _, v := range strings.Split(path, "/") {
		if v == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if v != "." && v != "" {
			stack = append(stack, v)
		}
	}
	return "/" + strings.Join(stack, "/")
}

func eight(path string) string {
	var stack []string
	for _, v := range strings.Split(path, "/") {
		if v == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if v != "." && v != "" {
			stack = append(stack, v)
		}
	}
	return "/" + strings.Join(stack, "/")
}

func main() {
	test := "//root/"
	fmt.Println(strings.Split(test, "/"))

	fmt.Println(simplifyPath(test))
}
