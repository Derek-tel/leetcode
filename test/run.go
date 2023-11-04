package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	i := 0
	go func() {
		fmt.Println(i)
		panic(any("already call"))
	}()

	for {
		i = i + 1
	}

	sliceA := make([]int, 0, 4)
	sliceA = append(sliceA, 1, 2, 3)
	sliceA = testSlice(sliceA)
	fmt.Println(sliceA)
}

func testSlice(sliceA []int) []int {
	sliceA = append(sliceA, 4)
	fmt.Println(sliceA)
	return sliceA
}
