package main
/*
	当你有看到一个接口类型的值时，你不知道它是什么，唯一知道的就是可以通过它的方法来做什么。
	fmt.Printf它会把结果写到标准输出和fmt.Sprintf它会把结果以字符串的形式返回。这两个函数都使用了另一个函数fmt.Fprintf来进行封装。
	io.Writer类型定义了函数Fprintf和这个函数调用者之间的约定。一方面这个约定需要调用者提供具体类型的值就像*os.File和*bytes.Buffer，这些类型都有一个特定签名和行为的Write的函数。另一方面
这个约定保证了Fprintf接受任何满足io.Writer接口的值都可以工作。
	一个类型可以自由的使用另一个满足相同接口的类型来进行替换被称作可替换性(LSP里氏替换)。这是一个面向对象的特征。

	除了io.Writer这个接口类型，还有另一个对fmt包很重要的接口类型。Fprintf和Fprintln函数向类型提供了一种控制它们值输出的途径。给一个类型定义String方法，可以让它满足最广泛使用之一的接口类型
fmt.Stringer.
*/
func main() {
	
}
