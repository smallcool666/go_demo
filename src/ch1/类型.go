package main

import "fmt"

/*
	变量或表达式的类型定义了对应存储值的属性特征，例如数值在内存的存储大小（或者是元素的bit个数），它们在内部是如何表达的，是否支持一些操作符，以及它们自己关联的方法集等。
	一个类型声明语句创建了一个新的类型名称，和现有类型具有相同的底层结构。新命名的类型提供了一个方法，用来分隔不同概念的类型，这样即使它们底层类型相同也是不兼容的。
	类型声明语句一般出现在包一级，因此如果新创建的类型名字的首字符大写，则在外部包也可以使用。
	对于每一个类型T，都有一个对应的类型转换操作T(x)，用于将x转为T类型（译注：如果T是指针类型，可能会需要用小括弧包装T，比如(*int)(0)）。只有当两个类型的底层基础类型相同时，
才允许这种转型操作，或者是两者都是指向相同底层结构的指针类型，这些转换只改变类型而不会影响值本身。
	底层数据类型决定了内部结构和表达方式，也决定是否可以像底层类型一样对内置运算符的支持。
	比较运算符==和<也可以用来比较一个命名类型的变量和另一个有相同类型的变量，或有着相同底层类型的未命名类型的值之间做比较。但是如果两个值有着不同的类型，则不能直接进行比较
	一个命名的类型可以提供书写方便，特别是可以避免一遍又一遍地书写复杂类型（译注：例如用匿名的结构体定义变量）。虽然对于像float64这种简单的底层类型没有简洁很多，但是如果是
复杂的类型将会简洁很多，特别是结构体类型。

	***命名类型还可以为该类型的值定义新的行为。这些行为表示为一组关联到该类型的函数集合，我们称为类型的方法集。
*/

/*
	在这个包声明了两种类型：Celsius和Fahrenheit分别对应不同的温度单位。它们虽然有着相同的底层类型float64，但是它们是不同的数据类型，因此它们不可以被相互比较或混在一个表达
式运算。刻意区分类型，可以避免一些像无意中使用不同单位的温度混合计算导致的错误；因此需要一个类似Celsius(t)或Fahrenheit(t)形式的显式转型操作才能将float64转为对应的类型。
Celsius(t)和Fahrenheit(t)是类型转换操作，它们并不是函数调用。类型转换不会改变值本身，但是会使它们的语义发生变化。
*/
type Celsius float64 // 摄氏温度
type Fahrenheit float64 // 华氏温度
const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC Celsius = 0 // 结冰点温度
	BoilingC Celsius = 100 // 沸水温度
)

 func CToF(c Celsius) Fahrenheit {
 	return Fahrenheit(c*9/5 + 32)
 }
 func FToC(f Fahrenheit) Celsius {
 	return Celsius((f - 32) * 5 / 9)
 }

 func main(){
	 //底层数据类型决定了内部结构和表达方式，也决定是否可以像底层类型一样对内置运算符的支持。
	 fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
	 boilingF := CToF(BoilingC)
	 fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180" °F
	 //fmt.Printf("%g\n", boilingF-FreezingC) // compile error: type mismatch

	 //比较运算符==和<也可以用来比较一个命名类型的变量和另一个有相同类型的变量，或有着相同底层类型的未命名类型的值之间做比较。但是如果两个值有着不同的类型，则不能直接进行比较
	 var c Celsius
	 var f Fahrenheit
	 fmt.Println(c == 0) // "true"
	 fmt.Println(f >= 0) // "true"
	 //fmt.Println(c == f) // compile error: type mismatch
	 fmt.Println(c == Celsius(f)) // "true"!
 }