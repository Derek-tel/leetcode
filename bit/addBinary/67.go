package addBinary

import (
	"strconv"
)

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

func two(a string, b string) string {
	resp := ""
	carry := 0
	lengthA, lengthB := len(a), len(b)
	n := max(lengthA, lengthB)
	for i := 0; i < n; i++ {
		if i < lengthA {
			carry += int(a[lengthA-i-1] - '0')
		}
		if i < lengthB {
			carry += int(b[lengthB-i-1] - '0')
		}
		resp = strconv.Itoa(carry%2) + resp
		carry = carry / 2
	}
	if carry > 0 {
		resp = "1" + resp
	}
	return resp
}

func three(a string, b string) string {
	resp := ""
	carry := 0
	lengthA, lengthB := len(a), len(b)
	n := max(lengthA, lengthB)
	for i := 0; i < n; i++ {
		if i < lengthA {
			carry = carry + int(a[lengthA-i-1]-'0')
		}
		if i < lengthB {
			carry = carry + int(b[lengthB-i-1]-'0')
		}
		resp = strconv.Itoa(carry%2) + resp
		carry = carry / 2
	}
	if carry > 0 {
		resp = "1" + resp
	}
	return resp
}

func four(a string, b string) string {
	resp := ""
	carry := 0
	lengthA, lengthB := len(a), len(b)
	n := max(lengthA, lengthB)
	for i := 0; i < n; i++ {
		if i < lengthA {
			carry = carry + int(a[lengthA-i-1]-'0')
		}
		if i < lengthB {
			carry = carry + int(b[lengthB-i-1]-'0')
		}
		resp = strconv.Itoa(carry%2) + resp
		carry = carry / 2
	}
	if carry > 0 {
		resp = "1" + resp
	}
	return resp
}

func five(a string, b string) string {
	var resp string
	carry := 0
	lengthA, lengthB := len(a), len(b)
	n := max(lengthA, lengthB)
	for i := 0; i < n; i++ {
		if i < lengthA {
			carry += int(a[lengthA-i-1] - '0')
		}
		if i < lengthB {
			carry += int(b[lengthB-i-1] - '0')
		}
		resp = strconv.Itoa(carry%2) + resp
		carry /= 2
	}
	if carry > 0 {
		resp = "1" + resp
	}
	return resp
}

func six(a, b string) string {
	var resp string
	carry := 0
	lengthA, lengthB := len(a), len(b)
	n := max(lengthA, lengthB)
	for i := 0; i < n; i++ {
		if i < lengthA {
			carry += int(a[lengthA-i-1] - '0')
		}
		if i < lengthB {
			carry += int(a[lengthB-i-1] - '0')
		}
		resp = strconv.Itoa(carry%2) + resp
		carry = carry / 2
	}
	if carry > 0 {
		resp = "1" + resp
	}
	return resp
}

func seven(a, b string) string {
	var resp string
	carry := 0
	lengthA, lengthB := len(a), len(b)
	n := max(lengthA, lengthB)
	for i := 0; i < n; i++ {
		if i < lengthA {
			carry += int(a[lengthA-i-1] - '0')
		}
		if i < lengthB {
			carry += int(b[lengthB-i-1] - '0')
		}
		resp = strconv.Itoa(carry%2) + resp
		carry = carry / 2
	}
	if carry > 0 {
		resp = "1" + resp
	}
	return resp
}

func eight(a, b string) string {
	var resp string
	carry := 0
	var lengthA, lengthB = len(a), len(b)
	n := max(lengthA, lengthB)
	for i := 0; i < n; i++ {
		if i < lengthA {
			carry += int(a[lengthA-i-1] - '0')
		}
		if i < lengthB {
			carry += int(b[lengthB-i-1] - '0')
		}
		resp = strconv.Itoa(carry%2) + resp
		carry = carry / 2
	}
	if carry > 0 {
		resp = "1" + resp
	}
	return resp
}

func nine(a, b string) string {
	var resp string
	carry := 0
	var lengthA, lengthB = len(a), len(b)
	n := max(lengthA, lengthB)
	for i := 0; i < n; i++ {
		if i < lengthA {
			carry += int(a[lengthA-i-1] - '0')
		}
		if i < lengthB {
			carry += int(b[lengthB-i-1] - '0')
		}
		resp = strconv.Itoa(carry%2) + resp
		carry = carry / 2
	}
	if carry > 0 {
		resp = "1" + resp
	}
	return resp
}

func ten(a, b string) string {
	var resp string
	carry := 0
	var lengthA, lengthB = len(a), len(b)
	n := max(lengthA, lengthB)
	for i := 0; i < n; i++ {
		if i < lengthA {
			carry += int(a[lengthA-i-1] - '0')
		}
		if i < lengthB {
			carry += int(b[lengthB-i-1] - '0')
		}
		resp = strconv.Itoa(carry%2) + resp
		carry = carry / 2
	}
	if carry > 0 {
		resp = "1" + resp
	}
	return resp
}

func eleven(a, b string) string {
	var resp string
	carry := 0
	var lengthA, lengthB = len(a), len(b)
	n := max(lengthA, lengthB)
	for i := 0; i < n; i++ {
		if i < lengthA {
			carry += int(a[lengthA-i-1] - '0')
		}
		if i < lengthB {
			carry += int(b[lengthB-i-1] - '0')
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
