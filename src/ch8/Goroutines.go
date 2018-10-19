package main

import (
	"fmt"
	"time"
)

/*
	Go语言中的并发程序可以用两种手段来实现。这一章会讲解goroutine和channel，其支持“顺序进程通信”(communicating sequential processes)或被简称为CSP。CSP是一个现代的并发编程模型，在这种
编程模型中值会在不同的运行实例(goroutine)中传递，尽管大多数情况下被限制在单一实例中。
	在Go语言中，每一个并发的执行单元叫作一个goroutine。
	当一个程序启动时，其主函数即在一个单独的goroutine中运行，我们叫它main goroutine。新的goroutine会用go语句来创建。在语法上，go语句是一个普通的函数或方法调用前加上关键字go。go语句
会使其语句中的函数在一个新创建的goroutine中运行。而go语句本身会迅速地完成。
	当主函数返回时，所有的goroutine都会直接打断，程序退出。除了从主函数退出或者直接退出程序之外，没有其它的编程方法能够让一个goroutine来打断另一个的执行，但是我们之后可以看到，可以通过
goroutine之间的通信来让一个goroutine请求请求其它的goroutine，并让其自己结束执行。

	go后跟的函数的参数会在go语句自身执行时被求值。
	让服务使用并发不只是处理多个客户端的请求，甚至在处理单个连接时也可能会用到。
*/
func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}