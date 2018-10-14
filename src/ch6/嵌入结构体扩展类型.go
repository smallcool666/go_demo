package main

import (
	"fmt"
	"image/color"
)

/*
	内嵌可以使我们在定义ColoredPoint时得到一种句法上的简写形式，并使其包含Point类型所具有的一切字段，然后再定义一些自己的。
	可以直接认为通过嵌入的字段就是ColoredPoint自身的字段，而完全不需要在调用时指出Point。
	对于Point中的方法我们也有类似的用法，我们可以把ColoredPoint类型当作接收器来调用Point里的方法，即使ColoredPoint里没有声明这些方法。
	读者如果对基于类来实现面向对象的语言比较熟悉的话，可能会倾向于将Point看作一个基类，而ColoredPoint看作其子类或者继承类，或者将ColoredPoint看作"is a" Point类型。但
这是错误的理解。请注意上面例子中对Distance方法的调用。Distance有一个参数是Point类型，但q并不是一个Point类，所以尽管q有着Point这个内嵌类型，我们也必须要显式地选择它。
	在类型中内嵌的匿名字段也可能是一个命名类型的指针，这种情况下字段和方法会被间接地引入到当前的类型中(译注：访问需要通过该指针指向的对象去取)。添加这一层间接关系让我
们可以共享通用的结构并动态地改变对象之间的关系。
	一个struct类型也可能会有多个匿名字段。然后这种类型的值便会拥有所有匿名类型的所有方法，以及直接定义在结构体中的方法。
	方法只能在命名类型(像Point)或者指向类型的指针上定义，但是多亏了内嵌，有些时候我们给匿名struct类型来定义方法也有了手段。
*/

type ColoredPoint struct {
	Point
	Color color.RGBA
}
//内嵌指针类型的匿名字段
type ColoredPoin1 struct {
	*Point
	Color color.RGBA
}

func main() {
	//可以直接认为通过嵌入的字段就是ColoredPoint自身的字段，而完全不需要在调用时指出Point
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X) // "1"
	cp.Point.Y = 2
	fmt.Println(cp.Y) // "2"
	//也可以直接认为通过嵌入的结构体的方法就是ColoredPoint自身的方法，而完全不需要在调用时指出Point
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}
	fmt.Println(p.Distance(q.Point)) // "5"
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point)) // "10"
	//内嵌指针类型的匿名字段
	p1 := ColoredPoin1{&Point{1, 1}, red}
	q1 := ColoredPoin1{&Point{5, 4}, blue}
	fmt.Println(p.Distance(*q1.Point)) // "5"
	q.Point = p.Point // p and q now share the same Point
	p.ScaleBy(2)
	fmt.Println(*p1.Point, *q1.Point) // "{2 2} {2 2}"
}
