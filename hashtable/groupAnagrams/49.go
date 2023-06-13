package groupAnagrams

import (
	"sort"
)

type sortRune []rune

func (s sortRune) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s sortRune) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s sortRune) Len() int {
	return len(s)
}

func groupAnagrams(strs []string) [][]string {
	length := len(strs)
	if length == 0 {
		return [][]string{}
	}
	mapDict := make(map[string][]string)

	for _, str := range strs {
		sByte := []rune(str)
		sort.Sort(sortRune(sByte))
		mapStr := mapDict[string(sByte)]
		mapStr = append(mapStr, str)
		mapDict[string(sByte)] = mapStr
	}
	res := [][]string{}

	for _, strings := range mapDict {
		res = append(res, strings)
	}
	return res
}

type so []rune

func (s so) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s so) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s so) Len() int {
	return len(s)
}
func test(strs []string) [][]string {
	mapDict := make(map[string][]string)
	for _, str := range strs {
		sByte := so(str)
		sort.Sort(sByte)
		s := mapDict[string(sByte)]
		s = append(s, str)
		mapDict[string(sByte)] = s
	}

	res := [][]string{}

	for _, strings := range mapDict {
		res = append(res, strings)
	}
	return res
}

type g0 []rune

func (g g0) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

func (g g0) Less(i, j int) bool {
	return g[i] < g[j]
}

func (g g0) Len() int {
	return len(g)
}

func get(str []string) [][]string {
	mapDict := make(map[string][]string)
	for _, s := range str {
		tag := g0(s)
		sort.Sort(tag)
		temp := mapDict[string(tag)]
		temp = append(temp, s)
		mapDict[string(tag)] = temp
	}
	var res [][]string
	for _, strings := range mapDict {
		res = append(res, strings)
	}
	return res
}

type helper []rune

func (h helper) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h helper) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h helper) Len() int {
	return len(h)
}

func four(str []string) [][]string {
	dict := make(map[string][]string)
	for _, s := range str {
		temp := helper(s)
		sort.Sort(temp)
		tag := dict[string(temp)]
		tag = append(tag, s)
		dict[string(temp)] = tag
	}
	var result [][]string
	for _, strings := range dict {
		result = append(result, strings)
	}
	return result
}

func five(str []string) [][]string {
	dict := make(map[string][]string)
	for _, s := range str {
		v := []rune(s)
		sort.Slice(v, func(i, j int) bool {
			return v[i] < v[j]
		})
		tag := dict[string(v)]
		tag = append(tag, s)
		dict[string(v)] = tag
	}
	var result [][]string
	for _, strings := range dict {
		result = append(result, strings)
	}
	return result
}

type sixHelper []rune

func (s sixHelper) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s sixHelper) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sixHelper) Len() int {
	return len(s)
}

func six(str []string) [][]string {
	var result [][]string
	dic := make(map[string][]string)
	for _, s := range str {
		temp := helper(s)
		sort.Sort(temp)
		tag := dic[string(temp)]
		tag = append(tag, s)
		dic[string(temp)] = tag
	}

	for _, strings := range dic {
		result = append(result, strings)
	}
	return result
}
