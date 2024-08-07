package main

import (
	"fmt"
)

func findAnagrams(s string, p string) []int {
	var freq [256]int
	lengthS := len(s)
	lengthP := len(p)
	for _, str := range p {
		freq[str-'a']++
	}

	left, right, count := 0, 0, lengthP
	res := []int{}
	for right < lengthS {
		if freq[s[right]-'a'] > 0 {
			count--
		}
		freq[s[right]-'a']--
		if count == 0 {
			res = append(res, left)
		}
		if right-left+1 == lengthP {
			if freq[s[left]-'a'] >= 0 {
				count++
			}
			freq[s[left]-'a']++
			left++
		}
		right++
	}
	return res
}

func findAnagrams1(s string, p string) []int {
	counter := make(map[byte]int)
	lengthS := len(s)
	lengthP := len(p)
	for i := 0; i < lengthP; i++ {
		counter[p[i]-'a']++
	}
	copy := copyMap(counter)
	res := []int{}
	for i, start := 0, 0; i < lengthS && start < lengthS; i++ {
		if copy[s[i]-'a'] >= 1 {
			copy[s[i]-'a']--
			if check(copy) {
				res = append(res, start)
			}
		} else {
			start++
			i = start - 1
			copy = copyMap(counter)
		}
	}
	return res
}

func check(s map[byte]int) bool {
	for _, i := range s {
		if i != 0 {
			return false
		}
	}
	return true
}

func copyMap(s map[byte]int) map[byte]int {
	c := map[byte]int{}
	for k, v := range s {
		c[k] = v
	}
	return c
}

func test(s string, p string) []int {
	lengthP := len(p)
	lengthS := len(s)
	var freq [256]int

	for _, i := range p {
		freq[i-'a']++
	}
	left, right := 0, 0
	count := lengthP
	var res []int
	for right < lengthS {
		if freq[s[right]-'a'] > 0 {
			count--
		}
		freq[s[right]-'a']--
		if count == 0 {
			res = append(res, left)
		}
		if right-left+1 == lengthP {
			if freq[s[left]-'a'] >= 0 {
				count++
			}
			freq[s[left]-'a']++
			left++
		}
		right++
	}
	return res
}

func get(s, p string) []int {
	lengthS := len(s)
	lengthP := len(p)

	var freq [256]int
	for i := 0; i < lengthP; i++ {
		freq[p[i]-'a']++
	}
	left, right := 0, 0
	count := lengthP
	var res []int
	for right < lengthS {
		if freq[s[right]-'a'] > 0 {
			count--
		}
		freq[s[right]-'a']--
		if count == 0 {
			res = append(res, left)
		}
		if right-left+1 == lengthP {
			if freq[s[left]-'a'] >= 0 {
				count++
			}
			freq[s[left]-'a']++
			left++
		}
		right++
	}
	return res
}

func five(s, p string) []int {
	lengthS := len(s)
	lengthP := len(p)

	var freq [256]int
	for i := 0; i < lengthP; i++ {
		freq[p[i]-'a']++
	}
	left, right := 0, 0
	count := lengthP
	var result []int
	for right < lengthS {
		if freq[s[right]-'a'] > 0 {
			count--
		}
		freq[s[right]-'a']--
		if count == 0 {
			result = append(result, left)
		}
		if right-left+1 == lengthP {
			if freq[s[left]-'a'] >= 0 {
				count++
			}
			freq[s[left]-'a']++
			left++
		}
		right++
	}
	return result
}

func six(s, p string) []int {
	lengthS := len(s)
	lengthP := len(p)

	var freq [256]int
	for i := 0; i < lengthP; i++ {
		freq[p[i]-'a']++
	}
	left, right := 0, 0
	count := lengthP
	var result []int
	for right < lengthS {
		if freq[s[right]-'a'] > 0 {
			count--
		}
		freq[s[right]-'a']--
		if count == 0 {
			result = append(result, left)
		}
		if right-left+1 == lengthP {
			if freq[s[left]-'a'] >= 0 {
				count++
			}
			freq[s[left]-'a']++
			left++
		}
		right++
	}
	return result
}

func seven(s, p string) []int {
	lengthS := len(s)
	lengthP := len(p)

	var freq [256]int
	for i := 0; i < lengthP; i++ {
		freq[p[i]-'a']++
	}
	left, right := 0, 0
	count := lengthP
	var result []int
	for right < lengthS {
		if freq[s[right]-'a'] > 0 {
			count--
		}
		freq[s[right]-'a']--
		if count == 0 {
			result = append(result, left)
		}
		if right-left+1 == lengthP {
			if freq[s[left]-'a'] >= 0 {
				count++
			}
			freq[s[left]-'a']++
			left++
		}
		right++
	}
	return result
}

func eight(s, p string) []int {
	lengthS := len(s)
	lengthP := len(p)

	freq := make(map[uint8]int)
	for i := 0; i < lengthP; i++ {
		freq[p[i]-'a']++
	}
	left, right := 0, 0
	count := 0

	var result []int
	tempFreq := make(map[uint8]int)
	for ; right < lengthS; right++ {
		if num, ok := freq[s[right]-'a']; ok {
			for tempFreq[s[right]-'a'] >= num {
				tempFreq[s[left]-'a']--
				count--
				left++
			}
			tempFreq[s[right]-'a']++
			count++
		} else {
			left = right + 1
			count = 0
			tempFreq = make(map[uint8]int)
		}
		if count == lengthP {
			result = append(result, left)
		}
	}
	return result
}

