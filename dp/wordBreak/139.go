package wordBreak

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dictMap := make(map[string]bool)
	for _, word := range wordDict {
		dictMap[word] = true
	}
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && dictMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func demo(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dictMap := make(map[string]bool)
	for _, v := range wordDict {
		dictMap[v] = true
	}
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && dictMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func test(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dictMap := make(map[string]bool)
	for _, v := range wordDict {
		dictMap[v] = true
	}

	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && dictMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func get(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dictMap := make(map[string]bool)
	for _, word := range wordDict {
		dictMap[word] = true
	}

	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && dictMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func five(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dictMap := make(map[string]bool)
	for _, word := range wordDict {
		dictMap[word] = true
	}

	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && dictMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func six(s string, wordDict []string) bool {
	dic := make(map[string]bool)
	for _, v := range wordDict {
		dic[v] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && dic[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func seven(s string, wordDict []string) bool {
	dic := make(map[string]bool)
	for _, v := range wordDict {
		dic[v] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && dic[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func eight(s string, wordDict []string) bool {
	dic := make(map[string]bool)
	for _, s2 := range wordDict {
		dic[s2] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && dic[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func nine(s string, wordDict []string) bool {
	dic := make(map[string]bool)
	for _, s2 := range wordDict {
		dic[s2] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && dic[s[j:i]] {
				dp[i] = true
			}
		}
	}
	return dp[len(s)]
}

func ten(s string, wordDict []string) bool {
	dic := make(map[string]bool)
	for _, word := range wordDict {
		dic[word] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && dic[s[j:i]] {
				dp[i] = true
			}
		}
	}
	return dp[len(s)]
}

func eleven(s string, wordDict []string) bool {
	if len(s) == 0 {
		return false
	}
	dic := make(map[string]bool)
	for _, w := range wordDict {
		dic[w] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= i; j++ {
			if dp[j-1] && dic[s[j-1:i-1+1]] {
				dp[i] = true
			}
		}
	}
	return dp[len(s)]
}

func twelve(s string, wordDict []string) bool {
	if len(s) == 0 {
		return false
	}
	dic := make(map[string]bool)
	for _, v := range wordDict {
		dic[v] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= i; j++ {
			if dp[j-1] && dic[s[j-1:i-1+1]] {
				dp[i] = true
			}
		}
	}
	return dp[len(s)]
}

func thirteen(s string, wordDict []string) bool {
	if len(s) == 0 {
		return false
	}
	dic := make(map[string]bool)
	for _, v := range wordDict {
		dic[v] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= i; j++ {
			if dp[j-1] && dic[s[j-1:i-1+1]] {
				dp[i] = true
			}
		}
	}
	return dp[len(s)]
}

func fourteen(s string, wordDict []string) bool {
	if len(s) == 0 {
		return false
	}
	dic := make(map[string]bool)
	for _, v := range wordDict {
		dic[v] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && dic[s[j+1-1:i-1+1]] {
				dp[i] = true
			}
		}
	}
	return dp[len(s)]
}

func fifteen(s string, wordDict []string) bool {
	if len(s) == 0 {
		return false
	}
	dic := make(map[string]bool)
	for _, v := range wordDict {
		dic[v] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && dic[s[j+1-1:i-1+1]] {
				dp[i] = true
			}
		}
	}
	return dp[len(s)]
}

func sixteen(s string, wordDict []string) bool {
	if len(s) == 0 {
		return false
	}
	dic := make(map[string]bool)
	for _, w := range wordDict {
		dic[w] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && dic[s[j+1-1:i-1+1]] {
				dp[i] = true
			}
		}
	}
	return dp[len(s)]
}
