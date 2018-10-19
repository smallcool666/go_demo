package main

import "fmt"

/*
	select语句的一般形式。和switch语句稍微有点相似，也会有几个case和最后的default选择支。每一个case代表一个通信操作(在某个channel上进行发送或者接收)并且会包含一些语句组成的一个语句
块。一个接收表达式可能只包含接收表达式自身(译注：不把接收到的值赋值给变量什么的)，就像上面的第一个case，或者包含在一个简短的变量声明中，像第二个case里一样；第二种形式让你能够引用接收
到的值。
	select会等待case中有能够执行的case时去执行。当条件满足时，select才会去通信并执行case之后的语句；这时候其它通信是不会执行的。一个没有任何case的select语句写作select{}，会永远地等
待下去。
	如果多个case同时就绪时，select会随机地选择一个执行，这样来保证每一个channel都有平等的被select的机会。增加前一个例子的buffer大小会使其输出变得不确定，因为当buffer既不为满也不为空
时，select语句的执行情况就像是抛硬币的行为一样是随机的。
	有时候我们希望能够从channel中发送或者接收值，并避免因为发送或者接收导致的阻塞，尤其是当channel没有准备好写或者读时。select语句就可以实现这样的功能。select会有一个default来设置当
其它的操作都不能够马上被处理时程序需要执行哪些逻辑。
	channel的零值是nil。也许会让你觉得比较奇怪，nil的channel有时候也是有一些用处的。因为对一个nil的channel发送和接收操作会永远阻塞，在select语句中操作nil的channel永远都不会被select到。
*/
func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x) // "0" "2" "4" "6" "8"
		case ch <- i:
		}
	}
}
