package main

import (
	"fmt"
	"sync"
)

func main() {

	var web3 sync.WaitGroup
	var chanInt = make(chan int)

	for i := 0; i < 10; i++ {
		go func(j int) {
			//get order info
			web3.Add(1)
			defer web3.Done()

			chanInt <- j
		}(i)
	}

	go func() {
		web3.Wait()
		close(chanInt)
	}()

	for order := range chanInt {
		fmt.Println(order)
	}
}
