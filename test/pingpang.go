package main

import (
	"fmt"
)

func main() {
	var chanA = make(chan int)
	var chanB = make(chan int, 1)
	var exit = make(chan int)

	go func() {
		defer func() {
			close(chanA)
			close(exit)
		}()
		for i := 0; i < 10; i++ {
			if _, ok := <-chanA; ok {
				fmt.Println("A")
				chanB <- 1
			}
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			if _, ok := <-chanB; ok {
				fmt.Println("B")
				chanA <- 1
			}
		}
	}()

	chanB <- 1

	<-exit

	close(chanB)
}
