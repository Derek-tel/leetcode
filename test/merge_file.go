package main

import (
	"fmt"
	"sync"
)

// 请编写一个Golang程序，并发地从多个URL下载文件，并将下载的文件合并到一个文件中。
func mergeFile() {
	wg := sync.WaitGroup{}
	ch := make(chan string, 10)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			//download file
			ch <- "file"
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for s := range ch {
		//write to file
		fmt.Println(s)
	}

}

func main() {
	t := test()
	fmt.Println(t)
	// print 1
}

func test() (i int) {
	i = 1
	defer func() {
		i++
		fmt.Println("defer1=", i)
	}()

	defer func() {
		i++
		fmt.Println("defer2=", i)
	}()

	return i
}
