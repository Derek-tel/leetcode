package main

import (
	"fmt"
	"runtime"
)

// 获取正在运行的函数名
func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func test1() {
	i := 0
	fmt.Println("i =", i)
	fmt.Println("FuncName1 =", runFuncName())
}

func test2() {
	i := 1
	fmt.Println("i =", i)
	fmt.Println("FuncName2 =", runFuncName())
}

type a struct {
	b
}

func (ss a) printStr(str string) {
	fmt.Println("aaa")
}

type b struct {
	i int
}

func (s b) printStr(i int) {
	fmt.Println("bbb")
}

func main() {
	fmt.Println("打印运行中的函数名")
	test1()
	test2()

	a := a{}
	a.printStr("sss")
	a.b.printStr(1)
	b := b{}
	b.printStr(1)
}
