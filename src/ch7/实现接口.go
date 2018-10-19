package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

/*
	一个类型如果拥有一个接口需要的所有方法，那么这个类型就实现了这个接口。
	接口指定的规则非常简单：表达一个类型属于某个接口只要这个类型实现这个接口。

	在T类型的参数上调用一个T的方法是合法的，只要这个参数是一个变量；编译器隐式的获取了它的地址。但这仅仅是一个语法糖：T类型的值不拥有所有*T指针的方法，那这样它就可能只实现更少的接口。
	就像信封封装和隐藏信件起来一样，接口类型封装和隐藏具体类型和它的值。即使具体类型有其它的方法也只有接口类型暴露出来的方法会被调用到。

	***实际上interface{}被称为空接口类型是不可或缺的。因为空接口类型对实现它的类型没有要求，所以我们可以将任意一个值赋给空接口类型。
	对于创建的一个interface{}值持有一个boolean，float，string，map，pointer，或者任意其它的类型；我们当然不能直接对它持有的值做操作，因为interface{}没有任何方法。
	因为接口实现只依赖于判断的两个类型的方法，所以没有必要定义一个具体类型和它实现的接口之间的关系。
	每一个具体类型的组基于它们相同的行为可以表示成一个接口类型。不像基于类的语言，他们一个类实现的接口集合需要进行显式的定义，在Go语言中我们可以在需要的时候定义一个新的抽象或者特定特点
的组，而不需要修改具体类型的定义。
*/
func main() {
	var w io.Writer
	w = os.Stdout // OK: *os.File has Write method
	w = new(bytes.Buffer) // OK: *bytes.Buffer has Write method
	w = time.Second // compile error: time.Duration lacks Write method
	var rwc io.ReadWriteCloser
	rwc = os.Stdout // OK: *os.File has Read, Write, Close methods
	rwc = new(bytes.Buffer) // compile error: *bytes.Buffer lacks Close method
	fmt.Println(w, rwc)

	w = rwc //接口实现接口
}
