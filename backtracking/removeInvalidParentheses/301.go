package main

import "fmt"

func isValid(s string) bool {
	cut := 0
	for _, str := range s {
		if str == '(' {
			cut++
		} else if str == ')' {
			cut--
			if cut < 0 {
				return false
			}
		}
	}
	if cut != 0 {
		return false
	}
	return true
}

func removeInvalidParentheses(s string) []string {
	lRemove, rRemove := 0, 0
	for _, str := range s {
		if str == '(' {
			lRemove++
		} else if str == ')' {
			if lRemove == 0 {
				rRemove++
			} else {
				lRemove--
			}
		}
	}
	ans := []string{}
	helper(&ans, s, 0, lRemove, rRemove)
	return ans
}

func helper(ans *[]string, str string, start int, lRemove int, rRemove int) {
	if lRemove == 0 && rRemove == 0 {
		if isValid(str) {
			*ans = append(*ans, str)
		}
		return
	}

	for i := start; i < len(str); i++ {
		if i != start && str[i] == str[i-1] {
			continue
		}
		if lRemove+rRemove > len(str)-i {
			return
		}
		if lRemove > 0 && str[i] == '(' {
			helper(ans, str[:i]+str[i+1:], i, lRemove-1, rRemove)
		}
		if rRemove > 0 && str[i] == ')' {
			helper(ans, str[:i]+str[i+1:], i, lRemove, rRemove-1)
		}
	}
}

func isV(s string) bool {
	count := 0
	for _, ch := range s {
		if ch == '(' {
			count++
		} else {
			count--
			if count < 0 {
				return false
			}
		}
	}
	if count != 0 {
		return false
	}
	return true
}

func remove(s string) []string {
	lmove := 0
	rmove := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			lmove++
		} else {
			if lmove > 0 {
				lmove--
			} else {
				rmove++
			}
		}
	}
	ans := []string{}
	handler(&ans, s, 0, lmove, rmove)
	return ans
}

func handler(ans *[]string, s string, start int, lmove int, rmove int) {
	if lmove == 0 && rmove == 0 {
		if isV(s) {
			*ans = append(*ans, s)
		}
	}

	for i := start; i < len(s); i++ {
		if i != start && s[i] == s[i-1] {
			continue
		}
		if lmove+rmove > len(s)-i {
			return
		}
		if s[i] == '(' && lmove > 0 {
			handler(ans, s[:i]+s[i+1:], i, lmove-1, rmove)
		}
		if s[i] == ')' && rmove > 0 {
			handler(ans, s[:i]+s[i+1:], i, lmove, rmove-1)
		}
	}

}

func three(s string) []string {
	lmove, rmove := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			lmove++
		} else if s[i] == ')' {
			if lmove > 0 {
				lmove--
			} else {
				rmove++
			}
		}
	}
	var result []string
	var backtrack func(string, int, int, int)
	backtrack = func(s string, start int, lmove int, rmove int) {
		if lmove == 0 && rmove == 0 {
			if judge(s) {
				result = append(result, s)
			}
		}
		for i := start; i < len(s); i++ {
			if i != start && s[i] == s[i-1] {
				continue
			}
			if lmove+rmove > len(s)-i {
				return
			}
			if lmove > 0 && s[i] == '(' {
				backtrack(s[:i]+s[i+1:], i, lmove-1, rmove)
			}
			if rmove > 0 && s[i] == ')' {
				backtrack(s[:i]+s[i+1:], i, lmove, rmove-1)
			}
		}
	}
	backtrack(s, 0, lmove, rmove)
	return result
}

func judge(s string) bool {
	cut := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			cut++
		} else if s[i] == ')' {
			cut--
			if cut < 0 {
				return false
			}
		}
	}
	return cut == 0
}

func four(s string) []string {
	var valid func(string) bool
	valid = func(str string) bool {
		cut := 0
		for i := 0; i < len(str); i++ {
			if str[i] == '(' {
				cut++
			} else if str[i] == ')' {
				cut--
				if cut < 0 {
					return false
				}
			}
		}
		return cut == 0
	}
	lmove, rmove := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			lmove++
		} else if s[i] == ')' {
			if lmove > 0 {
				lmove--
			} else {
				rmove++
			}
		}
	}

	var result []string
	var backtrack func(string, int, int, int)
	backtrack = func(str string, index int, lmove int, rmove int) {
		if lmove == 0 && rmove == 0 {
			if valid(str) {
				result = append(result, str)
			}
			return
		}
		for i := index; i < len(str); i++ {
			if i != index && str[i] == str[i-1] {
				continue
			}
			if lmove+rmove > len(str)-index {
				return
			}
			if str[i] == '(' && lmove > 0 {
				backtrack(str[:i]+str[i+1:], i, lmove-1, rmove)
			}
			if str[i] == ')' && rmove > 0 {
				backtrack(str[:i]+str[i+1:], i, lmove, rmove-1)
			}
		}
	}

	backtrack(s, 0, lmove, rmove)
	return result
}
func main() {
	test := "()())()"
	fmt.Println(removeInvalidParentheses(test))
}
