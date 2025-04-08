package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	rows := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	taskChan := make(chan int, len(rows)*2)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			defer func() {
				if r := recover(); r != nil {
					err, ok := r.(error)
					if !ok {
						err = fmt.Errorf("%v", r)
					}
					fmt.Printf("worker %d panic: %v\n", workerID, err)
				}
			}()
			for idx := range taskChan {
				time.Sleep(5 * time.Second)
				fmt.Printf("Worker %d processed index %d\n", workerID, idx)
			}
			fmt.Printf("Worker %d exit\n", workerID)
		}(i)
	}
	for _, row := range rows {
		taskChan <- row
	}
	close(taskChan)
	fmt.Println("xxxx")
	fmt.Println(fmt.Sprintf("%+v", len(taskChan)))

	wg.Wait()

	for idx := range taskChan {
		fmt.Printf("idx: %d\n", idx)
	}
	fmt.Println("main exit")
}
