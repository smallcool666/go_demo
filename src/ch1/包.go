package main

/*
	Go语言中的包和其他语言的库或模块的概念类似，目的都是为了支持模块化、封装、单独编译和代码重用。
	每个包都对应一个独立的名字空间。
	包还可以让我们通过控制哪些名字是外部可见的来隐藏内部实现信息。在Go语言中，一个简单的规则是：如果一个名字是大写字母开头的，那么该名字是导出的（译注：因为汉字不区分大小
写，因此汉字开头的名字是没有导出的）。
	每个源文件都是以包的声明语句开始，用来指名包的名字。当包被导入的时候，包内的成员将通过类似tempconv.CToF的形式访问。而包级别的名字，例如在一个文件声明的类型和常量，在同一个包的其他源
文件也是可以直接访问的，就好像所有代码都在一个文件一样。
	在每个源文件的包声明前仅跟着的注释是包注释。
*/
import (
	"ch1/example/tempconv"
	"fmt"
)
func main() {
	//因为包级别的常量名都是以大写字母开头，它们可以像tempconv.AbsoluteZeroC这样被外部代码访问
	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC) // "Brrrr! -273.15°C"
	fmt.Println(tempconv.CToF(tempconv.BoilingC)) // "212°F"
}
