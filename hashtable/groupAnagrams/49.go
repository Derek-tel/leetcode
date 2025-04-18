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

type sevenHelper []rune

func (s sevenHelper) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sevenHelper) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s sevenHelper) Len() int {
	return len(s)
}

func seven(str []string) [][]string {
	var result [][]string
	dic := make(map[string][]string)
	for _, s := range str {
		temp := sevenHelper(s)
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

type eightList []rune

func (e eightList) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
func (e eightList) Less(i, j int) bool {
	return e[i] < e[j]
}
func (e eightList) Len() int {
	return len(e)
}
func eight(str []string) [][]string {
	var result [][]string

	dic := make(map[string][]string)
	for _, s := range str {
		temp := eightList(s)
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

type nineList []rune

func (n nineList) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
func (n nineList) Less(i, j int) bool {
	return n[i] < n[j]
}
func (n nineList) Len() int {
	return len(n)
}

func nine(strs []string) [][]string {
	var result [][]string

	if len(strs) == 0 {
		return result
	}
	dic := make(map[string][]string)
	for _, str := range strs {
		temp := nineList(str)
		sort.Sort(temp)
		tag := dic[string(temp)]
		tag = append(tag, str)
		dic[string(temp)] = tag
	}

	for _, strings := range dic {
		result = append(result, strings)
	}
	return result
}

type tenList []rune

func (t tenList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t tenList) Less(i, j int) bool {
	return t[i] < t[j]
}

func (t tenList) Len() int {
	return len(t)
}

func ten(strs []string) [][]string {
	var result [][]string

	if len(strs) == 0 {
		return result
	}
	dic := make(map[string][]string)
	for _, str := range strs {
		temp := tenList(str)
		sort.Sort(temp)
		tag := dic[string(temp)]
		tag = append(tag, str)
		dic[string(temp)] = tag
	}

	for _, strings := range dic {
		result = append(result, strings)
	}
	return result
}

type elevenList []rune

func (t elevenList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t elevenList) Less(i, j int) bool {
	return t[i] < t[j]
}

func (t elevenList) Len() int {
	return len(t)
}

func eleven(strs []string) [][]string {
	var result [][]string
	if len(strs) == 0 {
		return result
	}
	dic := make(map[string][]string)
	for _, str := range strs {
		temp := elevenList(str)
		sort.Sort(temp)
		tag := dic[string(temp)]
		tag = append(tag, str)
		dic[string(temp)] = tag
	}
	for _, strings := range dic {
		result = append(result, strings)
	}
	return result
}

type twelveList []rune

func (t twelveList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t twelveList) Less(i, j int) bool {
	return t[i] < t[j]
}
func (t twelveList) Len() int {
	return len(t)
}
func twelve(strs []string) [][]string {
	var result [][]string
	if len(strs) == 0 {
		return result
	}
	dic := make(map[string][]string)
	for _, str := range strs {
		temp := twelveList(str)
		sort.Sort(temp)
		tag := dic[string(temp)]
		tag = append(tag, str)
		dic[string(temp)] = tag
	}
	for _, strings := range dic {
		result = append(result, strings)
	}
	return result
}

type thirteenList []rune

func (t thirteenList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
func (t thirteenList) Less(i, j int) bool {
	return t[i] < t[j]
}
func (t thirteenList) Len() int {
	return len(t)
}

func thirteen(strs []string) [][]string {
	var result [][]string
	if len(strs) == 0 {
		return result
	}
	dic := make(map[string][]string)
	for _, str := range strs {
		temp := thirteenList(str)
		sort.Sort(temp)
		tag := dic[string(temp)]
		tag = append(tag, str)
		dic[string(temp)] = tag
	}
	for _, v := range dic {
		result = append(result, v)
	}
	return result
}

type fourteenList []rune

func (t fourteenList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t fourteenList) Less(i, j int) bool {
	return t[i] < t[j]
}

func (t fourteenList) Len() int {
	return len(t)
}

func fourteen(strs []string) [][]string {
	var result [][]string

	if len(strs) == 0 {
		return result
	}
	dic := make(map[string][]string)
	for _, str := range strs {
		temp := fourteenList(str)
		sort.Sort(temp)
		tag := dic[string(temp)]
		tag = append(tag, str)
		dic[string(temp)] = tag
	}
	for _, v := range dic {
		result = append(result, v)
	}
	return result
}

type fifteenList []rune

func (t fifteenList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
func (t fifteenList) Less(i, j int) bool {
	return t[i] < t[j]
}
func (t fifteenList) Len() int {
	return len(t)
}

func fifteen(strs []string) [][]string {
	var result [][]string

	if len(strs) == 0 {
		return result
	}
	dic := make(map[string][]string)
	for _, str := range strs {
		temp := fifteenList(str)
		sort.Sort(temp)
		tag := dic[string(temp)]
		tag = append(tag, str)
		dic[string(temp)] = tag
	}
	for _, v := range dic {
		result = append(result, v)
	}
	return result
}
