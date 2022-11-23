package addBinary

import "strconv"

func addBinary(a string, b string) string {
	resp := ""
	carry := 0
	lengthA, lengthB := len(a), len(b)

	n := max(lengthA, lengthB)

	for i := 0; i < n; i++ {
		if i < lengthA {
			carry += int(a[lengthA-1-i] - '0')
		}
		if i < lengthB {
			carry += int(b[lengthB-1-i] - '0')
		}
		resp = strconv.Itoa(carry%2) + resp
		carry = carry / 2
	}
	if carry > 0 {
		resp = "1" + resp
	}
	return resp
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
