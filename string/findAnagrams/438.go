package main

import "fmt"

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

func main() {
	test := "abab"
	fmt.Println(findAnagrams1(test, "ab"))
}
