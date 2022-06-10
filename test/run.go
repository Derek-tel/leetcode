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
}
