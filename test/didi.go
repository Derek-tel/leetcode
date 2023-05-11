package main

import "fmt"

func didi(str string) int {
	result := 0
	length := len(str)
	if length == 0 {
		return result
	}

	dp := make([][]bool, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
	}

	for i := length - 1; i >= 0; i-- {
		for j := i; j < length; j++ {
			dp[i][j] = (str[i] == str[j]) && (j-i < 3 || dp[i+1][j-1])
			if dp[i][j] {
				result++
			}
		}
	}
	return result
}

func main() {
	str := "aaaaaa"
	fmt.Println(didi(str))
}
