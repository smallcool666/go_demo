package main
/*
	Go的runtime和工具链为我们装备了一个复杂但好用的动态分析工具，竞争检查器(the race detector)。
	只要在go build，go run或者go test命令后面加上-race的flag，就会使编译器创建一个你的应用的“修改”版或者一个附带了能够记录所有运行期对共享变量访问工具的test，并且会记录下每一个读
或者写共享变量的goroutine的身份信息。另外，修改版的程序会记录下所有的同步事件，比如go语句，channel操作，以及对(*sync.Mutex).Lock，(*sync.WaitGroup).Wait等等的调用。
	竞争检查器会报告所有的已经发生的数据竞争。然而，它只能检测到运行时的竞争条件；并不能证明之后不会发生数据竞争。所以为了使结果尽量正确，请保证你的测试并发地覆盖到了你到包。
*/

func main() {
	
}
