// 要求：
// 1. 只能编辑 foo 函数
// 2. foo 必须要调用 slow 函数
// 3. foo 函数在 ctx 超时后必须立刻返回
// 4. 【加分项】如果 slow 结束的比 ctx 快，也立刻返回
package main

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	foo(ctx)
	fmt.Println(strings.Split("", ","))
	fmt.Println(len(strings.Split("", ",")))
}

// 您需要实现 foo 函数，要求 foo 在 ctx 超时后立即返回
// foo 必须调用 slow 函数
func foo(ctx context.Context) {
	done := make(chan struct{})
	go func() {
		slow()
		close(done)
	}()

	select {
	case <-done:
		return
	case <-ctx.Done():
		return
	}
}

// 您不能修改 slow 函数
func slow() {
	n := rand.Intn(3)
	fmt.Printf("sleep %ds\n", n)
	time.Sleep(time.Duration(n) * time.Second)
}