func nine(s, p string) []int {
	lengthS, lengthP := len(s), len(p)
	freq := make(map[uint8]int)
	for i := 0; i < lengthP; i++ {
		freq[p[i]-'a']++
	}
	count := 0
	var result []int
	temp := make(map[uint8]int)
	for left, right := 0, 0; right < lengthS; right++ {
		if num, ok := freq[s[right]-'a']; ok {
			for temp[s[right]-'a'] >= num {
				temp[s[left]-'a']--
				count--
				left++
			}
			temp[s[right]-'a']++
			count++
		} else {
			left = right + 1
			count = 0
			temp = make(map[uint8]int)
		}
		if count == lengthP {
			result = append(result, left)
		}
	}
	return result
}

func ten(s, p string) []int {
	lengthS, lengthP := len(s), len(p)
	freq := make(map[int]int)
	for i := 0; i < lengthP; i++ {
		freq[int(p[i]-'a')]++
	}
	count := 0
	var result []int
	temp := make(map[int]int)
	for left, right := 0, 0; right < lengthS; right++ {
		if num, ok := freq[int(s[right]-'a')]; ok {
			for temp[int(s[right]-'a')] >= num {
				temp[int(s[left]-'a')]--
				count--
				left++
			}
			temp[int(s[right]-'a')]++
			count++
		} else {
			left = right + 1
			count = 0
			temp = make(map[int]int)
		}
		if count == lengthP {
			result = append(result, left)
		}
	}
	return result
}

func eleven(s, p string) []int {
	lengthS, lengthP := len(s), len(p)
	freq := make(map[int]int)
	for i := 0; i < len(p); i++ {
		freq[int(p[i]-'a')]++
	}
	count := 0
	var result []int
	temp := make(map[int]int)
	for left, right := 0, 0; right < lengthS; right++ {
		if num, ok := freq[int(s[right]-'a')]; ok {
			for temp[int(s[right]-'a')] >= num {
				temp[int(s[left]-'a')]--
				count--
				left++
			}
			temp[int(s[right]-'a')]++
			count++
		} else {
			left = right + 1
			count = 0
			temp = make(map[int]int)
		}
		if count == lengthP {
			result = append(result, left)
		}
	}
	return result
}

func twelve(s, p string) []int {
	lengthS, lengthP := len(s), len(p)
	freq := make(map[int]int)
	for i := 0; i < len(p); i++ {
		freq[int(p[i]-'a')]++
	}
	count := 0
	var result []int
	temp := make(map[int]int)
	for left, right := 0, 0; right < lengthS; right++ {
		if num, ok := freq[int(s[right]-'a')]; ok {
			for temp[int(s[right]-'a')] >= num {
				temp[int(s[left]-'a')]--
				count--
				left++
			}
			temp[int(s[right]-'a')]++
			count++
		} else {
			left = right + 1
			count = 0
			temp = make(map[int]int)
		}
		if count == lengthP {
			result = append(result, left)
		}
	}
	return result
}

func thirteen(s, p string) []int {
	lengthS, lengthP := len(s), len(p)
	freq := make(map[int]int)
	for i := 0; i < len(p); i++ {
		freq[int(p[i]-'a')]++
	}
	count := 0
	var result []int
	temp := make(map[int]int)
	for left, right := 0, 0; right < lengthS; right++ {
		if num, ok := freq[int(s[right]-'a')]; ok {
			for temp[int(s[right]-'a')] >= num {
				temp[int(s[left]-'a')]--
				count--
				left++
			}
			temp[int(s[right]-'a')]++
			count++
		} else {
			left = right + 1
			count = 0
			temp = make(map[int]int)
		}
		if count == lengthP {
			result = append(result, left)
		}
	}
	return result
}

func fourteen(s, p string) []int {
	lengthS, lengthP := len(s), len(p)
	freq := make(map[int]int)
	for i := 0; i < len(p); i++ {
		freq[int(p[i]-'a')]++
	}
	count := 0
	var result []int
	temp := make(map[int]int)
	for left, right := 0, 0; right < lengthS; right++ {
		if num, ok := freq[int(s[right]-'a')]; ok {
			for temp[int(s[right]-'a')] >= num {
				temp[int(s[left]-'a')]--
				left++
				count--
			}
			temp[int(s[right]-'a')]++
			count++
		} else {
			left = right + 1
			count = 0
			temp = make(map[int]int)
		}
		if count == lengthP {
			result = append(result, left)
		}
	}
	return result
}

func fifteen(s, p string) []int {
	lengthS, lengthP := len(s), len(p)
	freq := make(map[int]int)
	for i := 0; i < len(p); i++ {
		freq[int(p[i]-'a')]++
	}
	var result []int
	count := 0
	temp := make(map[int]int)
	for left, right := 0, 0; right < lengthS; right++ {
		if num, ok := freq[int(s[right]-'a')]; ok {
			for temp[int(s[right]-'a')] >= num {
				temp[int(s[left]-'a')]--
				left++
				count--
			}
			temp[int(s[right]-'a')]++
			count++
		} else {
			left = right + 1
			count = 0
			temp = make(map[int]int)
		}
		if count == lengthP {
			result = append(result, left)
		}
	}
	return result
}

func sixteen(s, p string) []int {
	lengthS, lengthP := len(s), len(p)
	freq := make(map[int]int)
	for i := 0; i < len(p); i++ {
		freq[int(p[i]-'a')]++
	}
	var result []int
	count := 0
	temp := make(map[int]int)
	for left, right := 0, 0; right < lengthS; right++ {
		if num, ok := freq[int(s[right]-'a')]; ok {
			for temp[int(s[right]-'a')] >= num {
				temp[int(s[left]-'a')]--
				left++
				count--
			}
			temp[int(s[right]-'a')]++
			count++
		} else {
			left = right + 1
			count = 0
			temp = make(map[int]int)
		}
		if count == lengthP {
			result = append(result, left)
		}
	}
	return result
}

func main() {
	test := "abab"
	fmt.Println(findAnagrams1(test, "ab"))
}
