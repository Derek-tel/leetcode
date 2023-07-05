package longestCommonPrefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var lcp func(int, int) string
	lcp = func(start int, end int) string {
		if start == end {
			return strs[start]
		}
		mid := (start + end) / 2
		left, right := lcp(start, mid), lcp(mid+1, end)
		minLength := min(len(left), len(right))
		for i := 0; i < minLength; i++ {
			if left[i] != right[i] {
				return left[:i]
			}
		}
		return left[:minLength]
	}
	return lcp(0, len(strs)-1)
}

func two(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var lcp func(int, int) string
	lcp = func(start int, end int) string {
		if start == end {
			return strs[start]
		}
		mid := (start + end) / 2
		left, right := lcp(start, mid), lcp(mid+1, end)
		minLength := min(len(left), len(right))
		for i := 0; i < minLength; i++ {
			if left[i] != right[i] {
				return left[:i]
			}
		}
		return left[:minLength]
	}
	return lcp(0, len(strs)-1)
}

func three(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var lcp func(int, int) string
	lcp = func(start int, end int) string {
		if start == end {
			return strs[start]
		}
		mid := (start + end) / 2
		left, right := lcp(start, mid), lcp(mid+1, end)
		minLength := min(len(left), len(right))
		for i := 0; i < minLength; i++ {
			if left[i] != right[i] {
				return left[:i]
			}
		}
		return left[:minLength]
	}
	return lcp(0, len(strs)-1)
}

func four(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var lcp func(int, int) string
	lcp = func(start int, end int) string {
		if start == end {
			return strs[start]
		}
		mid := (start + end) / 2
		left, right := lcp(start, mid), lcp(mid+1, end)
		minLength := min(len(left), len(right))
		for i := 0; i < minLength; i++ {
			if left[i] != right[i] {
				return left[:i]
			}
		}
		return left[:minLength]
	}
	return lcp(0, len(strs)-1)
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
