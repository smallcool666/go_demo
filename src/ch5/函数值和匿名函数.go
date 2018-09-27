package main

import (
	"fmt"
	"strings"
)

/*
	在Go中，函数被看作第一类值（first-class values）：函数像其他值一样，拥有类型，可以被赋值给其他变量，传递给函数，从函数返回。对函数值（function value）的调用类似函数调用。
	函数类型的零值是nil。调用值为nil的函数值会引起panic错误。
	函数值可以与nil比较。
	函数值之间是不可比较的，也不能用函数值作为map的key。
	函数值使得我们不仅仅可以通过数据来参数化函数，亦可通过行为。


	拥有函数名的函数只能在包级语法块中被声明，通过函数字面量（function literal），我们可绕过这一限制，在任何表达式中表示一个函数值。函数字面量的语法和函数声明相似，区别在于
func关键字后没有函数名。函数值字面量是一种表达式，它的值被成为匿名函数（anonymous function）。
	函数字面量允许我们在使用时函数时，再定义它。
	通过这种方式定义的函数可以访问完整的词法环境（lexical environment），这意味着在函数中定义的内部函数可以引用该函数的变量。
	squares的例子证明，函数值不仅仅是一串代码，还记录了状态。在squares中定义的匿名内部函数可以访问和更新squares中的局部变量，这意味着匿名函数和squares中，存在变量引用。这就
是函数值属于引用类型和函数值不可比较的原因。Go使用闭包（closures）技术实现函数值，Go程序员也把函数值叫做闭包。
	当匿名函数需要被递归调用时，我们必须首先声明一个变量，再将匿名函数赋值给这个变量。如果不分成两部，函数字面量无法与变量绑定，我们也无法递归调用该匿名函数。
*/

func square(n int) int { return n * n }
func negative(n int) int { return -n }
func product(m, n int) int { return m * n }
func add1(r rune) rune { return r + 1 }

//在函数中定义的内部函数可以引用该函数的变量
//函数squares返回另一个类型为 func() int 的函数。
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	//函数值的使用
	f := square
	fmt.Println(f(3)) // "9"
	f = negative
	fmt.Println(f(3)) // "-3"
	fmt.Printf("%T\n", f) // "func(int) int"
	//f = product // compile error: can't assign func(int, int) int to func(int) int

	//通过行为来参数化函数
	//strings.Map对字符串中的每个字符调用add1函数，并将每个add1函数的返回值组成一个新的字符串返回给调用者。
	fmt.Println(strings.Map(add1, "HAL-9000")) // "IBM.:111"
	fmt.Println(strings.Map(add1, "VMS")) // "WNT"
	fmt.Println(strings.Map(add1, "Admix")) // "Benjy"

	//使用函数时再定义
	strings.Map(func(r rune) rune { return r + 1 }, "HAL-9000")
	//在函数中定义的内部函数可以引用该函数的变量
	f1 := squares()
	fmt.Println(f1()) // "1"
	fmt.Println(f1()) // "4"
	fmt.Println(f1()) // "9"fmt.Println(f()) // "16"
}
