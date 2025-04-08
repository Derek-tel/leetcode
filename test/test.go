package main

import (
	"fmt"
)

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}

	//copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	//fmt.Println(slice2)

	copy(slice1, slice2)
	fmt.Println(slice1)

	var a uint16 = 0
	a = a | (1 << 15)
	a = a | (1 << 13)
	a = a | (1 << 0)
	fmt.Println(fmt.Sprintf("%x", a)) //a输出结果:[a001]

}
