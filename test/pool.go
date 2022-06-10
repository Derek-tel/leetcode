package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var pool = sync.Pool{
	New: func() any {
		return "123"
	},
}

func main() {
	t := pool.Get().(string)
	fmt.Println(t)

	pool.Put("321")
	pool.Put("321")
	pool.Put("321")
	pool.Put("321")

	runtime.GC()
	time.Sleep(1 * time.Second)

	t2 := pool.Get().(string)
	fmt.Println(t2)

	runtime.GC()
	time.Sleep(1 * time.Second)

	t2 = pool.Get().(string)
	fmt.Println(t2)
}
