package isPalindrome

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	flag := 0
	for x > flag {
		flag = flag*10 + x%10
		x = x / 10
	}
	return x == flag || x == flag/10
}
