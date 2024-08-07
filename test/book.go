package main

import (
	"fmt"
)

func main() {
	chA := make(chan bool, 1)
	ChB := make(chan bool)
	exit := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			if ok := <-chA; ok {
				fmt.Println(i*2 + 1)
				ChB <- true
			}
		}
	}()

	go func() {
		defer func() {
			close(exit)
		}()
		for i := 1; i <= 10; i++ {
			if ok := <-ChB; ok {
				fmt.Println(i * 2)
				chA <- true
			}
		}
	}()

	chA <- true
	<-exit
}

func one() {
	chA := make(chan bool, 1)
	chB := make(chan bool)
	exit := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			if ok := <-chA; ok {
				fmt.Println(i*2 + 1)
				chB <- true
			}
		}
	}()
	go func() {
		defer func() { close(exit) }()
		for i := 1; i <= 10; i++ {
			if ok := <-chB; ok {
				fmt.Println(i * 2)
				chA <- true
			}
		}
	}()

	chA <- true
	<-exit
}
