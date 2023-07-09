package main

import (
	"fmt"
	"github.com/imroc/biu"
	"sync"
)

func main() {

	dic := make(map[int]int)
	ch := make(chan int)

	wg := sync.WaitGroup{}
	go func() {
		key := 1
		wg.Add(1)
		defer wg.Done()
		_, ok := dic[key]
		fmt.Println(ok)
		if !ok {
			<-ch
			val, _ := dic[key]
			fmt.Println(val)
		}
	}()

	go func() {
		wg.Add(1)
		defer wg.Done()
		key := 1
		dic[key] = 1
		ch <- 1
	}()

	wg.Wait()

	fmt.Println(dic)

	const (
		a = iota
		b
		c
		d = "ha"
		e
		f = 100
		g
		h = iota
		i
	)
	const (
		mutexLocked           = 1 << iota // 持有锁的标记
		mutexWoken                        //唤醒标记
		mutexStarving                     // 饥饿标记
		mutexWaiterShift      = iota      //阻塞等待的数量
		starvationThresholdNs = 1e6       //饥饿阈值 1ms
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
	fmt.Println(mutexLocked, mutexWoken, mutexStarving, mutexWaiterShift, starvationThresholdNs)
	slc := make([]int, 100)
	leakSlice := slc[:10]
	fmt.Println(cap(leakSlice))
	leakSlice1 := slc[:10:10]
	fmt.Println(cap(leakSlice1))

	x := int8(+1)
	y := int8(-10)
	fmt.Println(biu.ToBinaryString(x))
	fmt.Println(biu.ToBinaryString(y))

}
