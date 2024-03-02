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

func five(strs []string) string {
	var lcp func(string, string) string
	lcp = func(left string, right string) string {
		minLength := min(len(left), len(right))
		for i := 0; i < minLength; i++ {
			if left[i] != right[i] {
				return left[:i]
			}
		}
		return left[:minLength]
	}
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	} else if len(strs) == 2 {
		return lcp(strs[0], strs[1])
	} else {
		mid := len(strs) / 2
		return lcp(five(strs[:mid]), five(strs[mid:]))
	}
}

func six(strs []string) string {
	var lcp func(string, string) string
	lcp = func(left string, right string) string {
		minLength := min(len(left), len(right))
		for i := 0; i < minLength; i++ {
			if left[i] != right[i] {
				return left[:i]
			}
		}
		return left[:minLength]
	}

	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	} else if len(strs) == 2 {
		return lcp(strs[0], strs[1])
	} else {
		mid := len(strs) / 2
		return lcp(six(strs[:mid]), six(strs[mid:]))
	}
}

func seven(strs []string) string {
	var lcp func(string, string) string
	lcp = func(left string, right string) string {
		minLength := min(len(left), len(right))
		for i := 0; i < minLength; i++ {
			if left[i] != right[i] {
				return left[:i]
			}
		}
		return left[:minLength]
	}

	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	} else if len(strs) == 2 {
		return lcp(strs[0], strs[1])
	} else {
		mid := len(strs) / 2
		return lcp(seven(strs[:mid]), seven(strs[mid:]))
	}
}

func eight(strs []string) string {
	var lcp func(string, string) string
	lcp = func(left string, right string) string {
		minLength := min(len(left), len(right))
		for i := 0; i < minLength; i++ {
			if left[i] != right[i] {
				return left[:i]
			}
		}
		return left[:minLength]
	}
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	} else if len(strs) == 2 {
		return lcp(strs[0], strs[1])
	} else {
		mid := len(strs) / 2
		return lcp(eight(strs[:mid]), eight(strs[mid:]))
	}
}

func nine(strs []string) string {
	var lcp func(string, string) string
	lcp = func(left string, right string) string {
		minLength := min(len(left), len(right))
		for i := 0; i < minLength; i++ {
			if left[i] != right[i] {
				return left[:i]
			}
		}
		return left[:minLength]
	}

	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	} else if len(strs) == 2 {
		return lcp(strs[0], strs[1])
	} else {
		mid := len(strs) / 2
		return lcp(nine(strs[:mid]), nine(strs[mid:]))
	}
}

func ten(strs []string) string {
	var lcp func(string, string) string
	lcp = func(left string, right string) string {
		minLength := min(len(left), len(right))
		for i := 0; i < minLength; i++ {
			if left[i] != right[i] {
				return left[:i]
			}
		}
		return left[:minLength]
	}
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	} else if len(strs) == 2 {
		return lcp(strs[0], strs[1])
	} else {
		mid := len(strs) / 2
		return lcp(ten(strs[:mid]), ten(strs[mid:]))
	}
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
