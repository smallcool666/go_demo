package main

import "fmt"

/*
	在现代计算机中可能会有一堆处理器，每一个都会有其本地缓存(local cache)。为了效率，对内存的写入一般会在每一个处理器中缓冲，并在必要时一起flush到主存。这种情况下这些数据可能会以与当初
goroutine写入顺序不同的顺序被提交到主存。像channel通信或者互斥量操作这样的原语会使处理器将其聚集的写入flush并commit，这样goroutine在某个时间点上的执行结果才能被其它处理器上运行的
goroutine得到。
	所有并发的问题都可以用一致的、简单的既定的模式来规避。所以可能的话，将变量限定在goroutine内部；如果是多个goroutine都需要访问的变量，使用互斥条件来访问。
	因为赋值和打印指向不同的变量，编译器可能会断定两条语句的顺序不会影响执行结果，并且会交换两个语句的执行顺序。如果两个goroutine在不同的CPU上执行，每一个核心有自己的缓存，这样一个
goroutine的写入对于其它goroutine的Print，在主存同步之前就是不可见的了。
	所有并发的问题都可以用一致的、简单的既定的模式来规避。所以可能的话，将变量限定在goroutine内部；如果是多个goroutine都需要访问的变量，使用互斥条件来访问。
*/

func main() {
	var x, y int
	go func() {
		x = 1 // A1
		fmt.Print("y:", y, " ") // A2
	}()
	go func() {
		y = 1 // B1
		fmt.Print("x:", x, " ") // B2
	}()
}
