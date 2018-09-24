package main

import "fmt"

/*
	Go语言提供了两种精度的复数类型：complex64和complex128，分别对应float32和float64两种浮点数精度。内置的complex函数用于构建复数，内建的real和imag函数分别返回复数的实部和虚部
	在常量算术规则下，一个复数常量可以加到另一个普通数值常量（整数或浮点数、实部或虚部），我们可以用自然的方式书写复数，就像1+2i或与之等价的写法2i+1。
	复数也可以用==和!=进行相等比较。只有两个复数的实部和虚部都相等的时候它们才是相等的.
*/
func main() {
	var x complex128 = complex(1, 2) // 1+2i
	var y complex128 = complex(3, 4) // 3+4i
	fmt.Println(x*y) // "(-5+10i)"
	fmt.Println(real(x*y)) // "-5"
	fmt.Println(imag(x*y)) // "10"

	x = 1 + 2i
	y = 3 + 4i
}
