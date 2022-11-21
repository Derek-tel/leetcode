package main

import (
	"fmt"
	"unsafe"
)

type Admin struct {
	Name string
	Age  int
}

func main() {
	ch := make(chan int)
	go Feibo(5, ch)
	fmt.Println("hello :", <-ch)

	unsafeTest()
	uintptrTest()

}

func unsafeTest() {
	i := 30
	iPtr1 := &i

	var iPtr2 *int64 = (*int64)(unsafe.Pointer(iPtr1))

	*iPtr2 = 8

	fmt.Println(i)
}

func uintptrTest() {
	admin := Admin{
		Name: "seekload",
		Age:  18,
	}
	ptr := &admin
	fmt.Println(uintptr(unsafe.Pointer(ptr)))
	fmt.Println(unsafe.Offsetof(ptr.Age))
	name := (*string)(unsafe.Pointer(ptr)) // 1

	*name = "四哥"

	fmt.Println(*ptr)

	age := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + unsafe.Offsetof(ptr.Age))) // 2
	*age = 35

	fmt.Println(*ptr)
}

func Feibo(n int, ch chan<- int) {
	if n <= 0 {
		ch <- 0
	} else if n < 3 {
		ch <- 1
	} else {
		ch1 := make(chan int)
		go Feibo(n-1, ch1)
		fmt.Println("goroutine ch1")
		ch2 := make(chan int)
		go Feibo(n-2, ch2)
		fmt.Println("goroutine ch2")
		ch <- <-ch1 + <-ch2
	}

}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func mainTest() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
