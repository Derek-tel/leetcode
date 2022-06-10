package main

import "fmt"

func uintRevert(num uint32) uint32 {
	var res uint32

	for num != 0 {
		mod := num % 10
		num = num / 10
		res = res*10 + mod
	}
	return res
}

func main() {
	var num uint32
	num = 233455
	fmt.Println(uintRevert(num))
}
