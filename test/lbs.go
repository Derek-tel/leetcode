package main

import (
	"fmt"
)

/**
Go两个goroutine交替打印1-10
例如输出结果：
G1 1
G2 2
G1 3
G2 4
…
G1 9
G2 10

*/

func main() {
	chanA := make(chan int, 1)
	chanB := make(chan int)
	end := make(chan int)

	first := func() {
		for i := 0; i < 5; i++ {
			<-chanA
			fmt.Println(i * 2)
			chanB <- 1
		}
	}

	second := func() {
		defer func() {
			end <- 1
		}()
		for i := 0; i < 5; i++ {
			<-chanB
			fmt.Println(i*2 + 1)
			chanA <- 1
		}
	}

	go first()
	go second()
	chanA <- 1

	<-end
}

/**
长度10初始化都为0的数组，顺序填入1-10，因为中间一个数值漏掉了，填到最后发现数组没有填满，找出漏掉了哪个数字？
例如给出填好的数组[1,2,3,4,5,7,8,9,10,0]，返回6    0-9   1-10
*/
