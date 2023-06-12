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
