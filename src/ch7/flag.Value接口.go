package main

import (
	"flag"
	"fmt"
	"time"
)

/*
	本节了解标准的接口类型flag.Value是怎么帮助命令行标记定义新的符号的。
	flag.Duration函数创建一个time.Duration类型的标记变量并且允许用户通过多种用户友好的方式来设置这个变量的大小，这种方式还包括和String方法相同的符号排版形式。这种对称设计使得用户交互良好。
	因为时间周期标记值非常的有用，所以这个特性被构建到了flag包中；但是我们为我们自己的数据类型定义新的标记符号是简单容易的。我们只需要定义一个实现flag.Value接口的类型（只有String和Set
两个方法）。
*/

var period = flag.Duration("period", 1*time.Second, "sleep period")
func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()

	//$ ./sleep
	//Sleeping for 1s...

	//默认情况下，休眠周期是一秒，但是可以通过 -period 这个命令行标记来控制。flag.Duration函数创建一个time.Duration类型的标记变量并且允许用户通过多种用户友好的方式来设置这个变量的大小
	//$ ./sleep -period 50ms
	//Sleeping for 50ms...
	//$ ./sleep -period 2m30s
	//Sleeping for 2m30s...
}
