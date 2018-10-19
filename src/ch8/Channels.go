package main

import "fmt"

/*
	如果说goroutine是Go语言程序的并发体的话，那么channelsj就是它们之间的通信机制。一个channels是一个通信机制，它可以让一个goroutine通过它给另一个goroutine发送值信息。每个channel都有一个
特殊的类型，也就是channels可发送数据的类型。一个可以发送int类型数据的channel一般写为chan int。
	使用内置的make函数，我们可以创建一个channel。
	和map类似，channel也一个对应make创建的底层数据结构的引用。和其它的引用类型一样，channel的零值也是nil。
	一个channel有发送和接受两个主要操作，都是通信行为。一个发送语句将一个值从一个goroutine通过channel发送到另一个执行接收操作的goroutine。发送和接收两个操作都是用<-运算符。在发送语句中，
<-运算符分割channel和要发送的值。在接收语句中，<-运算符写在channel对象之前。一个不使用接收结果的接收操作也是合法的。
	Channel还支持close操作，用于关闭channel，随后对基于该channel的任何发送操作都将导致panic异常。对一个已经被close过的channel之行接收操作依然可以接受到之前已经成功发送的数据；如果channel
中已经没有数据的话将产生一个零值的数据。
	以最简单方式调用make函数创建的时一个无缓存的channel，但是我们也可以指定第二个整形参数，对应channel的容量。如果channel的容量大于零，那么该channel就是带缓存的channel。

1. 不带缓存的Channels
	一个基于无缓存Channels的发送操作将导致发送者goroutine阻塞，直到另一个goroutine在相同的Channels上执行接收操作，当发送的值通过Channels成功传输之后，两个goroutine可以继续执行后面的
语句。反之，如果接收操作先发生，那么接收者goroutine也将阻塞，直到有另一个goroutine在相同的Channels上执行发送操作。
	基于无缓存Channels的发送和接收操作将导致两个goroutine做一次同步操作。因为这个原因，无缓存Channels有时候也被称为同步Channels。当通过一个无缓存Channels发送数据时，接收者收到数据发生
在唤醒发送者goroutine之前（译注：happens before，这是Go语言并发内存模型的一个关键术语！）。
	在讨论并发编程时，当我们说x事件在y事件之前发生（happens before），我们并不是说x事件在时间上比y时间更早；我们要表达的意思是要保证在此之前的事件都已经完成了。
	当我们说x事件既不是在y事件之前发生也不是在y事件之后发生，我们就说x事件和y事件是并发的。这并不是意味着x事件和y事件就一定是同时发生的，我们只是不能确定这两个事件发生的先后顺序。
	基于channels发送消息有两个重要方面。首先每个消息都有一个值，但是有时候通讯的事实和发生的时刻也同样重要。当我们更希望强调通讯发生的时刻时，我们将它称为消息事件。有些消息事件并不携带
额外的信息，它仅仅是用作两个goroutine之间的同步，这时候我们可以用struct{}空结构体作为channels元素的类型，虽然也可以使用bool或int类型实现同样的功能，done <- 1语句也比done <-struct{}{}更短。

2.串联的Channels（Pipeline）
	Channels也可以用于将多个goroutine链接在一起，一个Channels的输出作为下一个Channels的输入。这种串联的Channels就是所谓的管道（pipeline）。
	如果发送者知道，没有更多的值需要发送到channel的话，那么让接收者也能及时知道没有多余的值可接收将是有用的，因为接收者可以停止不必要的接收等待。这可以通过内置的close函数来关闭channel实现。
	当一个channel被关闭后，再向该channel发送数据将导致panic异常。当一个被关闭的channel中已经发送的数据都被成功接收后，后续的接收操作将不再阻塞，它们会立即返回一个零值。
	没有办法直接测试一个channel是否被关闭，但是接收操作有一个变体形式：它多接收一个结果，多接收的第二个结果是一个布尔值ok，ture表示成功从channels接收到值，false表示channels已经被关闭并且
里面没有值可接收。
	Go语言的range循环可直接在channels上面迭代。使用range循环是上面处理模式的简洁语法，它依次从channel接收数据，当channel被关闭并且没有值可接收时跳出循环。
	并不需要关闭每一个channel。只要当需要告诉接收者goroutine，所有的数据已经全部发送时才需要关闭channel。不管一个channel是否被关闭，当它没有被引用时将会被Go语言的垃圾自动回收器回收。
	试图重复关闭一个channel将导致panic异常，试图关闭一个nil值的channel也将导致panic异常。关闭一个channels还会触发一个广播机制。

3.单方向的Channel
	Go语言的类型系统提供了单方向的channel类型，分别用于只发送或只接收的channel。类型chan<- int表示一个只发送int的channel，只能发送不能接收。相反，类型<-chanint表示一个只接收int的channel，
只能接收不能发送。（箭头<-和关键字chan的相对位置表明了channel的方向。）这种限制将在编译期检测。
	因为关闭操作只用于断言不再向channel发送新的数据，所以只有在发送者所在的goroutine才会调用close函数，因此对一个只接收的channel调用close将是一个编译错误。
	在函数参数中定义了单向channels，传进来双向channels会有隐式转换。

4.带缓存的Channels
	带缓存的Channel内部持有一个元素队列。队列的最大容量是在调用make函数创建channel时通过第二个参数指定的。
	向缓存Channel的发送操作就是向内部缓存队列的尾部插入元素，接收操作则是从队列的头部删除元素。如果内部缓存队列是满的，那么发送操作将阻塞直到因另一个goroutine执行接收操作而释放了新的队列
空间。相反，如果channel是空的，接收操作将阻塞直到有另一个goroutine执行发送操作而向队列插入元素。
	在某些特殊情况下，程序可能需要知道channel内部缓存的容量，可以用内置的cap函数获取。对于内置的len函数，如果传入的是channel，那么将返回channel内部缓存队列中有效元素的个数。
	如果我们使用了无缓存的channel，那么两个慢的goroutines将会因为没有人接收而被永远卡住。这种情况，称为goroutines泄漏，这将是一个BUG。和垃圾变量不同，泄漏的goroutines并不会被自动回收，因
此确保每个不再需要的goroutine能正常退出是重要的。
*/

func main() {
	ch := make(chan int) // ch has type 'chan int
	fmt.Printf("%T\n", ch)
	//channel发送接收
	x := 111
	ch <- x // a send statement
	x = <-ch // a receive expression in an assignment statement
	<-ch // a receive statement; result is discarded
}
