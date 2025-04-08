package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	f, _ := os.Create("trace.output")
	defer f.Close()
	_ = trace.Start(f)
	defer trace.Stop()
	var wg sync.WaitGroup
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			t := 0
			for i := 0; i < 1e8; i++ {
				t += 2
			}
			fmt.Println("total:", t)
		}()
	}
	wg.Wait()
}
