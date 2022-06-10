package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

// 一段有问题的代码
func do() {
	var c chan int
	for {
		select {
		case v := <-c:
			fmt.Printf("我是有问题的那一行，因为收不到值：%v", v)
		default:
		}
	}
}

func main() {

	// 创建分析文件
	file, err := os.Create("./cpu-profile.prof")
	if err != nil {
		fmt.Printf("创建采集文件失败, err:%v\n", err)
		return
	}

	// 进行cpu数据的获取
	pprof.StartCPUProfile(file)
	defer pprof.StopCPUProfile()
	// 执行一段有问题的代码
	for i := 0; i < 50; i++ {
		go do()
	}
	time.Sleep(10 * time.Second)

	println("程序执行完毕...")
}
