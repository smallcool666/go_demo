package main

/*
	接口类型具体描述了一系列方法的集合，一个实现了这些方法的具体类型是这个接口类型的实例。
	和结构内嵌相似的语法，我们可以用这种方式以一个简写命名另一个接口，而不用声明它所有的方法。这种方式本称为接口内嵌。
*/

type ReadWriter interface {
	Read(p []byte) (n int, err error) //不适用内嵌声明接口方法
	Writer //io包中的接口
}
type ReadWriteCloser interface {
Reader
Writer
Closer
}

func main() {
	
}
