package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Product struct {
	value int
}

//采用通道实现产品双向传递，缓冲区设为10
//通道的运作方式替代了伪代码中 empty 和 full 信号量的作用
//通道的另一个特性：任意时刻只能有一个协程能对 channel 中某一个item进行访问。替代了伪代码中 mutex 信号量的作用
var ch = make(chan Product, 10)

//终止生产者生产的信号，否则无限执行下去
var stop = false

func Producer(wg *sync.WaitGroup, p_num int) {
	for !stop { //不断生成，直到主线程执行到终止信号
		p := Product{value: rand.Int()}
		ch <- p
		fmt.Printf("producer %v produce a product: %#vn", p_num, p)
		time.Sleep(time.Duration(200+rand.Intn(1000)) * time.Millisecond) //延长执行时间
	}
	wg.Done()
}

func Consumer(wg *sync.WaitGroup, c_num int) {
	//通道里没有产品了就停止消费
	for p := range ch {
		fmt.Printf("consumer %v consume a product: %#vn", c_num, p)
		time.Sleep(time.Duration(200+rand.Intn(1000)) * time.Millisecond)
	}
	wg.Done()
}

func main() { //主线程
	var wgp sync.WaitGroup
	var wgc sync.WaitGroup
	wgp.Add(5)
	wgc.Add(5)
	//设5个生产者、5个消费者
	for i := 0; i < 5; i++ {
		go Producer(&wgp, i)
		go Consumer(&wgc, i)
	}
	time.Sleep(time.Duration(1) * time.Second)
	stop = true
	wgp.Wait()
	close(ch)
	wgc.Wait()
}
