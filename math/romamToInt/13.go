package romamToInt

func one(s string) int {
	result := 0
	n := len(s)
	for i := range s {
		value := symbolValue[s[i]]
		if i < n-1 && value < symbolValue[s[i+1]] {
			result -= value
		} else {
			result += value
		}
	}
	return result
}

func two(s string) int {
	result := 0
	n := len(s)
	for i := range s {
		value := symbolValue[s[i]]
		if i < n-1 && value < symbolValue[s[i+1]] {
			result -= value
		} else {
			result += value
		}
	}
	return result
}

func three(s string) int {
	result := 0
	n := len(s)
	for i := range s {
		value := symbolValue[s[i]]
		if i < n-1 && value < symbolValue[s[i+1]] {
			result -= value
		} else {
			result += value
		}
	}
	return result
}

func four(s string) int {
	result := 0
	n := len(s)
	for i := range s {
		value := symbolValue[s[i]]
		if i < n-1 && value < symbolValue[s[i+1]] {
			result -= value
		} else {
			result += value
		}
	}
	return result
}

var symbolValue = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}
