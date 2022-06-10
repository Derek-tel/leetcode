package main

import "fmt"

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
	demo := map[string]string{
		"a":   "b",
		"aa":  "bb",
		"aaa": "bbbb",
	}
	fmt.Println(Map2Kvs(demo))
}

func Map2Kvs(m map[string]string) []string {
	kvs := []string{}
	for k, v := range m {
		kvs = append(kvs, k, v)
	}
	return kvs
}
