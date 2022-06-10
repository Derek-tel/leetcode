package minDistance

func minDistance(word1 string, word2 string) int {
	lengthA := len(word1)
	lengthB := len(word2)

	if lengthA == 0 || lengthB == 0 {
		return lengthA + lengthB
	}
	dp := make([][]int, lengthA+1)
	for index, _ := range dp {
		dp[index] = make([]int, lengthB+1)
	}
	for i := 0; i <= lengthA; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= lengthB; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= lengthA; i++ {
		for j := 1; j <= lengthB; j++ {
			left := dp[i-1][j] + 1
			down := dp[i][j-1] + 1
			leftDown := dp[i-1][j-1] + 1
			if word1[i-1] == word2[j-1] {
				leftDown--
			}
			dp[i][j] = min(left, min(down, leftDown))
		}
	}
	return dp[lengthA][lengthB]
}

func demo(word1, word2 string) int {
	lengthA := len(word1)
	lengthB := len(word2)
	if lengthA == 0 || lengthB == 0 {
		return lengthA + lengthB
	}
	dp := make([][]int, lengthA+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, lengthB+1)
	}
	for i := 0; i <= lengthA; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= lengthB; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= lengthA; i++ {
		for j := 1; j <= lengthB; j++ {
			left := dp[i-1][j] + 1
			down := dp[i][j-1] + 1
			leftDown := dp[i-1][j-1] + 1
			if word1[i-1] == word2[j-1] {
				leftDown--
			}
			dp[i][j] = min(left, min(down, leftDown))
		}
	}
	return dp[lengthA][lengthB]
}

func test(word1 string, word2 string) int {
	lengthA := len(word1)
	lengthB := len(word2)
	if lengthA == 0 || lengthB == 0 {
		return lengthA + lengthB
	}
	dp := make([][]int, lengthA+1)
	for i, _ := range dp {
		dp[i] = make([]int, lengthB+1)
	}

	for i := 0; i <= lengthA; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= lengthB; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= lengthA; i++ {
		for j := 1; j <= lengthB; j++ {
			left := dp[i-1][j] + 1
			down := dp[i][j-1] + 1
			leftDown := dp[i-1][j-1] + 1
			if word1[i-1] == word2[j-1] {
				leftDown--
			}
			dp[i][j] = min(left, min(down, leftDown))
		}
	}
	return dp[lengthA][lengthB]
}

func get(word1, word2 string) int {
	lengthA := len(word1)
	lengthB := len(word2)
	if lengthA == 0 || lengthB == 0 {
		return lengthA + lengthB
	}
	dp := make([][]int, lengthA+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, lengthB+1)
	}

	for i := 0; i <= lengthA; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= lengthB; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= lengthA; i++ {
		for j := 1; j <= lengthB; j++ {
			left := dp[i-1][j] + 1
			down := dp[i][j-1] + 1
			leftDone := dp[i-1][j-1] + 1
			if word1[i-1] == word2[j-1] {
				leftDone--
			}
			dp[i][j] = min(min(left, down), leftDone)
		}
	}
	return dp[lengthA][lengthB]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
